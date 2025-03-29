package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pessoa struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nome      string             `bson:"nome" json:"nome" binding:"required"`
	Email     string             `bson:"email" json:"email" binding:"required"`
	Telefones []Telefone         `bson:"telefones,omitempty" json:"telefones,omitempty"`
	Contextos []Contexto         `bson:"contextos,omitempty" json:"contextos,omitempty"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

type Telefone struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Numero   string             `bson:"numero" json:"numero" binding:"required"`
	Tipo     string             `bson:"tipo" json:"tipo" binding:"required"`
	PessoaID primitive.ObjectID `bson:"pessoa_id" json:"pessoa_id"`
}

type Contexto struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nome       string             `bson:"nome" json:"nome" binding:"required"`
	Descricao  string             `bson:"descricao" json:"descricao"`
	DataInicio time.Time          `bson:"data_inicio" json:"data_inicio"`
	DataFim    time.Time          `bson:"data_fim" json:"data_fim"`
	Pessoas    []Pessoa           `bson:"pessoas,omitempty" json:"pessoas,omitempty"`
	Prompts    []Prompt           `bson:"prompts,omitempty" json:"prompts,omitempty"`
}

type Prompt struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Conteudo   string             `bson:"conteudo" json:"conteudo" binding:"required"`
	ContextoID primitive.ObjectID `bson:"contexto_id" json:"contexto_id"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
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
