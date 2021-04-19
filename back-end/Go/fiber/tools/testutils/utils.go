package testutils

import (
	"database/sql"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"

	"github.com/ohdata/blog/internal/database"
)

var router *fiber.App

// SetupRouter register route
func SetupRouter(route func(app *fiber.App)) {
	router = fiber.New(fiber.Config{
		CaseSensitive:         true,
		Immutable:             true,
		DisableStartupMessage: true,
	})
	route(router)
}

// GetRouter ...
func GetRouter() *fiber.App {
	return router
}

// SetCommonObject ...
func SetCommonObject(db *sql.DB) {
	database.DB = sqlx.NewDb(db, "mysql")
}

func MockRequest(method, target string, body io.Reader) *http.Request {
	req := httptest.NewRequest(method, target, body)
	req.Header.Set("Content-Type", "application/json")
	return req
}
