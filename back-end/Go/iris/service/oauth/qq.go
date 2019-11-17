package oauth

import (
	"blog/helper"
	"errors"
	"github.com/tidwall/gjson"
	"net/url"
	"regexp"
	"strings"
)

/**
获取qq授权连接
*/
func (this OauthConfig) GetOauthUrlFromQQ(state string) string {
	v := url.Values{}
	v.Set("redirect_uri", this.Callback)
	return "https://graph.qq.com/oauth2.0/authorize?response_type=code&client_id=" +
		this.ClientID + "&" + v.Encode() + "&state=" + state
}

/**
获取qq的access_token
*/
func (this OauthConfig) GetAccessTokenFromQQ(code string) (string, error) {
	v := url.Values{}
	v.Set("redirect_uri", this.Callback)
	urll := "https://graph.qq.com/oauth2.0/token?grant_type=authorization_code&client_id=" +
		this.ClientID + "&client_secret=" + this.Secret + "&code=" + code + "&" + v.Encode()
	body, _, e := helper.GetHttpResponse(urll, "GET", nil)
	if e != nil {
		return "", e
	}
	body_str := string(body)
	err_match, _ := regexp.MatchString("error", body_str)
	if err_match {
		re, _ := regexp.Compile("({.*})")
		err_slice := re.FindStringSubmatch(body_str)
		err_str := err_slice[0]
		msg := gjson.Get(err_str, "error_description").String()
		return "", errors.New(msg)
	} else {
		re, _ := regexp.Compile("access_token=(.*)&expires_in")
		token_slice := re.FindStringSubmatch(body_str)
		if len(token_slice) != 2 {
			msg := "获取access_token失败" + body_str
			return "", errors.New(msg)
		} else {
			return token_slice[1], nil
		}
	}
}

/**
获取qq的openid
*/
func (this OauthConfig) GetOpenIDFromQQ(access_token string) (string, error) {
	urll := "https://graph.qq.com/oauth2.0/me?access_token=" + access_token
	body, _, e := helper.GetHttpResponse(urll, "GET", nil)
	if e != nil {
		return "", e
	}
	body_str := string(body)
	err_match, _ := regexp.MatchString("error", body_str)
	if err_match {
		re, _ := regexp.Compile("({.*})")
		err_slice := re.FindStringSubmatch(body_str)
		err_str := err_slice[0]
		msg := gjson.Get(err_str, "error_description").String()
		return "", errors.New(msg)
	} else {
		openid := gjson.Get(body_str, "openid").String()
		if openid == "" {
			msg := gjson.Get(body_str, "msg").String()
			if msg == "" {
				msg = "获取openid失败"
			}
			return "", errors.New(msg)
		}
		return openid, nil
	}
}

/**
获取用户信息
*/
func (this OauthConfig) GetUserFromQQ(code string) (map[string]string, error) {
	user := make(map[string]string)
	access_token, e := this.GetAccessTokenFromQQ(code)
	if e != nil {
		return user, e
	}
	openid, e := this.GetOpenIDFromQQ(access_token)
	if e != nil {
		return user, e
	}
	urll := "https://graph.qq.com/user/get_user_info?access_token=" + access_token +
		"&oauth_consumer_key=" + this.ClientID + "&openid=" + openid
	body, _, e := helper.GetHttpResponse(urll, "GET", nil)
	if e != nil {
		return user, e
	}
	body_str := string(body)
	err_match, _ := regexp.MatchString("error", body_str)
	if err_match {
		re, _ := regexp.Compile("({.*})")
		err_slice := re.FindStringSubmatch(body_str)
		err_str := err_slice[0]
		msg := gjson.Get(err_str, "error_description").String()
		return user, errors.New(msg)
	} else {
		err := gjson.Get(body_str, "ret").Int()
		if err != 0 {
			msg := gjson.Get(body_str, "msg").String()
			return user, errors.New(msg)
		}
		user["access_token"] = access_token
		user["openid"] = openid
		user["name"] = gjson.Get(body_str, "nickname").String()
		avatar := gjson.Get(body_str, "figureurl_qq_1").String()
		if strings.Index(avatar, "https") == -1 {
			avatar = strings.Replace(avatar, "http", "https", -1)
		}
		user["avatar"] = avatar
		return user, nil
	}

}
