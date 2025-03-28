package postgres

import (
	"vend/internal/domain"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreatePessoa(pessoa *domain.Pessoa) error {
	return r.db.Create(pessoa).Error
}

func (r *Repository) GetPessoa(id uint) (*domain.Pessoa, error) {
	var pessoa domain.Pessoa
	err := r.db.Preload("Telefones").Preload("Contextos").First(&pessoa, id).Error
	return &pessoa, err
}

func (r *Repository) ListPessoas() ([]domain.Pessoa, error) {
	var pessoas []domain.Pessoa
	err := r.db.Preload("Telefones").Preload("Contextos").Find(&pessoas).Error
	return pessoas, err
}

func (r *Repository) UpdatePessoa(pessoa *domain.Pessoa) error {
	return r.db.Save(pessoa).Error
}

func (r *Repository) DeletePessoa(id uint) error {
	return r.db.Delete(&domain.Pessoa{}, id).Error
}

func (r *Repository) CreateTelefone(telefone *domain.Telefone) error {
	return r.db.Create(telefone).Error
}

func (r *Repository) GetTelefone(id uint) (*domain.Telefone, error) {
	var telefone domain.Telefone
	err := r.db.First(&telefone, id).Error
	return &telefone, err
}

func (r *Repository) ListTelefones() ([]domain.Telefone, error) {
	var telefones []domain.Telefone
	err := r.db.Find(&telefones).Error
	return telefones, err
}

func (r *Repository) UpdateTelefone(telefone *domain.Telefone) error {
	return r.db.Save(telefone).Error
}

func (r *Repository) DeleteTelefone(id uint) error {
	return r.db.Delete(&domain.Telefone{}, id).Error
}

func (r *Repository) CreateContexto(contexto *domain.Contexto) error {
	return r.db.Create(contexto).Error
}

func (r *Repository) GetContexto(id uint) (*domain.Contexto, error) {
	var contexto domain.Contexto
	err := r.db.Preload("Pessoas").Preload("Prompts").First(&contexto, id).Error
	return &contexto, err
}

func (r *Repository) ListContextos() ([]domain.Contexto, error) {
	var contextos []domain.Contexto
	err := r.db.Preload("Pessoas").Preload("Prompts").Find(&contextos).Error
	return contextos, err
}

func (r *Repository) UpdateContexto(contexto *domain.Contexto) error {
	return r.db.Save(contexto).Error
}

func (r *Repository) DeleteContexto(id uint) error {
	return r.db.Delete(&domain.Contexto{}, id).Error
}

func (r *Repository) CreatePrompt(prompt *domain.Prompt) error {
	return r.db.Create(prompt).Error
}

func (r *Repository) GetPrompt(id uint) (*domain.Prompt, error) {
	var prompt domain.Prompt
	err := r.db.First(&prompt, id).Error
	return &prompt, err
}

func (r *Repository) ListPrompts() ([]domain.Prompt, error) {
	var prompts []domain.Prompt
	err := r.db.Find(&prompts).Error
	return prompts, err
}

func (r *Repository) UpdatePrompt(prompt *domain.Prompt) error {
	return r.db.Save(prompt).Error
}

func (r *Repository) DeletePrompt(id uint) error {
	return r.db.Delete(&domain.Prompt{}, id).Error
}
