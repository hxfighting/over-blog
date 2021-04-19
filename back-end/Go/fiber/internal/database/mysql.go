package database

import (
	"database/sql"
	"fmt"
	stdLog "log"
	"os"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/ohdata/blog/configs"
	"github.com/ohdata/blog/tools/log"
	"github.com/ohdata/blog/tools/util"
)

var (
	DB   *gorm.DB
	once sync.Once
)

func New() (err error) {
	cfg := configs.Config.Mysql
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4", cfg.Username, cfg.Password, cfg.Host, cfg.DB)
	once.Do(func() {
		err = util.Retry(func() error {
			d, err := sql.Open("mysql", dns)
			if err != nil {
				return err
			}
			d.SetMaxIdleConns(cfg.MaxIdleConns)
			d.SetMaxOpenConns(cfg.MaxOpenConns)
			d.SetConnMaxLifetime(cfg.MaxLifeTime)
			d.SetConnMaxIdleTime(cfg.MaxIdleTime)
			err = d.Ping()
			if err != nil {
				return err
			}
			level := logger.Error
			if configs.Config.Server.Debug {
				level = logger.Info
			}
			newLogger := logger.New(
				stdLog.New(os.Stdout, "\r\n", stdLog.LstdFlags), // io writer
				logger.Config{
					SlowThreshold: time.Second * 5,             // 慢 SQL 阈值
					LogLevel:      level,                       // Log level
					Colorful:      configs.Config.Server.Debug, // 禁用彩色打印
				},
			)
			db, err := gorm.Open(mysql.New(mysql.Config{
				DefaultStringSize:         256,
				DisableDatetimePrecision:  true,
				DontSupportRenameIndex:    true,
				DontSupportRenameColumn:   true,
				SkipInitializeWithVersion: false,
				Conn:                      d,
			}), &gorm.Config{
				SkipDefaultTransaction: true,
				PrepareStmt:            true,
				Logger:                 newLogger,
			})
			if err != nil {
				d.Close()
				return err
			}
			DB = db
			return nil
		}, 3, time.Second*3)
	})
	return
}

func Close() {
	if DB != nil {
		d, err := DB.DB()
		if err != nil {
			log.Log.Err(err).Send()
		}
		if err := d.Close(); err != nil {
			log.Log.Err(err).Send()
		}
	}
}
