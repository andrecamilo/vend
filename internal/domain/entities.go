package domain

import (
	"time"
)

type Pessoa struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	Nome      string     `json:"nome" validate:"required"`
	Email     string     `json:"email" validate:"required,email"`
	Telefones []Telefone `json:"telefones" gorm:"foreignKey:PessoaID"`
	Contextos []Contexto `json:"contextos" gorm:"many2many:pessoa_contextos;"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type Telefone struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Numero   string `json:"numero" validate:"required"`
	Tipo     string `json:"tipo" validate:"required"`
	PessoaID uint   `json:"pessoa_id"`
	Pessoa   Pessoa `json:"pessoa" gorm:"foreignKey:PessoaID"`
}

type Contexto struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Nome       string    `json:"nome" validate:"required"`
	Descricao  string    `json:"descricao"`
	DataInicio time.Time `json:"data_inicio"`
	DataFim    time.Time `json:"data_fim"`
	Pessoas    []Pessoa  `json:"pessoas" gorm:"many2many:pessoa_contextos;"`
	Prompts    []Prompt  `json:"prompts" gorm:"foreignKey:ContextoID"`
}

type Prompt struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Conteudo   string    `json:"conteudo" validate:"required"`
	ContextoID uint      `json:"contexto_id"`
	Contexto   Contexto  `json:"contexto" gorm:"foreignKey:ContextoID"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Repository interface {
	CreatePessoa(pessoa *Pessoa) error
	GetPessoa(id uint) (*Pessoa, error)
	ListPessoas() ([]Pessoa, error)
	UpdatePessoa(pessoa *Pessoa) error
	DeletePessoa(id uint) error

	CreateTelefone(telefone *Telefone) error
	GetTelefone(id uint) (*Telefone, error)
	ListTelefones() ([]Telefone, error)
	UpdateTelefone(telefone *Telefone) error
	DeleteTelefone(id uint) error

	CreateContexto(contexto *Contexto) error
	GetContexto(id uint) (*Contexto, error)
	ListContextos() ([]Contexto, error)
	UpdateContexto(contexto *Contexto) error
	DeleteContexto(id uint) error

	CreatePrompt(prompt *Prompt) error
	GetPrompt(id uint) (*Prompt, error)
	ListPrompts() ([]Prompt, error)
	UpdatePrompt(prompt *Prompt) error
	DeletePrompt(id uint) error
}
