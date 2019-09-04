package errorcollector

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

// WriteLog is use for write log to defined log file
func WriteLog(logMsg string) {
	f, err := os.OpenFile(viper.GetString("log.path")+viper.GetString("log.filename"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
	}
	defer f.Close()
	var buf bytes.Buffer
	logger := log.New(&buf, "[", log.Ldate|log.Ltime)
	logger.SetOutput(f)
	logger.Println(logMsg, "]")
	fmt.Print(&buf)
}

// WritePanic is use for write panic to defined log file
func WritePanic(e interface{}, stack []byte) {
	// var buf bytes.Buffer
	// // logger := log.New(&buf, "", log.Ldate|log.Ltime|log.Llongfile)
	// logger := log.New(&buf, "", log.Ldate|log.Ltime)
	// logger.Println(err)
	// fmt.Print(&buf)

	f, err := os.OpenFile(viper.GetString("log.path")+viper.GetString("log.filename"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
	}
	defer f.Close()
	var buf bytes.Buffer
	logger := log.New(&buf, "[", log.Ldate|log.Ltime)
	logger.SetOutput(f)
	logger.Println(e, "] \n", string(stack))
	fmt.Print(&buf)
}

// // MonitorPanic for handling if panic is happen
// func MonitorPanic() {
// 	exitStatus, err := panicwrap.BasicWrap(panicHandler)
// 	if err != nil {
// 		// Something went wrong setting up the panic wrapper. Unlikely,
// 		// but possible.
// 		panic(err)
// 	}
// 	// If exitStatus >= 0, then we're the parent process and the panicwrap
// 	// re-executed ourselves and completed. Just exit with the proper status.
// 	if exitStatus >= 0 {
// 		os.Exit(exitStatus)
// 	}
// }

// func panicHandler(output string) {
// 	WriteLog(output)
// 	os.Exit(1)
// }
