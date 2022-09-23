package handler

import (
	"fmt"
	"log"
	"os"
)

// 记录日志
// eg: log.Println(fmt.Sprintf("UID: %s", userID), mString)
func NotesLog() (err error) {
	logFile, err := os.OpenFile("./log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetPrefix("[log]")
	log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.Ldate)
	return
}
