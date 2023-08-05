package endpoints

import (
	"errors"
	"github.com/go-chi/render"
	internal_errors "github.com/niltonSugawara/api-rest-golang/internal/internal-errors"
	"net/http"
)

type EndpointFunc func(w http.ResponseWriter, r *http.Request) (interface{}, int, error)

func HandlerError(endpointFunc EndpointFunc) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		objeto, status, err := endpointFunc(writer, request)
		if err != nil {

			if errors.Is(err, internal_errors.ErroInterno) {
				render.Status(request, http.StatusInternalServerError)
			} else {
				render.Status(request, http.StatusBadRequest)
			}
			render.JSON(writer, request, map[string]string{"erro": err.Error()})
			return
		}
		render.Status(request, status)
		if objeto != nil {
			render.JSON(writer, request, objeto)
		}
	})
}
