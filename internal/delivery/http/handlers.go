package http

import (
	"net/http"
	"vend/internal/domain"
	"vend/internal/usecase"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handler struct {
	pessoaUseCase   *usecase.PessoaUseCase
	telefoneUseCase *usecase.TelefoneUseCase
	contextoUseCase *usecase.ContextoUseCase
	promptUseCase   *usecase.PromptUseCase
}

func NewHandler(
	pessoaUseCase *usecase.PessoaUseCase,
	telefoneUseCase *usecase.TelefoneUseCase,
	contextoUseCase *usecase.ContextoUseCase,
	promptUseCase *usecase.PromptUseCase,
) *Handler {
	return &Handler{
		pessoaUseCase:   pessoaUseCase,
		telefoneUseCase: telefoneUseCase,
		contextoUseCase: contextoUseCase,
		promptUseCase:   promptUseCase,
	}
}

// @Summary     Criar pessoa
// @Description Cria uma nova pessoa no sistema
// @Tags        pessoas
// @Accept      json
// @Produce     json
// @Param       pessoa body domain.Pessoa true "Dados da pessoa"
// @Success     201 {object} domain.Pessoa
// @Failure     400 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /pessoas [post]
func (h *Handler) CreatePessoa(c *gin.Context) {
	var pessoa domain.Pessoa
	if err := c.ShouldBindJSON(&pessoa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := h.pessoaUseCase.CreatePessoa(&pessoa); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, pessoa)
}

// @Summary     Buscar pessoa
// @Description Retorna os dados de uma pessoa específica
// @Tags        pessoas
// @Accept      json
// @Produce     json
// @Param       id path string true "ID da pessoa"
// @Success     200 {object} domain.Pessoa
// @Failure     400 {object} map[string]string
// @Failure     404 {object} map[string]string
// @Router      /pessoas/{id} [get]
func (h *Handler) GetPessoa(c *gin.Context) {
	id := c.Param("id")
	if !primitive.IsValidObjectID(id) {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	pessoa, err := h.pessoaUseCase.GetPessoa(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Pessoa não encontrada"})
		return
	}

	c.JSON(http.StatusOK, pessoa)
}

// @Summary     Listar pessoas
// @Description Retorna a lista de todas as pessoas cadastradas
// @Tags        pessoas
// @Accept      json
// @Produce     json
// @Success     200 {array} domain.Pessoa
// @Failure     500 {object} map[string]string
// @Router      /pessoas [get]
func (h *Handler) ListPessoas(c *gin.Context) {
	pessoas, err := h.pessoaUseCase.ListPessoas()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, pessoas)
}

// @Summary     Atualizar pessoa
// @Description Atualiza os dados de uma pessoa específica
// @Tags        pessoas
// @Accept      json
// @Produce     json
// @Param       id path string true "ID da pessoa"
// @Param       pessoa body domain.Pessoa true "Dados da pessoa"
// @Success     200 {object} domain.Pessoa
// @Failure     400 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /pessoas/{id} [put]
func (h *Handler) UpdatePessoa(c *gin.Context) {
	id := c.Param("id")
	if !primitive.IsValidObjectID(id) {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	var pessoa domain.Pessoa
	if err := c.ShouldBindJSON(&pessoa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	pessoa.ID = objectID
	if err := h.pessoaUseCase.UpdatePessoa(&pessoa); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, pessoa)
}

// @Summary     Deletar pessoa
// @Description Remove uma pessoa do sistema
// @Tags        pessoas
// @Accept      json
// @Produce     json
// @Param       id path string true "ID da pessoa"
// @Success     200 {object} map[string]string
// @Failure     400 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /pessoas/{id} [delete]
func (h *Handler) DeletePessoa(c *gin.Context) {
	id := c.Param("id")
	if !primitive.IsValidObjectID(id) {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	if err := h.pessoaUseCase.DeletePessoa(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensagem": "Pessoa deletada com sucesso"})
}

// @Summary     Listar telefones
// @Description Retorna a lista de todos os telefones cadastrados
// @Tags        telefones
// @Accept      json
// @Produce     json
// @Success     200 {array} domain.Telefone
// @Failure     500 {object} map[string]string
// @Router      /telefones [get]
func (h *Handler) ListTelefones(c *gin.Context) {
	telefones, err := h.telefoneUseCase.ListTelefones()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, telefones)
}

// @Summary     Criar telefone
// @Description Cria um novo telefone no sistema
// @Tags        telefones
// @Accept      json
// @Produce     json
// @Param       telefone body domain.Telefone true "Dados do telefone"
// @Success     201 {object} domain.Telefone
// @Failure     400 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /telefones [post]
func (h *Handler) CreateTelefone(c *gin.Context) {
	var telefone domain.Telefone
	if err := c.ShouldBindJSON(&telefone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := h.telefoneUseCase.CreateTelefone(&telefone); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, telefone)
}

// @Summary     Buscar telefone
// @Description Retorna os dados de um telefone específico
// @Tags        telefones
// @Accept      json
// @Produce     json
// @Param       id path string true "ID do telefone"
// @Success     200 {object} domain.Telefone
// @Failure     400 {object} map[string]string
// @Failure     404 {object} map[string]string
// @Router      /telefones/{id} [get]
func (h *Handler) GetTelefone(c *gin.Context) {
	id := c.Param("id")
	if !primitive.IsValidObjectID(id) {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	telefone, err := h.telefoneUseCase.GetTelefone(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Telefone não encontrado"})
		return
	}

	c.JSON(http.StatusOK, telefone)
}

// @Summary     Atualizar telefone
// @Description Atualiza os dados de um telefone específico
// @Tags        telefones
// @Accept      json
// @Produce     json
// @Param       id path string true "ID do telefone"
// @Param       telefone body domain.Telefone true "Dados do telefone"
// @Success     200 {object} domain.Telefone
// @Failure     400 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /telefones/{id} [put]
func (h *Handler) UpdateTelefone(c *gin.Context) {
	id := c.Param("id")
	if !primitive.IsValidObjectID(id) {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	var telefone domain.Telefone
	if err := c.ShouldBindJSON(&telefone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	telefone.ID = objectID
	if err := h.telefoneUseCase.UpdateTelefone(&telefone); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, telefone)
}

// @Summary     Deletar telefone
// @Description Remove um telefone do sistema
// @Tags        telefones
// @Accept      json
// @Produce     json
// @Param       id path string true "ID do telefone"
// @Success     200 {object} map[string]string
// @Failure     400 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /telefones/{id} [delete]
func (h *Handler) DeleteTelefone(c *gin.Context) {
	id := c.Param("id")
	if !primitive.IsValidObjectID(id) {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	if err := h.telefoneUseCase.DeleteTelefone(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensagem": "Telefone deletado com sucesso"})
}

// @Summary     Listar contextos
// @Description Retorna a lista de todos os contextos cadastrados
// @Tags        contextos
// @Accept      json
// @Produce     json
// @Success     200 {array} domain.Contexto
// @Failure     500 {object} map[string]string
// @Router      /contextos [get]
func (h *Handler) ListContextos(c *gin.Context) {
	contextos, err := h.contextoUseCase.ListContextos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, contextos)
}

// @Summary     Criar contexto
// @Description Cria um novo contexto no sistema
// @Tags        contextos
// @Accept      json
// @Produce     json
// @Param       contexto body domain.Contexto true "Dados do contexto"
// @Success     201 {object} domain.Contexto
// @Failure     400 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /contextos [post]
func (h *Handler) CreateContexto(c *gin.Context) {
	var contexto domain.Contexto
	if err := c.ShouldBindJSON(&contexto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := h.contextoUseCase.CreateContexto(&contexto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, contexto)
}

// @Summary     Buscar contexto
// @Description Retorna os dados de um contexto específico
// @Tags        contextos
// @Accept      json
// @Produce     json
// @Param       id path string true "ID do contexto"
// @Success     200 {object} domain.Contexto
// @Failure     400 {object} map[string]string
// @Failure     404 {object} map[string]string
// @Router      /contextos/{id} [get]
func (h *Handler) GetContexto(c *gin.Context) {
	id := c.Param("id")
	if !primitive.IsValidObjectID(id) {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	contexto, err := h.contextoUseCase.GetContexto(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Contexto não encontrado"})
		return
	}

	c.JSON(http.StatusOK, contexto)
}

// @Summary     Atualizar contexto
// @Description Atualiza os dados de um contexto específico
// @Tags        contextos
// @Accept      json
// @Produce     json
// @Param       id path string true "ID do contexto"
// @Param       contexto body domain.Contexto true "Dados do contexto"
// @Success     200 {object} domain.Contexto
// @Failure     400 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /contextos/{id} [put]
func (h *Handler) UpdateContexto(c *gin.Context) {
	id := c.Param("id")
	if !primitive.IsValidObjectID(id) {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	var contexto domain.Contexto
	if err := c.ShouldBindJSON(&contexto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	contexto.ID = objectID
	if err := h.contextoUseCase.UpdateContexto(&contexto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, contexto)
}

// @Summary     Deletar contexto
// @Description Remove um contexto do sistema
// @Tags        contextos
// @Accept      json
// @Produce     json
// @Param       id path string true "ID do contexto"
// @Success     200 {object} map[string]string
// @Failure     400 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /contextos/{id} [delete]
func (h *Handler) DeleteContexto(c *gin.Context) {
	id := c.Param("id")
	if !primitive.IsValidObjectID(id) {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	if err := h.contextoUseCase.DeleteContexto(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensagem": "Contexto deletado com sucesso"})
}

// @Summary     Listar prompts
// @Description Retorna a lista de todos os prompts cadastrados
// @Tags        prompts
// @Accept      json
// @Produce     json
// @Success     200 {array} domain.Prompt
// @Failure     500 {object} map[string]string
// @Router      /prompts [get]
func (h *Handler) ListPrompts(c *gin.Context) {
	prompts, err := h.promptUseCase.ListPrompts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, prompts)
}

// @Summary     Criar prompt
// @Description Cria um novo prompt no sistema
// @Tags        prompts
// @Accept      json
// @Produce     json
// @Param       prompt body domain.Prompt true "Dados do prompt"
// @Success     201 {object} domain.Prompt
// @Failure     400 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /prompts [post]
func (h *Handler) CreatePrompt(c *gin.Context) {
	var prompt domain.Prompt
	if err := c.ShouldBindJSON(&prompt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := h.promptUseCase.CreatePrompt(&prompt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, prompt)
}

// @Summary     Buscar prompt
// @Description Retorna os dados de um prompt específico
// @Tags        prompts
// @Accept      json
// @Produce     json
// @Param       id path string true "ID do prompt"
// @Success     200 {object} domain.Prompt
// @Failure     400 {object} map[string]string
// @Failure     404 {object} map[string]string
// @Router      /prompts/{id} [get]
func (h *Handler) GetPrompt(c *gin.Context) {
	id := c.Param("id")
	if !primitive.IsValidObjectID(id) {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	prompt, err := h.promptUseCase.GetPrompt(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Prompt não encontrado"})
		return
	}

	c.JSON(http.StatusOK, prompt)
}

// @Summary     Atualizar prompt
// @Description Atualiza os dados de um prompt específico
// @Tags        prompts
// @Accept      json
// @Produce     json
// @Param       id path string true "ID do prompt"
// @Param       prompt body domain.Prompt true "Dados do prompt"
// @Success     200 {object} domain.Prompt
// @Failure     400 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /prompts/{id} [put]
func (h *Handler) UpdatePrompt(c *gin.Context) {
	id := c.Param("id")
	if !primitive.IsValidObjectID(id) {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	var prompt domain.Prompt
	if err := c.ShouldBindJSON(&prompt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	prompt.ID = objectID
	if err := h.promptUseCase.UpdatePrompt(&prompt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, prompt)
}

// @Summary     Deletar prompt
// @Description Remove um prompt do sistema
// @Tags        prompts
// @Accept      json
// @Produce     json
// @Param       id path string true "ID do prompt"
// @Success     200 {object} map[string]string
// @Failure     400 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /prompts/{id} [delete]
func (h *Handler) DeletePrompt(c *gin.Context) {
	id := c.Param("id")
	if !primitive.IsValidObjectID(id) {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	if err := h.promptUseCase.DeletePrompt(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensagem": "Prompt deletado com sucesso"})
}
