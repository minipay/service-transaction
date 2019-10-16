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
	writer(logMsg, "]")
}

// WritePanic is use for write panic to defined log file
func WritePanic(e interface{}, stack []byte) {
	writer(e, "] \n", string(stack))
}

func writer(errFormat ...interface{}) {
	logFilename := viper.GetString("log.filename")
	if viper.GetString("log.filename") == "" {
		logFilename = "error.log"
	}
	f, err := os.OpenFile(viper.GetString("log.path")+logFilename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
	}
	defer f.Close()
	var buf bytes.Buffer
	logger := log.New(&buf, "[", log.Ldate|log.Ltime)
	logger.SetOutput(f)
	logger.Println(errFormat...)
	fmt.Print(&buf)
}
