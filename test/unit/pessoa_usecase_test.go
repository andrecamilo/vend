package unit

import (
	"testing"
	"vend/internal/domain"
	"vend/internal/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreatePessoa(pessoa *domain.Pessoa) error {
	args := m.Called(pessoa)
	return args.Error(0)
}

func (m *MockRepository) GetPessoa(id uint) (*domain.Pessoa, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Pessoa), args.Error(1)
}

func (m *MockRepository) ListPessoas() ([]domain.Pessoa, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.Pessoa), args.Error(1)
}

func (m *MockRepository) UpdatePessoa(pessoa *domain.Pessoa) error {
	args := m.Called(pessoa)
	return args.Error(0)
}

func (m *MockRepository) DeletePessoa(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockRepository) CreateTelefone(telefone *domain.Telefone) error {
	args := m.Called(telefone)
	return args.Error(0)
}

func (m *MockRepository) GetTelefone(id uint) (*domain.Telefone, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Telefone), args.Error(1)
}

func (m *MockRepository) ListTelefones() ([]domain.Telefone, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.Telefone), args.Error(1)
}

func (m *MockRepository) UpdateTelefone(telefone *domain.Telefone) error {
	args := m.Called(telefone)
	return args.Error(0)
}

func (m *MockRepository) DeleteTelefone(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockRepository) CreateContexto(contexto *domain.Contexto) error {
	args := m.Called(contexto)
	return args.Error(0)
}

func (m *MockRepository) GetContexto(id uint) (*domain.Contexto, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Contexto), args.Error(1)
}

func (m *MockRepository) ListContextos() ([]domain.Contexto, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.Contexto), args.Error(1)
}

func (m *MockRepository) UpdateContexto(contexto *domain.Contexto) error {
	args := m.Called(contexto)
	return args.Error(0)
}

func (m *MockRepository) DeleteContexto(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockRepository) CreatePrompt(prompt *domain.Prompt) error {
	args := m.Called(prompt)
	return args.Error(0)
}

func (m *MockRepository) GetPrompt(id uint) (*domain.Prompt, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Prompt), args.Error(1)
}

func (m *MockRepository) ListPrompts() ([]domain.Prompt, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.Prompt), args.Error(1)
}

func (m *MockRepository) UpdatePrompt(prompt *domain.Prompt) error {
	args := m.Called(prompt)
	return args.Error(0)
}

func (m *MockRepository) DeletePrompt(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreatePessoa(t *testing.T) {
	mockRepo := new(MockRepository)
	useCase := usecase.NewPessoaUseCase(mockRepo)

	pessoa := &domain.Pessoa{
		Nome:  "Teste",
		Email: "teste@teste.com",
	}

	mockRepo.On("CreatePessoa", pessoa).Return(nil)

	err := useCase.CreatePessoa(pessoa)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetPessoa(t *testing.T) {
	mockRepo := new(MockRepository)
	useCase := usecase.NewPessoaUseCase(mockRepo)

	expectedPessoa := &domain.Pessoa{
		ID:    1,
		Nome:  "Teste",
		Email: "teste@teste.com",
	}

	mockRepo.On("GetPessoa", uint(1)).Return(expectedPessoa, nil)

	pessoa, err := useCase.GetPessoa(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedPessoa, pessoa)
	mockRepo.AssertExpectations(t)
}

func TestListPessoas(t *testing.T) {
	mockRepo := new(MockRepository)
	useCase := usecase.NewPessoaUseCase(mockRepo)

	expectedPessoas := []domain.Pessoa{
		{
			ID:    1,
			Nome:  "Teste 1",
			Email: "teste1@teste.com",
		},
		{
			ID:    2,
			Nome:  "Teste 2",
			Email: "teste2@teste.com",
		},
	}

	mockRepo.On("ListPessoas").Return(expectedPessoas, nil)

	pessoas, err := useCase.ListPessoas()

	assert.NoError(t, err)
	assert.Equal(t, expectedPessoas, pessoas)
	mockRepo.AssertExpectations(t)
}

func TestUpdatePessoa(t *testing.T) {
	mockRepo := new(MockRepository)
	useCase := usecase.NewPessoaUseCase(mockRepo)

	pessoa := &domain.Pessoa{
		ID:    1,
		Nome:  "Teste Atualizado",
		Email: "teste.atualizado@teste.com",
	}

	mockRepo.On("UpdatePessoa", pessoa).Return(nil)

	err := useCase.UpdatePessoa(pessoa)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeletePessoa(t *testing.T) {
	mockRepo := new(MockRepository)
	useCase := usecase.NewPessoaUseCase(mockRepo)

	mockRepo.On("DeletePessoa", uint(1)).Return(nil)

	err := useCase.DeletePessoa(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
