package service

import (
	"blog/config"
	"github.com/kataras/golog"
	"os"
	"time"
)

var (
	Log *golog.Logger
)

func getFileName() string {
	today := time.Now().Format("20060102")
	return "blog_" + today + ".log"
}

func NewLogFile() *os.File {
	filename := config.ConfigPath + string(os.PathSeparator) + getFileName()
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return f
}
