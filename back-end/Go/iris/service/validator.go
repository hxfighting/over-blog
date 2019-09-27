package service

import (
	"errors"
	"gopkg.in/go-playground/validator.v9"
)

var (
	Validate = newInstance()
)

type BlogValidate struct {
	Key        string
	Validation string
	Err        string
}

func newInstance() *validator.Validate {
	return validator.New()
}

/**
验证数据
*/
func ValidateField(validates []BlogValidate) error {
	if len(validates) > 0 {
		for _, value := range validates {
			e := Validate.Var(value.Key, value.Validation)
			if e != nil {
				return errors.New(value.Err)
			}
		}
	}
	return nil
}
