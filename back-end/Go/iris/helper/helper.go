package helper

import (
	"blog/config"
	"errors"
	"github.com/kataras/iris"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	YMDHIS = "2006-01-02 15:04:05"
	YMD    = "2006-01-02"
)

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

/**
验证电话号码
*/
func VerifyMobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

/**
检查是否调试
*/
func CheckDebug() bool {
	var res bool
	var err error
	debugStr := config.GetConfig("app.debug").(string)
	if debugStr != "" {
		res, err = strconv.ParseBool(debugStr)
		if err != nil {
			res = false
		}
	}
	return res
}

/**
格式化时间戳
*/
func GetDateTime(unix int64, format string) string {
	return time.Unix(unix, 0).Format(format)
}

/**
获取post、put的json请求数据
*/
func GetRequestData(ctx iris.Context) (map[string]interface{}, error) {
	var request interface{}
	err := ctx.ReadJSON(&request)
	if err != nil {
		return map[string]interface{}{}, err
	}
	request_values, ok := request.(map[string]interface{})
	if !ok {
		return map[string]interface{}{}, errors.New("数据格式错误！")
	}
	return request_values, nil
}
