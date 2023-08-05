package domain

type Repository interface {
	Create(produto *Produto) error
}
