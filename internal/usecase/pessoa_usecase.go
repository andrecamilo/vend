package usecase

import (
	"vend/internal/domain"
)

type Repository interface {
	// Métodos de Pessoa
	CreatePessoa(pessoa *domain.Pessoa) error
	GetPessoa(id string) (*domain.Pessoa, error)
	ListPessoas() ([]domain.Pessoa, error)
	UpdatePessoa(pessoa *domain.Pessoa) error
	DeletePessoa(id string) error

	// Métodos de Telefone
	CreateTelefone(telefone *domain.Telefone) error
	GetTelefone(id string) (*domain.Telefone, error)
	ListTelefones() ([]domain.Telefone, error)
	UpdateTelefone(telefone *domain.Telefone) error
	DeleteTelefone(id string) error

	// Métodos de Contexto
	CreateContexto(contexto *domain.Contexto) error
	GetContexto(id string) (*domain.Contexto, error)
	ListContextos() ([]domain.Contexto, error)
	UpdateContexto(contexto *domain.Contexto) error
	DeleteContexto(id string) error

	// Métodos de Prompt
	CreatePrompt(prompt *domain.Prompt) error
	GetPrompt(id string) (*domain.Prompt, error)
	ListPrompts() ([]domain.Prompt, error)
	UpdatePrompt(prompt *domain.Prompt) error
	DeletePrompt(id string) error
}

type PessoaUseCase struct {
	repo Repository
}

func NewPessoaUseCase(repo Repository) *PessoaUseCase {
	return &PessoaUseCase{repo: repo}
}

func (u *PessoaUseCase) CreatePessoa(pessoa *domain.Pessoa) error {
	return u.repo.CreatePessoa(pessoa)
}

func (u *PessoaUseCase) GetPessoa(id string) (*domain.Pessoa, error) {
	return u.repo.GetPessoa(id)
}

func (u *PessoaUseCase) ListPessoas() ([]domain.Pessoa, error) {
	return u.repo.ListPessoas()
}

func (u *PessoaUseCase) UpdatePessoa(pessoa *domain.Pessoa) error {
	return u.repo.UpdatePessoa(pessoa)
}

func (u *PessoaUseCase) DeletePessoa(id string) error {
	return u.repo.DeletePessoa(id)
}
