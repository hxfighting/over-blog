package service

import (
	"blog/config"
	"blog/helper"
	"bufio"
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	stdLog "log"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

var (
	Log *logrus.Logger
)

func NewLogger() {
	Log = logrus.New()
	Log.SetReportCaller(true)
	if !helper.CheckDebug() {
		formater := &logrus.JSONFormatter{
			TimestampFormat: helper.YMDHIS,
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				s := strings.Split(f.Function, ".")
				funcname := s[len(s)-1]
				_, filename := path.Split(f.File)
				return funcname, filename + ":" + fmt.Sprintf("%d", f.Line)
			},
		}
		setNull()
		configPath := config.ConfigPath + "/blog.log"
		writer, err := rotatelogs.New(
			configPath+".%Y%m%d",
			rotatelogs.WithLinkName(configPath),
			rotatelogs.WithRotationTime(24*time.Hour),
			rotatelogs.WithRotationCount(30),
		)
		if err != nil {
			stdLog.Fatalln(err.Error())
		}
		Log.AddHook(lfshook.NewHook(
			lfshook.WriterMap{
				logrus.InfoLevel:  writer,
				logrus.ErrorLevel: writer,
				logrus.DebugLevel: writer,
				logrus.FatalLevel: writer,
				logrus.PanicLevel: writer,
				logrus.WarnLevel:  writer,
				logrus.TraceLevel: writer,
			},
			formater,
		))
	} else {
		Log.Formatter = &logrus.TextFormatter{
			TimestampFormat: helper.YMDHIS,
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				s := strings.Split(f.Function, ".")
				funcname := s[len(s)-1]
				_, filename := path.Split(f.File)
				return "方法:" + funcname + " 信息:", "文件:" + filename + " 行号:" + fmt.Sprintf("%d", f.Line)
			},
		}
	}
}

/**
非调试模式就不输出到控制台
*/
func setNull() {
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	writer := bufio.NewWriter(src)
	Log.SetOutput(writer)
}
