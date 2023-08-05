package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/niltonSugawara/api-rest-golang/internal/domain"
	"github.com/niltonSugawara/api-rest-golang/internal/endpoints"
	"github.com/niltonSugawara/api-rest-golang/internal/infra/database"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	produtoService := domain.ServiceImp{
		Repository: &database.ProdutoRepository{},
	}
	handler := endpoints.Handler{
		ProdutoService: &produtoService,
	}
	r.Post("/produtos", endpoints.HandlerError(handler.ProdutoPost))
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		return
	}
}
