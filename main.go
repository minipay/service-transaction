package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"time"

	"service-transaction/middleware"
	transactionHttpDeliver "service-transaction/transaction/delivery/http"
	transactionRepository "service-transaction/transaction/repository"
	transactionUsecase "service-transaction/transaction/usecase"
	"service-transaction/utils/errorcollector"

	//"github.com/bugsnag/bugsnag-go"
	//bugsnagnegroni "github.com/bugsnag/bugsnag-go/negroni"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"
	"github.com/urfave/negroni"
)

func init() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool("debug") {
		log.Println("DEBUG mode is enabled")
	}
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			errorcollector.WritePanic(r, debug.Stack())
			panic(r)
		}
	}()
	dbConn, err := sql.Open("mysql", viper.GetString("database.mysql.user")+":"+viper.GetString("database.mysql.password")+"@tcp("+viper.GetString("database.mysql.host")+":"+viper.GetString("database.mysql.port")+")/"+viper.GetString("database.mysql.dbname")+"?charset=utf8&parseTime=True&loc=Local")
	dbConn.SetMaxOpenConns(100)
	dbConn.SetMaxIdleConns(100)
	if err != nil && viper.GetBool("debug") {
		log.Println(err)
	}
	err = dbConn.Ping()
	if err != nil {
		panic(fmt.Errorf("Error ping : %s", err))
	}
	defer func() {
		dbConn.Close()
	}()
	router := httprouter.New()
	n := negroni.New()
	mid := middleware.InitMiddleware()
	n.Use(negroni.HandlerFunc(mid.CORS))
	n.Use(negroni.HandlerFunc(mid.PanicCatcher))
	n.UseHandler(router)
	if viper.GetBool("debug") {
		n.Use(negroni.NewLogger())
	}
	timeoutContext := time.Duration(viper.GetInt("server.timeout")) * time.Second
	tr := transactionRepository.NewMysqlTransactionRepository(dbConn)
	tu := transactionUsecase.NewTransactionUsecase(tr, timeoutContext)
	transactionHttpDeliver.NewTransactionHandler(router, tu)

	s := &http.Server{
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         viper.GetString("server.host") + ":" + viper.GetString("server.port"),
		Handler:      n,
	}
	s.ListenAndServe()
}
