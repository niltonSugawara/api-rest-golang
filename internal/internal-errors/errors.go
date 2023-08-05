package internal_errors

import (
	"errors"
	"net/http"
)

var ErroInterno error = errors.New("Internal Server Error")

func ProcessarErroParaRetornar(erro error) error {
	if !errors.Is(erro, http.ErrBodyNotAllowed) {
		return ErroInterno
	}
	return erro
}
