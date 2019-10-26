package service

import (
	"blog/config"
	"github.com/phachon/go-logger"
)

var (
	Log *go_logger.Logger
)

func NewLog() {
	path := config.ConfigPath + "/"
	logger := go_logger.NewLogger()
	logger.Detach("console")
	fileConfig := &go_logger.FileConfig{
		Filename: path + "blog.log",
		LevelFileName: map[int]string{
			logger.LoggerLevel("error"): path + "error.log",
			logger.LoggerLevel("info"):  path + "info.log",
			logger.LoggerLevel("debug"): path + "debug.log",
		},
		MaxSize:    1024 * 1024,
		MaxLine:    10000,
		DateSlice:  "d",
		JsonFormat: false,
		Format:     "%millisecond_format% [%level_string%] [%file%:%line%] %body%",
	}
	logger.Attach("file", go_logger.LOGGER_LEVEL_DEBUG, fileConfig)
	Log = logger
}
