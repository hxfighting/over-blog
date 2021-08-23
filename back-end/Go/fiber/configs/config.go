package configs

import (
	"bytes"
	_ "embed"
	"time"

	"github.com/spf13/viper"
)

type GlobalConfig struct {
	Server struct {
		Addr          string        `yaml:"addr" json:"addr" cfg:"addr"`
		Debug         bool          `yaml:"debug" json:"debug" cfg:"debug"`
		EnableSwagger bool          `yaml:"enableswagger" json:"enableswagger" cfg:"enableswagger"`
		PProfToken    string        `yaml:"pprofToken" json:"pprofToken" cfg:"pprofToken"`
		Name          string        `yaml:"name" json:"name" cfg:"name"`
		MaxBody       int           `yaml:"maxbody" json:"maxbody" cfg:"maxbody"`
		Concurrency   int           `yaml:"concurrency" json:"concurrency" cfg:"concurrency"`
		ReadTimeout   time.Duration `yaml:"readtimeout" json:"readtimeout" cfg:"readtimeout"`
		WriteTimeout  time.Duration `yaml:"writetimeout" json:"writetimeout" cfg:"writetimeout"`
	} `yaml:"server" json:"server" cfg:"server"`
	Mysql struct {
		MaxOpenConns int           `yaml:"maxopenconns" json:"maxopenconns" cfg:"maxopenconns"`
		MaxIdleConns int           `yaml:"maxidleconns" json:"maxidleconns" cfg:"maxidleconns"`
		MaxLifeTime  time.Duration `yaml:"maxlifetime" json:"maxlifetime" cfg:"maxlifetime"`
		MaxIdleTime  time.Duration `yaml:"maxidletime" json:"maxidletime" cfg:"maxidletime"`
		Username     string        `yaml:"username" json:"username" cfg:"username"`
		Password     string        `yaml:"password" json:"password" cfg:"password"`
		Host         string        `yaml:"host" json:"host" cfg:"host"`
		DB           string        `yaml:"db" json:"db" cfg:"db"`
	} `yaml:"mysql" json:"mysql" cfg:"mysql"`
	CORS struct {
		AllowOrigins     string `json:"alloworigins" yaml:"alloworigins"`
		AllowMethods     string `json:"allowmethods" yaml:"allowmethods"`
		AllowHeaders     string `json:"allowheaders" yaml:"allowheaders"`
		ExposeHeaders    string `json:"exposeheaders" yaml:"exposeheaders"`
		AllowCredentials bool   `json:"allowcredentials" yaml:"allowcredentials"`
		MaxAge           int    `json:"maxage" yaml:"maxage"`
		Enable           bool   `json:"enable" yaml:"enable"`
	} `yaml:"cors" json:"cors"`
	RateLimit struct {
		Max    int           `json:"max" yaml:"max"`
		Expire time.Duration `json:"expire" yaml:"expire"`
	} `yaml:"ratelimit" json:"ratelimit"`
	Log struct {
		Path   string `json:"path" yaml:"path"`
		Output string `json:"output" yaml:"output"`
	} `yaml:"log" json:"log"`
	JWT struct {
		Secret           string        `json:"secret" yaml:"secret" cfg:"secret"`
		Expire           time.Duration `json:"expire" yaml:"expire"`
		MaxRefresh       int           `json:"maxrefresh" yaml:"maxrefresh"`
		SigningAlgorithm string        `json:"signingalgorithm" yaml:"signingalgorithm"`
	} `yaml:"jwt" json:"jwt" cfg:"jwt"`
	HTTPClient struct {
		RequestTimeout            time.Duration `yaml:"requesttimeout"`
		ReadTimeout               time.Duration `yaml:"readtimeout"`
		WriteTimeout              time.Duration `yaml:"writetimeout"`
		MaxConnDuration           time.Duration `yaml:"maxconnduration"`
		MaxConnWaitTimeout        time.Duration `yaml:"maxconnwaittimeout"`
		MaxConnsPerHost           int           `yaml:"maxconnsperhost"`
		MaxIdleConnDuration       time.Duration `yaml:"maxidleconnduration"`
		MaxIdemponentCallAttempts int           `yaml:"maxidemponentcallattempts"`
	} `yaml:"httpclient" json:"httpclient"`
	QiNiu struct {
		AccessKey string `json:"accesskey" yaml:"accesskey" cfg:"accesskey"`
		SecretKey string `json:"secretkey" yaml:"secretkey" cfg:"secretkey"`
		Bucket    string `json:"bucket" yaml:"bucket" cfg:"bucket"`
	} `yaml:"qiniu" json:"qiniu"`
	Redis struct {
		Host         string        `json:"host" yaml:"host"`
		Password     string        `json:"password" yaml:"password"`
		Port         string        `json:"port" yaml:"port"`
		DB           int           `json:"db" yaml:"db"`
		MinIdleConns int           `json:"minidleconns" yaml:"minidleconns"`
		PoolSize     int           `json:"poolsize" yaml:"poolsize"`
		MaxConnAge   time.Duration `json:"maxconnage" yaml:"maxconnage"`
		ReadTimeout  time.Duration `json:"readtimeout" yaml:"readtimeout"`
		WriteTimeout time.Duration `json:"writetimeout" yaml:"writetimeout"`
		DialTimeout  time.Duration `json:"dialtimeout" yaml:"dialtimeout"`
	} `yaml:"redis" json:"redis" cfg:"redis"`
	Event struct {
		Size       int `json:"size" yaml:"size"`
		Concurrent int `json:"concurrent" yaml:"concurrent"`
	} `yaml:"event" json:"event" cfg:"event"`
}

var Config *GlobalConfig

//go:embed config.yaml
var c []byte

func New() error {
	viper.SetConfigType("yaml")
	if err := viper.ReadConfig(bytes.NewBuffer(c)); err != nil {
		return err
	}
	return viper.Unmarshal(&Config)
}
