package wechat

import (
	"blog/helper"
	"blog/service"
	"bytes"
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"strings"
)

//小程序access_token缓存key
const access_token_key = "miniProgram_access_token"

type QRCoder struct {
	// page 必须是已经发布的小程序存在的页面,根路径前不要填加 /,不能携带参数（参数请放在scene字段里），如果不填写这个字段，默认跳主页面
	Page string `json:"page,omitempty"`
	// path 扫码进入的小程序页面路径
	Path string `json:"path,omitempty"`
	// width 图片宽度
	Width int `json:"width,omitempty"`
	// scene 最大32个可见字符，只支持数字，大小写英文以及部分特殊字符：!#$&'()*+,/:;=?@-._~，其它字符请自行编码为合法字符（因不支持%，中文无法使用 urlencode 处理，请使用其他编码方式）
	Scene string `json:"scene,omitempty"`
	// autoColor 自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调
	AutoColor bool `json:"auto_color,omitempty"`
	// lineColor AutoColor 为 false 时生效，使用 rgb 设置颜色 例如 {"r":"xxx","g":"xxx","b":"xxx"},十进制表示
	LineColor Color `json:"line_color,omitempty"`
	// isHyaline 是否需要透明底色
	IsHyaline bool `json:"is_hyaline,omitempty"`
}

type Color struct {
	R string `json:"r"`
	G string `json:"g"`
	B string `json:"b"`
}

/**
获取小程序access_token
*/
func (this Config) GetAccessToken() (access_token string, err error) {
	urll := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + this.ClientID +
		"&secret=" + this.Secret
	body, _, err := helper.GetHttpResponse(urll, "GET", nil)
	if err != nil {
		return
	}
	body_str := string(body)
	error_code := gjson.Get(body_str, "errcode").Int()
	if error_code != 0 {
		err_msg := gjson.Get(body_str, "errmsg").String()
		err = errors.New(err_msg)
		return
	}
	access_token = gjson.Get(body_str, "access_token").String()
	if access_token == "" {
		err = errors.New("缺少access_token字段")
		return
	}
	return access_token, nil
}

/**
获取二维码
*/
func (this Config) GetQrCode(scene string, coder QRCoder) (response []byte, err error) {
	access_token, err := this.GetAccessToken()
	if err != nil {
		return
	}
	urll := "https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=" + access_token
	coder_byte, err := service.FastJson.Marshal(coder)
	if err != nil {
		return
	}
	response, contentType, err := helper.GetHttpResponse(urll, "POST", bytes.NewReader(coder_byte))
	if strings.HasPrefix(contentType, "application/json") {
		// 返回错误信息
		response_str := string(response)
		err_msg := gjson.Get(response_str, "errmsg").String()
		err = errors.New(err_msg)
		return
	} else if contentType == "image/jpeg" {
		// 返回文件
		return response, nil
	} else {
		err = fmt.Errorf("fetchCode error : unknown response content type - %v", contentType)
		return nil, err
	}
}

/**
获取openID
*/
func (this Config) GetOpenID(jsCode string) (openid, sessionKey string, err error) {
	urll := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	urll = fmt.Sprintf(urll, this.ClientID, this.Secret, jsCode)
	response, _, err := helper.GetHttpResponse(urll, "GET", nil)
	if err != nil {
		return
	}
	response_str := string(response)
	errcode := gjson.Get(response_str, "errcode").Int()
	if errcode != 0 {
		err_msg := gjson.Get(response_str, "errmsg").String()
		err = errors.New(err_msg)
		return
	}
	openid = gjson.Get(response_str, "openid").String()
	sessionKey = gjson.Get(response_str, "session_key").String()
	return
}
