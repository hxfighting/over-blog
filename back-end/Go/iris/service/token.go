package service

import (
	"blog/config"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
	"time"
)

var response Response

const (
	JWT_KEY = "Authorization"
)

/**
注册中间件
*/
func GetJWTHandler() *jwtmiddleware.Middleware {
	return jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetConfig("jwt.secret").(string)), nil
		},
		// When set, the middleware verifies that tokens are signed with the specific signing algorithm
		// If the signing method is not constant the ValidationKeyGetter callback can be used to implement additional checks
		// Important to avoid security issues described here: https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
		SigningMethod: jwt.SigningMethodHS256,
		Expiration:    true,
		ContextKey:    JWT_KEY,
		ErrorHandler: func(ctx iris.Context, s error) {
			if s.Error() == "Token is expired" {
				ctx.Next()
				//response.RenderError(ctx, "token已过期", nil)
				return
			} else {
				response.RenderError(ctx, "非法token", nil)
				return
			}
		},
	})
}

/**
生成token
*/
func GenerateToken(user_id uint, exp_end int64) (map[string]interface{}, error) {
	if exp_end == 0 {
		exp_end = time.Now().Unix() + 86400*7
	}
	exp := time.Now().Add(time.Minute * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo":     "bar",
		"nbf":     time.Now().Unix(),
		"exp":     exp,
		"id":      float64(user_id),
		"exp_end": exp_end,
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(config.GetConfig("jwt.secret").(string)))
	if err != nil {
		Log.Error(err.Error())
		return nil, errors.New("token生成失败！")
	}
	tokenString = "Bearer " + tokenString
	if cacheToken(user_id, tokenString) {
		data := map[string]interface{}{
			"token":  tokenString,
			"expire": exp,
		}
		return data, nil
	}
	return nil, errors.New("token缓存失败！")
}

/**
缓存token
*/
func cacheToken(user_id uint, token string) bool {
	val := md5.Sum([]byte(token))
	set := Redis.Set(fmt.Sprintf("%x", val), user_id, time.Minute*1)
	_, e := set.Result()
	if e != nil {
		return false
	}
	return true
}

/**
获取缓存的token
*/
func getCacheToken(ctx iris.Context) (string, error) {
	authHeader := ctx.GetHeader("Authorization")
	key := md5.Sum([]byte(authHeader))
	userIdStr := Redis.Get(fmt.Sprintf("%x", key))
	return userIdStr.Result()
}

/**
获取用户ID
*/
func GetUserId(ctx iris.Context) uint {
	claims := getClaims(ctx)
	if val, ok := claims["id"]; ok {
		id := val.(float64)
		return uint(id)
	}
	return 0
}

func getClaims(ctx iris.Context) jwt.MapClaims {
	token := ctx.Values().Get(JWT_KEY).(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	return claims
}

/**
刷新token
*/
func RefreshToken(ctx iris.Context) (map[string]interface{}, error) {
	s, e := jwtmiddleware.FromAuthHeader(ctx)
	if e != nil {
		return nil, errors.New("缺少token")
	}
	_, e = new(jwt.Parser).Parse(s, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(config.GetConfig("jwt.secret").(string)), nil
	})
	if e != nil && e.Error() != "Token is expired" {
		return nil, errors.New("非法token")
	}
	key := fmt.Sprintf("%x", md5.Sum([]byte("Bearer "+s)))
	exist_black_list := Redis.SIsMember("token_black_list", key)
	if exist_black_list.Val() {
		return nil, errors.New("非法token")
	}
	my_claims := jwt.MapClaims{
		"foo":     "",
		"nbf":     "",
		"exp":     "",
		"id":      "",
		"exp_end": "",
	}
	_, _, _ = new(jwt.Parser).ParseUnverified(s, &my_claims)
	exp_float := my_claims["exp"].(float64)
	exp := int64(exp_float)
	now_unix := time.Now().Unix()
	exp_end := int64(my_claims["exp_end"].(float64))
	if now_unix-exp > 86400 || exp_end < now_unix {
		return nil, errors.New("Token已过期，请重新登录！")
	}
	new_token, e := GenerateToken(uint(my_claims["id"].(float64)), exp_end)
	if e != nil {
		return nil, errors.New("Token已过期，请重新登录！")
	}
	res := Redis.SAdd("token_black_list", key)
	if res.Err() != nil {
		return nil, errors.New("Token已过期，请重新登录！")
	}
	return new_token, nil
}
