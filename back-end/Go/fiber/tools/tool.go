package tools

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/sync/singleflight"

	"github.com/ohdata/blog/tools/log"
)

var (
	ErrServer = errors.New(http.StatusText(http.StatusInternalServerError))
	Validate  = validator.New()
	OnceDo    = new(singleflight.Group)
)

func baseResponse(ctx *fiber.Ctx, code int, msg string, data ...interface{}) error {
	var d interface{}
	if len(data) > 0 {
		d = data[0]
	}
	return ctx.JSON(fiber.Map{"code": code, "msg": msg, "data": d})
}

func Success(ctx *fiber.Ctx, msg string, data ...interface{}) error {
	return baseResponse(ctx, http.StatusOK, msg, data...)
}

func Error(ctx *fiber.Ctx, msg string, data ...interface{}) error {
	return baseResponse(ctx, -1, msg, data...)
}

func ValidateStruct(ctx *fiber.Ctx, data interface{}) error {
	if err := ctx.BodyParser(data); err != nil {
		return err
	}
	return Validate.Struct(data)
}

func SingleRecordResponse(err error, annotate string) error {
	if err == sql.ErrNoRows {
		return errors.New(annotate)
	}
	return ErrServer
}

func ServerErrResponse(ctx *fiber.Ctx, err error) error {
	log.Log.Err(err).Send()
	return Error(ctx, ErrServer.Error())
}
