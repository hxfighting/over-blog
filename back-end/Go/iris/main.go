package main

import (
	"blog/config"
	"blog/controllers/frontend"
	"blog/database"
	"blog/helper"
	"blog/queue"
	"blog/routes"
	"blog/service"
	stdContext "context"
	"flag"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	requestLogger "github.com/kataras/iris/v12/middleware/logger"
	"log"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var (
	app = iris.New()
)

func init() {
	getConfigPath()
	config.NewConfig()
	service.NewLogger()
	database.NewDB()
	service.NewGeoDb()
	service.NewRedis()
	service.NewEmail()
	frontend.InitData()
	if e := sentry.Init(sentry.ClientOptions{
		Dsn: config.GetConfig("sentry.dsn").(string),
	}); e != nil {
		log.Fatalln("缺少sentry dsn")
	}
}

func main() {
	dir := config.ConfigPath
	app.Use(panicCapture())
	if helper.CheckDebug() {
		app.Use(requestLogger.New())
	}
	app.Use(service.NewSentry(service.SentryOptions{}))
	app.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
		MaxAge:           600,
		AllowedMethods:   []string{iris.MethodGet, iris.MethodPost, iris.MethodOptions, iris.MethodHead, iris.MethodDelete, iris.MethodPut},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Authorization"},
	}))
	app.AllowMethods(iris.MethodOptions)
	app.Favicon(dir + "/public/image/favicon.ico")
	app.HandleDir("/css", dir+"/public/css")
	app.HandleDir("/js", dir+"/public/js")
	app.HandleDir("/image", dir+"/public/image")
	app.HandleDir("/static", dir+"/public/static")
	routes.RegisterRoutes(app)
	go func() {
		queue.HandleQueue()
	}()
	iris.RegisterOnInterrupt(func() {
		println("\n服务器关闭...")
		timeout := 5 * time.Second
		ctx, cancel := stdContext.WithTimeout(stdContext.Background(), timeout)
		defer cancel()
		app.Shutdown(ctx)
		database.Db.Close()
		service.Redis.Close()
	})
	if run := app.Run(iris.Addr(":3024"), iris.WithoutInterruptHandler, iris.WithOptimizations); run != nil {
		log.Fatalln(run.Error())
	}
}

/**
获取配置文件路径
*/
func getConfigPath() {
	var configPath string
	flag.StringVar(&configPath, "config", "", "获取配置文件目录")
	flag.Parse()
	if configPath == "" {
		log.Fatalln("缺少配置文件路径")
	}
	config.ConfigPath = configPath
}

/**
捕获博客的panic异常
*/
func panicCapture() context.Handler {
	return func(ctx context.Context) {
		defer func() {
			if err := recover(); err != nil {
				if ctx.IsStopped() {
					return
				}
				var stacktrace string
				for i := 1; ; i++ {
					_, f, l, got := runtime.Caller(i)
					if !got {
						break
					}

					stacktrace += fmt.Sprintf("%s:%d\n", f, l)
				}
				// when stack finishes
				logMessage := fmt.Sprintf("Recovered from a route's Handler('%s')\n", ctx.HandlerName())
				logMessage += fmt.Sprintf("At Request: %s\n", getRequestLogs(ctx))
				logMessage += fmt.Sprintf("Trace: %s\n", err)
				logMessage += fmt.Sprintf("\n%s", stacktrace)
				if !helper.CheckDebug() {
					logMessage = strings.Replace(logMessage, "\n", ",", -1)
				}
				service.Log.Error(logMessage)
				ctx.StatusCode(500)
				ctx.StopExecution()
			}
		}()
		ctx.Next()
	}
}

/**
获取请求数据
*/
func getRequestLogs(ctx context.Context) string {
	var status, ip, method, path string
	status = strconv.Itoa(ctx.GetStatusCode())
	path = ctx.Path()
	method = ctx.Method()
	ip = ctx.RemoteAddr()
	// the date should be logged by iris' Logger, so we skip them
	return fmt.Sprintf("%v %s %s %s", status, path, method, ip)
}
