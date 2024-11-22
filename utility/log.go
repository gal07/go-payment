package utility

import (
	"log"
	"os"
)

func CreateLog(msg error) {
	// create info log
	fileInfo, err := openLogFile("../writeable/info.log")
	if err != nil {
		log.Fatal(err)
	}
	infoLog := log.New(fileInfo, "[info]", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
	infoLog.Println(msg)

	// create error log
	fileError, err := openLogFile("../writeable/error.log")
	if err != nil {
		log.Fatal(err)
	}
	errorLog := log.New(fileError, "[error]", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
	errorLog.Println(msg)
}

func openLogFile(path string) (*os.File, error) {
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}
