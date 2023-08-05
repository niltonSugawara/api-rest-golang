package domain

import (
	"github.com/google/uuid"
	internal_errors "github.com/niltonSugawara/api-rest-golang/internal/internal-errors"
	"time"
)

type Produto struct {
	Id          string    `json:"id" validate:"required"`
	Nome        string    `json:"nome" validate:"required"`
	Descricao   string    `json:"descricao"`
	Preco       float64   `json:"preco" validate:"required"`
	Categoria   string    `json:"categoria" validate:"required"`
	DataCriacao time.Time `json:"data-criacao" validate:"required"`
}
type ProdutoResponse struct {
	Nome      string `json:"nome"`
	Descricao string
	Preco     float64
	Categoria string
}

type Contato struct {
	Telefone string
	Email    string
	Endereco string
}

type NovoProduto struct {
	Nome      string  `json:"nome"`
	Descricao string  `json:"descricao"`
	Preco     float64 `json:"preco"`
	Categoria string  `json:"categoria"`
}

func NovoProdutoCadastrado(nome string, descricao string, preco float64, categoria string) (*Produto, error) {
	produtoCadastrado := &Produto{
		Id:          uuid.New().String(),
		Nome:        nome,
		Descricao:   descricao,
		Preco:       preco,
		Categoria:   categoria,
		DataCriacao: time.Now(),
	}
	err := internal_errors.ValidarStruct(produtoCadastrado)
	if err == nil {
		return produtoCadastrado, nil
	}
	return nil, err
}
