package service

import (
	"gopkg.in/go-playground/validator.v9"
)

var (
	Validate = newInstance()
)

func newInstance() *validator.Validate {
	validate := validator.New()
	return validate
}
