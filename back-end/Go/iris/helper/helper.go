package helper

import (
	"blog/config"
	"errors"
	"github.com/kataras/iris/v12"
	"io"
	"io/ioutil"
	"log"
	"net/http"
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
时间转时间戳
*/
func GetUnixTimeFromDate(date, format string) (int64, error) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return 0, errors.New("时区设置失败")
	}
	tt, err := time.ParseInLocation(format, date, loc)
	if err != nil {
		return 0, errors.New("时间转换失败！")
	}
	return tt.Unix(), nil
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

/**
获取http请求结果
*/
func GetHttpResponse(http_url, method string, body io.Reader) ([]byte, string, error) {
	res := []byte{}
	request, e := http.NewRequest(method, http_url, body)
	if e != nil {
		return res, "", e
	}
	response, e := http.DefaultClient.Do(request)
	if e != nil {
		return res, "", e
	}
	defer response.Body.Close()
	body_byte, e := ioutil.ReadAll(response.Body)
	if e != nil {
		return res, "", e
	}
	contentType := response.Header.Get("Content-Type")
	return body_byte, contentType, nil
}

/**
获取今日23:59:59
*/
func GetTimeRemainingToday() (time.Time, error) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return time.Time{}, errors.New("时区设置失败")
	}
	tt, err := time.ParseInLocation(YMDHIS, time.Now().Format(YMD)+" 23:59:59", loc)
	if err != nil {
		return time.Time{}, errors.New("时间转换失败！")
	}
	return tt, nil
}

/**
获取两个日期相差多少天
dayFirst Y-m-d
dayLast Y-m-d
*/
func GetDateDiffDay(dayFirst, dayLast string) (int, error) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return 0, errors.New("时区设置失败")
	}
	tt, err := time.ParseInLocation(YMD, dayFirst, loc)
	if err != nil {
		return 0, errors.New("时间转换失败！")
	}
	tt_two, err := time.ParseInLocation(YMD, dayLast, loc)
	if err != nil {
		return 0, errors.New("时间转换失败！")
	}
	hours := tt_two.Sub(tt).Hours()
	if hours <= 0 {
		return 0, errors.New("时间错误")
	}
	if hours < 24 {
		// may same day
		t1y, t1m, t1d := tt.Date()
		t2y, t2m, t2d := tt_two.Date()
		isSameDay := (t1y == t2y && t1m == t2m && t1d == t2d)
		if isSameDay {
			return 0, errors.New("时间错误")
		} else {
			return 1, nil
		}
	} else {
		if (hours/24)-float64(int(hours/24)) == 0 {
			return int(hours / 24), nil
		} else {
			return int(hours/24) + 1, nil
		}
	}
}

/**
反转义html
*/
func DecodeHtml(str string) string {
	var s = ""
	if len(str) <= 0 {
		return ""
	}
	s = strings.Replace(str, "&amp;", "&", -1)
	s = strings.Replace(s, "&lt;", "<", -1)
	s = strings.Replace(s, "&gt;", ">", -1)
	s = strings.Replace(s, "&#39;", "'", -1)
	s = strings.Replace(s, "&quot;", "\"", -1)
	s = strings.Replace(s, "&huhu;", "\n", -1)
	return s
}
