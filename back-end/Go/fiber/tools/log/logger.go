package log

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rs/zerolog"

	"github.com/ohdata/blog/configs"
	"github.com/ohdata/blog/tools/util"
)

var (
	Log zerolog.Logger
)

func New() {
	zerolog.TimeFieldFormat = util.YMDHIS
	if configs.Config.Log.Output == "console" {
		Log = zerolog.New(os.Stderr).With().Caller().Timestamp().Logger().Output(formatConsoleLog())
	} else {
		Log = zerolog.New(getFileWriter("debug_log")).With().Timestamp().Logger()
	}
}

func formatConsoleLog() zerolog.ConsoleWriter {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: util.YMDHIS}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatMessage = func(i interface{}) string {
		if i != nil {
			return fmt.Sprintf("***%s***", i)
		}
		return ""
	}
	return output
}

func getFileWriter(name string) *rotatelogs.RotateLogs {
	filePath := filepath.Join(configs.Config.Log.Path, name)
	logf, err := rotatelogs.New(
		filePath+".%Y%m%d",
		rotatelogs.WithLinkName(filePath),
		rotatelogs.WithRotationCount(180),
	)
	if err != nil {
		panic(err)
	}
	return logf
}
