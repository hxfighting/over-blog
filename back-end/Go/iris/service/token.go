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
		ErrorHandler: func(ctx iris.Context, s string) {
			if s == "Token is expired" {
				token, err := renewalToken(ctx)
				if err != nil {
					response.RenderError(ctx, err.Error(), nil)
					return
				}
				ctx.Header(JWT_KEY, token)
				ctx.Next()
			} else {
				response.RenderError(ctx, s, nil)
			}
		},
	})
}

/**
生成token
*/
func GenerateToken(user_id uint, exp_end int64) (string, *jwt.Token, error) {
	if exp_end == 0 {
		exp_end = time.Now().Unix() + 86400*7
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo":     "bar",
		"nbf":     time.Now().Unix(),
		"exp":     time.Now().Add(time.Second * 60).Unix(),
		"id":      float64(user_id),
		"exp_end": time.Now().Unix() + 86400*7,
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(config.GetConfig("jwt.secret").(string)))
	if err != nil {
		Log.Error(err.Error())
		return "", nil, errors.New("token生成失败！")
	}
	tokenString = "Bearer " + tokenString
	if cacheToken(user_id, tokenString) {
		return tokenString, token, nil
	}
	return "", &jwt.Token{}, errors.New("token缓存失败！")
}

/**
缓存token
*/
func cacheToken(user_id uint, token string) bool {
	val := md5.Sum([]byte(token))
	set := Redis.Set(fmt.Sprintf("%x", val), user_id, time.Second*120)
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
一周内续签token
*/
func renewalToken(ctx iris.Context) (string, error) {
	userIdStr, err := getCacheToken(ctx)
	if err != nil || userIdStr == "" {
		return "", errors.New("Token已过期，请重新登录！")
	}
	s, e := jwtmiddleware.FromAuthHeader(ctx)
	if e != nil {
		Log.Error(e.Error())
		return "", errors.New("Token已过期，请重新登录！")
	}
	my_claims := jwt.MapClaims{
		"foo":     "",
		"nbf":     "",
		"exp":     "",
		"id":      "",
		"exp_end": "",
	}
	_, _, e = new(jwt.Parser).ParseUnverified(s, &my_claims)
	if e != nil {
		Log.Error(e.Error())
		return "", errors.New("Token已过期，请重新登录！")
	}
	exp_end := int64(my_claims["exp_end"].(float64))
	if exp_end < time.Now().Unix() {
		return "", errors.New("Token已过期，请重新登录！")
	}
	new_token, parseToken, e := GenerateToken(uint(my_claims["id"].(float64)), exp_end)
	if e != nil {
		return "", errors.New("Token已过期，请重新登录！")
	}
	token := "Bearer " + s
	key := md5.Sum([]byte(token))
	res := Redis.Del(fmt.Sprintf("%x", key))
	_, e = res.Result()
	if e != nil {
		return "", errors.New("Token已过期，请重新登录！")
	}
	ctx.Values().Set(JWT_KEY, parseToken)
	return new_token, nil
}
