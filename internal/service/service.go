package service

import (
	"vend/internal/domain"
	"vend/internal/infrastructure/postgres"
)

type PessoaService struct {
	repo *postgres.Repository
}

func NewPessoaService(repo *postgres.Repository) *PessoaService {
	return &PessoaService{
		repo: repo,
	}
}

func (s *PessoaService) CreatePessoa(pessoa *domain.Pessoa) error {
	return s.repo.CreatePessoa(pessoa)
}

func (s *PessoaService) GetPessoa(id uint) (*domain.Pessoa, error) {
	return s.repo.GetPessoa(id)
}

func (s *PessoaService) ListPessoas() ([]domain.Pessoa, error) {
	return s.repo.ListPessoas()
}

func (s *PessoaService) UpdatePessoa(pessoa *domain.Pessoa) error {
	return s.repo.UpdatePessoa(pessoa)
}

func (s *PessoaService) DeletePessoa(id uint) error {
	return s.repo.DeletePessoa(id)
}
