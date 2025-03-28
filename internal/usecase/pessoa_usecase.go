package usecase

import (
	"vend/internal/domain"
)

type PessoaUseCase struct {
	repo domain.Repository
}

func NewPessoaUseCase(repo domain.Repository) *PessoaUseCase {
	return &PessoaUseCase{repo: repo}
}

func (u *PessoaUseCase) CreatePessoa(pessoa *domain.Pessoa) error {
	return u.repo.CreatePessoa(pessoa)
}

func (u *PessoaUseCase) GetPessoa(id uint) (*domain.Pessoa, error) {
	return u.repo.GetPessoa(id)
}

func (u *PessoaUseCase) ListPessoas() ([]domain.Pessoa, error) {
	return u.repo.ListPessoas()
}

func (u *PessoaUseCase) UpdatePessoa(pessoa *domain.Pessoa) error {
	return u.repo.UpdatePessoa(pessoa)
}

func (u *PessoaUseCase) DeletePessoa(id uint) error {
	return u.repo.DeletePessoa(id)
}
