package repository

import (
	"context"
	"os"
	"time"
	"vend/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PessoaRepository struct {
	db *mongo.Database
}

func NewPessoaRepository(client *mongo.Client) *PessoaRepository {
	dbName := "vend"
	if dbNameEnv := os.Getenv("MONGODB_DATABASE"); dbNameEnv != "" {
		dbName = dbNameEnv
	}
	return &PessoaRepository{db: client.Database(dbName)}
}

// Métodos de Pessoa
func (r *PessoaRepository) CreatePessoa(pessoa *domain.Pessoa) error {
	collection := r.db.Collection("pessoas")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pessoa.CreatedAt = time.Now()
	pessoa.UpdatedAt = time.Now()

	result, err := collection.InsertOne(ctx, pessoa)
	if err != nil {
		return err
	}

	pessoa.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *PessoaRepository) GetPessoa(id string) (*domain.Pessoa, error) {
	collection := r.db.Collection("pessoas")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var pessoa domain.Pessoa
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&pessoa)
	if err != nil {
		return nil, err
	}

	return &pessoa, nil
}

func (r *PessoaRepository) ListPessoas() ([]domain.Pessoa, error) {
	collection := r.db.Collection("pessoas")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var pessoas []domain.Pessoa
	if err = cursor.All(ctx, &pessoas); err != nil {
		return nil, err
	}

	return pessoas, nil
}

func (r *PessoaRepository) UpdatePessoa(pessoa *domain.Pessoa) error {
	collection := r.db.Collection("pessoas")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pessoa.UpdatedAt = time.Now()

	_, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": pessoa.ID},
		bson.M{"$set": pessoa},
	)
	return err
}

func (r *PessoaRepository) DeletePessoa(id string) error {
	collection := r.db.Collection("pessoas")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}

// Métodos de Telefone
func (r *PessoaRepository) CreateTelefone(telefone *domain.Telefone) error {
	collection := r.db.Collection("telefones")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, telefone)
	if err != nil {
		return err
	}

	telefone.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *PessoaRepository) GetTelefone(id string) (*domain.Telefone, error) {
	collection := r.db.Collection("telefones")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var telefone domain.Telefone
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&telefone)
	if err != nil {
		return nil, err
	}

	return &telefone, nil
}

func (r *PessoaRepository) ListTelefones() ([]domain.Telefone, error) {
	collection := r.db.Collection("telefones")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var telefones []domain.Telefone
	if err = cursor.All(ctx, &telefones); err != nil {
		return nil, err
	}

	return telefones, nil
}

func (r *PessoaRepository) UpdateTelefone(telefone *domain.Telefone) error {
	collection := r.db.Collection("telefones")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": telefone.ID},
		bson.M{"$set": telefone},
	)
	return err
}

func (r *PessoaRepository) DeleteTelefone(id string) error {
	collection := r.db.Collection("telefones")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}

// Métodos de Contexto
func (r *PessoaRepository) CreateContexto(contexto *domain.Contexto) error {
	collection := r.db.Collection("contextos")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, contexto)
	if err != nil {
		return err
	}

	contexto.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *PessoaRepository) GetContexto(id string) (*domain.Contexto, error) {
	collection := r.db.Collection("contextos")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var contexto domain.Contexto
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&contexto)
	if err != nil {
		return nil, err
	}

	return &contexto, nil
}

func (r *PessoaRepository) ListContextos() ([]domain.Contexto, error) {
	collection := r.db.Collection("contextos")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var contextos []domain.Contexto
	if err = cursor.All(ctx, &contextos); err != nil {
		return nil, err
	}

	return contextos, nil
}

func (r *PessoaRepository) UpdateContexto(contexto *domain.Contexto) error {
	collection := r.db.Collection("contextos")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": contexto.ID},
		bson.M{"$set": contexto},
	)
	return err
}

func (r *PessoaRepository) DeleteContexto(id string) error {
	collection := r.db.Collection("contextos")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}

// Métodos de Prompt
func (r *PessoaRepository) CreatePrompt(prompt *domain.Prompt) error {
	collection := r.db.Collection("prompts")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	prompt.CreatedAt = time.Now()
	prompt.UpdatedAt = time.Now()

	result, err := collection.InsertOne(ctx, prompt)
	if err != nil {
		return err
	}

	prompt.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *PessoaRepository) GetPrompt(id string) (*domain.Prompt, error) {
	collection := r.db.Collection("prompts")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var prompt domain.Prompt
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&prompt)
	if err != nil {
		return nil, err
	}

	return &prompt, nil
}

func (r *PessoaRepository) ListPrompts() ([]domain.Prompt, error) {
	collection := r.db.Collection("prompts")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var prompts []domain.Prompt
	if err = cursor.All(ctx, &prompts); err != nil {
		return nil, err
	}

	return prompts, nil
}

func (r *PessoaRepository) UpdatePrompt(prompt *domain.Prompt) error {
	collection := r.db.Collection("prompts")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	prompt.UpdatedAt = time.Now()

	_, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": prompt.ID},
		bson.M{"$set": prompt},
	)
	return err
}

func (r *PessoaRepository) DeletePrompt(id string) error {
	collection := r.db.Collection("prompts")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}
