package gamelog

import (
	"log"
	"os"
)

var logFile *os.File

func InitLogger() error {
	logfile, err := os.Create("game.log")
	if err != nil {
		return err
	}

	log.SetOutput(logfile)
	return nil
}

func CloseLogger() {
	// 关闭日志文件
	if logFile != nil {
		logFile.Close()
	}
}

func Log(message string) {
	log.Println(message)
}
