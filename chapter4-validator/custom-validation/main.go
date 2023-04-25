package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// MyStruct ..
type MyStruct struct {
	String string `validate:"is-awesome"`
}

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func main() {

	validate = validator.New()
	validate.RegisterValidation("is-awesome", ValidateMyVal)

	s := MyStruct{String: "awesome"}

	err := validate.Struct(s)
	if err != nil {
		fmt.Printf("验证1，Err(s):\n%+v\n", err)
	}

	fmt.Printf("验证1结束\n")

	s.String = "not awesome"
	err = validate.Struct(s)
	if err != nil {
		fmt.Printf("验证2，Err(s):\n%+v\n", err)
	}
}

// ValidateMyVal implements validator.Func
func ValidateMyVal(fl validator.FieldLevel) bool {
	fmt.Printf("%+v\n", fl.Field().String())
	return fl.Field().String() == "awesome"
}
