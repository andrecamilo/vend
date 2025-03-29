package usecase

import (
	"vend/internal/domain"
)

type TelefoneUseCase struct {
	repo Repository
}

func NewTelefoneUseCase(repo Repository) *TelefoneUseCase {
	return &TelefoneUseCase{repo: repo}
}

func (u *TelefoneUseCase) CreateTelefone(telefone *domain.Telefone) error {
	return u.repo.CreateTelefone(telefone)
}

func (u *TelefoneUseCase) GetTelefone(id string) (*domain.Telefone, error) {
	return u.repo.GetTelefone(id)
}

func (u *TelefoneUseCase) ListTelefones() ([]domain.Telefone, error) {
	return u.repo.ListTelefones()
}

func (u *TelefoneUseCase) UpdateTelefone(telefone *domain.Telefone) error {
	return u.repo.UpdateTelefone(telefone)
}

func (u *TelefoneUseCase) DeleteTelefone(id string) error {
	return u.repo.DeleteTelefone(id)
}
