package internal_errors

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"strings"
)

func ValidarStruct(objeto interface{}) error {
	validate := validator.New()
	err := validate.Struct(objeto)
	if err == nil {
		return nil
	}
	validarErros := err.(validator.ValidationErrors)
	validarErro := validarErros[0]

	campo := strings.ToLower(validarErro.StructField())
	switch validarErro.Tag() {
	case "required":
		return errors.New(campo + " is required")
	}
	return nil
}
