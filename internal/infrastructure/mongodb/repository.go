package mongodb

import (
	"context"
	"time"
	"vend/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{db: db}
}

// Pessoas
func (r *Repository) CreatePessoa(pessoa *domain.Pessoa) error {
	pessoa.CreatedAt = time.Now()
	pessoa.UpdatedAt = time.Now()

	result, err := r.db.Collection("pessoas").InsertOne(context.Background(), pessoa)
	if err != nil {
		return err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		pessoa.ID = oid
	}
	return nil
}

func (r *Repository) GetPessoa(id string) (*domain.Pessoa, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var pessoa domain.Pessoa
	err = r.db.Collection("pessoas").FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&pessoa)
	if err != nil {
		return nil, err
	}
	return &pessoa, nil
}

func (r *Repository) ListPessoas() ([]domain.Pessoa, error) {
	var pessoas []domain.Pessoa
	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := r.db.Collection("pessoas").Find(context.Background(), bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err = cursor.All(context.Background(), &pessoas); err != nil {
		return nil, err
	}
	return pessoas, nil
}

func (r *Repository) UpdatePessoa(pessoa *domain.Pessoa) error {
	pessoa.UpdatedAt = time.Now()

	_, err := r.db.Collection("pessoas").UpdateOne(
		context.Background(),
		bson.M{"_id": pessoa.ID},
		bson.M{"$set": bson.M{
			"nome":       pessoa.Nome,
			"email":      pessoa.Email,
			"telefones":  pessoa.Telefones,
			"contextos":  pessoa.Contextos,
			"updated_at": pessoa.UpdatedAt,
		}},
	)
	return err
}

func (r *Repository) DeletePessoa(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.db.Collection("pessoas").DeleteOne(context.Background(), bson.M{"_id": objectID})
	return err
}

// Implementar métodos similares para Telefone, Contexto e Prompt...
