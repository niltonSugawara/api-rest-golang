package domain

import internal_errors "github.com/niltonSugawara/api-rest-golang/internal/internal-errors"

type Service interface {
	Create(novoProduto NovoProduto) (string, error)
}

type ServiceImp struct {
	Repository Repository
}

func (s *ServiceImp) Create(produto NovoProduto) (string, error) {
	produtoCadastrado, err := NovoProdutoCadastrado(produto.Nome, produto.Descricao, produto.Preco, produto.Categoria)
	if err != nil {
		return "ha algum campo nao preenchido", err
	}
	err = s.Repository.Create(produtoCadastrado)
	if err != nil {
		return "", internal_errors.ErroInterno
	}
	return produtoCadastrado.Id, nil
}

//func (s *ServiceImp) FindBy(id string) (*Produto, error) {
//	return Produto{}
//}
