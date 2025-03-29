package usecase

import (
	"vend/internal/domain"
)

type PromptUseCase struct {
	repo Repository
}

func NewPromptUseCase(repo Repository) *PromptUseCase {
	return &PromptUseCase{repo: repo}
}

func (u *PromptUseCase) CreatePrompt(prompt *domain.Prompt) error {
	return u.repo.CreatePrompt(prompt)
}

func (u *PromptUseCase) GetPrompt(id string) (*domain.Prompt, error) {
	return u.repo.GetPrompt(id)
}

func (u *PromptUseCase) ListPrompts() ([]domain.Prompt, error) {
	return u.repo.ListPrompts()
}

func (u *PromptUseCase) UpdatePrompt(prompt *domain.Prompt) error {
	return u.repo.UpdatePrompt(prompt)
}

func (u *PromptUseCase) DeletePrompt(id string) error {
	return u.repo.DeletePrompt(id)
}
