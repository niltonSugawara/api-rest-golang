package database

import "github.com/niltonSugawara/api-rest-golang/internal/domain"

type ProdutoRepository struct {
	produtos []domain.Produto
}

func (p *ProdutoRepository) Create(produto *domain.Produto) error {
	p.produtos = append(p.produtos, *produto)
	return nil
}
