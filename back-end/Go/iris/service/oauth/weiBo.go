package oauth

import (
	"blog/helper"
	"errors"
	"github.com/tidwall/gjson"
	"net/url"
	"strings"
)

/**
获取微博授权URL
*/
func (this OauthConfig) GetOauthUrlFromWeiBo() string {
	v := url.Values{}
	v.Set("redirect_uri", this.Callback)
	return "https://api.weibo.com/oauth2/authorize?client_id=" + this.ClientID + "&response_type=code&" + v.Encode()
}

/**
获取微博access_token
*/
func (this OauthConfig) GetAccessTokenFromWeiBo(code string) (string, error) {
	v := url.Values{}
	v.Set("redirect_uri", this.Callback)
	urll := "https://api.weibo.com/oauth2/access_token?client_id=" + this.ClientID + "&client_secret=" +
		this.Secret + "&grant_type=authorization_code&" + v.Encode() + "&code=CODE"
	s_byte, _, e := helper.GetHttpResponse(urll, "GET", nil)
	if e != nil {
		return "", e
	}
	s := string(s_byte)
	access_token := gjson.Get(s, "access_token").String()
	if access_token == "" {
		msg := gjson.Get(s, "error").String()
		if msg == "" {
			msg = "获取access_token失败"
		}
		return "", errors.New(msg)
	}
	return access_token, nil
}

/**
获取微博openID
*/
func (this OauthConfig) GetOpenIDFromWeiBo(access_token string) (string, error) {
	urll := "https://api.weibo.com/2/account/get_uid.json?access_token=" + access_token
	s_byte, _, e := helper.GetHttpResponse(urll, "GET", nil)
	if e != nil {
		return "", e
	}
	s := string(s_byte)
	openid := gjson.Get(s, "uid").String()
	if openid == "" {
		msg := gjson.Get(s, "error").String()
		if msg == "" {
			msg = "获取openid失败"
		}
		return "", errors.New(msg)
	}
	return openid, nil
}

/**
获取用户信息
*/
func (this OauthConfig) GetUserFromWeiBo(code string) (map[string]string, error) {
	user := make(map[string]string)
	access_token, e := this.GetAccessTokenFromWeiBo(code)
	if e != nil {
		return user, e
	}
	openid, e := this.GetOpenIDFromWeiBo(access_token)
	if e != nil {
		return user, e
	}
	urll := "https://api.weibo.com/2/users/show.json?access_token=" + access_token + "&uid=" + openid
	body, _, e := helper.GetHttpResponse(urll, "GET", nil)
	if e != nil {
		return user, e
	}
	body_str := string(body)
	err := gjson.Get(body_str, "error").String()
	if err != "" {
		return user, errors.New(err)
	}
	user["access_token"] = access_token
	user["openid"] = openid
	user["name"] = gjson.Get(body_str, "name").String()
	user["avatar"] = strings.Replace(gjson.Get(body_str, "profile_image_url").String(), "http", "https", -1)

	return user, nil
}
