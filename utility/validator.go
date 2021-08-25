package utility

import "github.com/go-playground/validator"

type IValidator interface {
	Validate(interface{}) error
}

type vengine struct {
	engine *validator.Validate
}

func (v *vengine) Validate(val interface{}) error {
	return v.engine.Struct(val)
}

var Validator IValidator = &vengine{engine: validator.New()}
