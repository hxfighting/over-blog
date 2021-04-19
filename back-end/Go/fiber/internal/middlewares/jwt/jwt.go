package jwt

import (
	stdJson "encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/kataras/jwt"

	"github.com/ohdata/blog/tools"
	"github.com/ohdata/blog/tools/util"
)

// Config defines the config for BasicAuth middleware
type Config struct {
	// Next defines a function to skip middleware.
	// Optional. Default: nil
	Next func(*fiber.Ctx) bool

	// SuccessHandler defines a function which is executed for a valid token.
	// Optional. Default: nil
	SuccessHandler fiber.Handler

	// ErrorHandler defines a function which is executed for an invalid token.
	// It may be used to define a custom JWT error.
	// Optional. Default: 401 Invalid or expired JWT
	ErrorHandler fiber.ErrorHandler

	// Signing key to validate token. Used as fallback if SigningKeys has length 0.
	// Required.
	SigningKey string

	// Signing method, used to check token signing method.
	// Optional. Default: "HS256".
	// Possible values: "HS256", "HS384", "HS512", "ES256", "ES384", "ES512", "RS256", "RS384", "RS512", "EdDSA"
	SigningMethod string

	// Context key to store user information from the token into context.
	// Optional. Default: "user".
	ContextKey string

	// TokenLookup is a string in the form of "<source>:<name>" that is used
	// to extract token from the request.
	// Optional. Default value "header:Authorization".
	// Possible values:
	// - "header:<name>"
	// - "query:<name>"
	// - "param:<name>"
	// - "cookie:<name>"
	TokenLookup string

	// AuthScheme to be used in the Authorization header.
	// Optional. Default: "Bearer".
	AuthScheme string

	// Expire is the token expire time
	Expire time.Duration

	alg jwt.Alg

	privateKey interface{}
	publicKey  interface{}
	blockList  *jwt.Blocklist
	extractors []func(c *fiber.Ctx) (string, error)
}

var cfg Config

// New ...
func New(config ...Config) fiber.Handler {
	// Init config
	if len(config) > 0 {
		cfg = config[0]
	}
	if cfg.SuccessHandler == nil {
		cfg.SuccessHandler = func(c *fiber.Ctx) error {
			return c.Next()
		}
	}
	if cfg.ErrorHandler == nil {
		cfg.ErrorHandler = func(c *fiber.Ctx, err error) error {
			if err.Error() == "Missing or malformed JWT" {
				return c.Status(fiber.StatusBadRequest).SendString("Missing or malformed JWT")
			}
			return c.Status(fiber.StatusUnauthorized).SendString("Invalid or expired JWT")
		}
	}
	if cfg.SigningKey == "" {
		panic("Fiber: JWT middleware requires signing key")
	}
	if cfg.SigningMethod == "" {
		cfg.SigningMethod = "HS256"
	}
	if cfg.ContextKey == "" {
		cfg.ContextKey = "user"
	}
	if cfg.TokenLookup == "" {
		cfg.TokenLookup = "header:" + fiber.HeaderAuthorization
	}
	if cfg.AuthScheme == "" {
		cfg.AuthScheme = "Bearer"
	}
	cfg.alg = getSignMethod(cfg.SigningMethod)
	cfg.privateKey, cfg.publicKey = getPrivateKeyAndPublicKey(cfg.SigningMethod, cfg.SigningKey)
	// Initialize
	cfg.extractors = make([]func(c *fiber.Ctx) (string, error), 0)
	rootParts := strings.Split(cfg.TokenLookup, ",")
	for _, rootPart := range rootParts {
		parts := strings.Split(strings.TrimSpace(rootPart), ":")

		switch parts[0] {
		case "header":
			cfg.extractors = append(cfg.extractors, jwtFromHeader(parts[1], cfg.AuthScheme))
		case "query":
			cfg.extractors = append(cfg.extractors, jwtFromQuery(parts[1]))
		case "param":
			cfg.extractors = append(cfg.extractors, jwtFromParam(parts[1]))
		case "cookie":
			cfg.extractors = append(cfg.extractors, jwtFromCookie(parts[1]))
		}
	}
	cfg.blockList = jwt.NewBlocklist(time.Hour)
	// Return middleware handler
	return func(c *fiber.Ctx) error {
		// Filter request to skip middleware
		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}
		var auth string
		var err error

		for _, extractor := range cfg.extractors {
			auth, err = extractor(c)
			if auth != "" && err == nil {
				break
			}
		}
		if err != nil {
			return cfg.ErrorHandler(c, err)
		}
		token, err := jwt.Verify(cfg.alg, cfg.publicKey, util.StrToBytes(auth), cfg.blockList)
		if err != nil {
			return cfg.ErrorHandler(c, err)
		}
		claims := jwt.Map{}
		if err = token.Claims(&claims); err != nil {
			return cfg.ErrorHandler(c, err)
		}
		if _, ok := claims["subject"]; !ok {
			claims["subject"] = token.StandardClaims.Subject
		}
		c.Locals(cfg.ContextKey, claims)
		return cfg.SuccessHandler(c)
	}
}

func getSignMethod(alg string) jwt.Alg {
	switch alg {
	case "HS256":
		return jwt.HS256
	case "HS384":
		return jwt.HS384
	case "HS512":
		return jwt.HS512
	case "RS256":
		return jwt.RS256
	case "RS384":
		return jwt.RS384
	case "RS512":
		return jwt.RS512
	case "ES256":
		return jwt.ES256
	case "ES384":
		return jwt.ES384
	case "ES512":
		return jwt.ES512
	case "EdDSA":
		return jwt.EdDSA
	default:
		panic("this signing algorithms not support")
	}
}

func getPrivateKeyAndPublicKey(alg string, pk string) (interface{}, interface{}) {
	k := []byte(pk)
	switch alg {
	case "HS256", "HS384", "HS512":
		return k, k
	case "RS256", "RS384", "RS512":
		pk, err := jwt.ParsePrivateKeyRSA(k)
		if err != nil {
			panic(err)
		}
		return pk, pk.Public()
	case "ES256", "ES384", "ES512":
		pk, err := jwt.ParsePrivateKeyECDSA(k)
		if err != nil {
			panic(err)
		}
		return pk, pk.Public()
	case "EdDSA":
		pk, err := jwt.ParsePrivateKeyEdDSA(k)
		if err != nil {
			panic(err)
		}
		return pk, pk.Public()
	default:
		return nil, nil
	}
}

// jwtFromHeader returns a function that extracts token from the request header.
func jwtFromHeader(header string, authScheme string) func(c *fiber.Ctx) (string, error) {
	return func(c *fiber.Ctx) (string, error) {
		auth := c.Get(header)
		l := len(authScheme)
		if len(auth) > l+1 && auth[:l] == authScheme {
			return auth[l+1:], nil
		}
		return "", errors.New("Missing or malformed JWT")
	}
}

// jwtFromQuery returns a function that extracts token from the query string.
func jwtFromQuery(param string) func(c *fiber.Ctx) (string, error) {
	return func(c *fiber.Ctx) (string, error) {
		token := c.Query(param)
		if token == "" {
			return "", errors.New("Missing or malformed JWT")
		}
		return token, nil
	}
}

// jwtFromParam returns a function that extracts token from the url param string.
func jwtFromParam(param string) func(c *fiber.Ctx) (string, error) {
	return func(c *fiber.Ctx) (string, error) {
		token := c.Params(param)
		if token == "" {
			return "", errors.New("Missing or malformed JWT")
		}
		return token, nil
	}
}

// jwtFromCookie returns a function that extracts token from the named cookie.
func jwtFromCookie(name string) func(c *fiber.Ctx) (string, error) {
	return func(c *fiber.Ctx) (string, error) {
		token := c.Cookies(name)
		if token == "" {
			return "", errors.New("Missing or malformed JWT")
		}
		return token, nil
	}
}

// GenerateToken is method to generate token
func GenerateToken(uid int64, username, role string) (map[string]interface{}, error) {
	expire := jwt.Clock().Add(cfg.Expire).Unix()
	claims := jwt.Map{
		"subject":  uid,
		"role":     role,
		"username": username,
		"jti":      utils.UUID(),
		"exp":      expire,
		"iat":      jwt.Clock().Unix(),
	}
	token, err := jwt.Sign(cfg.alg, cfg.privateKey, claims, jwt.MaxAge(cfg.Expire))
	return map[string]interface{}{
		"token":  "Bearer " + util.BytesToStr(token),
		"expire": expire,
	}, err
}

// BlockToken is the method to invalid token
func BlockToken(ctx *fiber.Ctx) error {
	var auth string
	var err error
	for _, extractor := range cfg.extractors {
		auth, err = extractor(ctx)
		if auth != "" && err == nil {
			break
		}
	}
	if err != nil {
		return err
	}
	token, err := jwt.Verify(cfg.alg, cfg.publicKey, util.StrToBytes(auth), cfg.blockList)
	if err != nil {
		return err
	}
	return cfg.blockList.InvalidateToken(token.Token, token.StandardClaims)
}

// GetUserInfo ...
func GetUserInfo(ctx *fiber.Ctx) (map[string]interface{}, error) {
	claims := ctx.Locals(cfg.ContextKey)
	if claims == nil {
		return nil, tools.ErrServer
	}
	c, ok := claims.(jwt.Map)
	if !ok {
		return nil, tools.ErrServer
	}
	return c, nil
}

// GetUID ...
func GetUID(ctx *fiber.Ctx) (int64, error) {
	info, err := GetUserInfo(ctx)
	if err != nil {
		return 0, err
	}
	uidJson := info["subject"].(stdJson.Number)
	uid, err := uidJson.Int64()
	if err != nil {
		return 0, tools.ErrServer
	}
	return uid, nil
}
