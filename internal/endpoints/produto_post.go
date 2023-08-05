package endpoints

import (
	"github.com/go-chi/render"
	"github.com/niltonSugawara/api-rest-golang/internal/domain"
	"net/http"
)

func (h *Handler) ProdutoPost(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var requisicao domain.NovoProduto
	render.DecodeJSON(r.Body, &requisicao)
	id, err := h.ProdutoService.Create(requisicao)
	return map[string]string{"id": id}, 201, err
}
