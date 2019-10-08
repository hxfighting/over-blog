package backend

import (
	"blog/service"
	"errors"
	"github.com/kataras/iris"
	"github.com/tidwall/gjson"
	"strconv"
)

var response = service.Response{}

/**
生成验证码
*/
func GetCaptcha(ctx iris.Context) {
	key, captchaBase64 := service.GenerateCaptcha()
	data := make(map[string]string)
	data["img"] = captchaBase64
	data["key"] = key
	response.RenderSuccess(ctx, "登录成功！", data)
}

/**
获取需要的数据,并且验证数据
*/
func getRequestData(ctx iris.Context, validates []service.BlogValidate) (map[string]interface{}, error) {
	jsonBytes, _ := ctx.GetBody()
	jsonData := string(jsonBytes)
	if !gjson.Valid(jsonData) {
		return nil, errors.New("json格式错误!")
	}
	mapData := make(map[string]interface{})
	for _, validate := range validates {
		data := gjson.Get(jsonData, validate.Key).Value()
		e := service.Validate.Var(data, validate.Validation)
		if e != nil {
			return nil, errors.New(validate.Err)
		}
		mapData[validate.Key] = getRightData(validate.Type, data.(string))
	}
	return mapData, nil
}

/**
获取正确类型数据
*/
func getRightData(types, data string) interface{} {
	switch types {
	case "float64":
		f, e := strconv.ParseFloat(data, 64)
		if e != nil {
			return data
		}
		return f
	case "int64":
		i, err := strconv.ParseInt(data, 10, 64)
		if err != nil {
			return data
		}
		return i
	case "uint":
		u, e := strconv.ParseUint(data, 10, 0)
		if e != nil {
			return data
		}
		return u
	case "int":
		u, e := strconv.ParseInt(data, 10, 0)
		if e != nil {
			return data
		}
		return u
	default:
		return data
	}
}
