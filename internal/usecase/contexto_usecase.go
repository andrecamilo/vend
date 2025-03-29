package usecase

import (
	"vend/internal/domain"
)

type ContextoUseCase struct {
	repo Repository
}

func NewContextoUseCase(repo Repository) *ContextoUseCase {
	return &ContextoUseCase{repo: repo}
}

func (u *ContextoUseCase) CreateContexto(contexto *domain.Contexto) error {
	return u.repo.CreateContexto(contexto)
}

func (u *ContextoUseCase) GetContexto(id string) (*domain.Contexto, error) {
	return u.repo.GetContexto(id)
}

func (u *ContextoUseCase) ListContextos() ([]domain.Contexto, error) {
	return u.repo.ListContextos()
}

func (u *ContextoUseCase) UpdateContexto(contexto *domain.Contexto) error {
	return u.repo.UpdateContexto(contexto)
}

func (u *ContextoUseCase) DeleteContexto(id string) error {
	return u.repo.DeleteContexto(id)
}
