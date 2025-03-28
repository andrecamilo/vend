package integration

import (
	"testing"
	"vend/internal/domain"
	postgresRepo "vend/internal/infrastructure/postgres"
	"vend/internal/usecase"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=vend_test port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	// Limpar e migrar tabelas
	db.Migrator().DropTable(&domain.Pessoa{}, &domain.Telefone{}, &domain.Contexto{}, &domain.Prompt{})
	db.AutoMigrate(&domain.Pessoa{}, &domain.Telefone{}, &domain.Contexto{}, &domain.Prompt{})

	return db
}

func TestPessoaIntegration(t *testing.T) {
	db := setupTestDB(t)
	repo := postgresRepo.NewRepository(db)
	useCase := usecase.NewPessoaUseCase(repo)

	t.Run("Criar e recuperar pessoa", func(t *testing.T) {
		pessoa := &domain.Pessoa{
			Nome:  "Teste Integração",
			Email: "teste.integracao@teste.com",
		}

		err := useCase.CreatePessoa(pessoa)
		assert.NoError(t, err)
		assert.NotZero(t, pessoa.ID)

		recuperada, err := useCase.GetPessoa(pessoa.ID)
		assert.NoError(t, err)
		assert.Equal(t, pessoa.Nome, recuperada.Nome)
		assert.Equal(t, pessoa.Email, recuperada.Email)
	})

	t.Run("Listar pessoas", func(t *testing.T) {
		pessoas := []domain.Pessoa{
			{
				Nome:  "Teste 1",
				Email: "teste1@teste.com",
			},
			{
				Nome:  "Teste 2",
				Email: "teste2@teste.com",
			},
		}

		for _, p := range pessoas {
			err := useCase.CreatePessoa(&p)
			assert.NoError(t, err)
		}

		lista, err := useCase.ListPessoas()
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(lista), 2)
	})

	t.Run("Atualizar pessoa", func(t *testing.T) {
		pessoa := &domain.Pessoa{
			Nome:  "Teste Atualização",
			Email: "teste.atualizacao@teste.com",
		}

		err := useCase.CreatePessoa(pessoa)
		assert.NoError(t, err)

		pessoa.Nome = "Teste Atualizado"
		err = useCase.UpdatePessoa(pessoa)
		assert.NoError(t, err)

		atualizada, err := useCase.GetPessoa(pessoa.ID)
		assert.NoError(t, err)
		assert.Equal(t, "Teste Atualizado", atualizada.Nome)
	})

	t.Run("Deletar pessoa", func(t *testing.T) {
		pessoa := &domain.Pessoa{
			Nome:  "Teste Deleção",
			Email: "teste.delecao@teste.com",
		}

		err := useCase.CreatePessoa(pessoa)
		assert.NoError(t, err)

		err = useCase.DeletePessoa(pessoa.ID)
		assert.NoError(t, err)

		_, err = useCase.GetPessoa(pessoa.ID)
		assert.Error(t, err)
	})
}
