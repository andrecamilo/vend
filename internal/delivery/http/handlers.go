package http

import (
	"net/http"
	"strconv"
	"vend/internal/domain"
	"vend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	pessoaUseCase *usecase.PessoaUseCase
}

func NewHandler(pessoaUseCase *usecase.PessoaUseCase) *Handler {
	return &Handler{
		pessoaUseCase: pessoaUseCase,
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
// @Param       id path int true "ID da pessoa"
// @Success     200 {object} domain.Pessoa
// @Failure     400 {object} map[string]string
// @Failure     404 {object} map[string]string
// @Router      /pessoas/{id} [get]
func (h *Handler) GetPessoa(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	pessoa, err := h.pessoaUseCase.GetPessoa(uint(id))
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
// @Param       id path int true "ID da pessoa"
// @Param       pessoa body domain.Pessoa true "Dados da pessoa"
// @Success     200 {object} domain.Pessoa
// @Failure     400 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /pessoas/{id} [put]
func (h *Handler) UpdatePessoa(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	var pessoa domain.Pessoa
	if err := c.ShouldBindJSON(&pessoa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	pessoa.ID = uint(id)
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
// @Param       id path int true "ID da pessoa"
// @Success     200 {object} map[string]string
// @Failure     400 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /pessoas/{id} [delete]
func (h *Handler) DeletePessoa(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	if err := h.pessoaUseCase.DeletePessoa(uint(id)); err != nil {
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
// @Failure     501 {object} map[string]string
// @Router      /telefones [get]
func (h *Handler) ListTelefones(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"erro": "Método não implementado"})
}

// @Summary     Criar telefone
// @Description Cria um novo telefone no sistema
// @Tags        telefones
// @Accept      json
// @Produce     json
// @Param       telefone body domain.Telefone true "Dados do telefone"
// @Success     201 {object} domain.Telefone
// @Failure     501 {object} map[string]string
// @Router      /telefones [post]
func (h *Handler) CreateTelefone(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"erro": "Método não implementado"})
}

// @Summary     Buscar telefone
// @Description Retorna os dados de um telefone específico
// @Tags        telefones
// @Accept      json
// @Produce     json
// @Param       id path int true "ID do telefone"
// @Success     200 {object} domain.Telefone
// @Failure     501 {object} map[string]string
// @Router      /telefones/{id} [get]
func (h *Handler) GetTelefone(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"erro": "Método não implementado"})
}

// @Summary     Atualizar telefone
// @Description Atualiza os dados de um telefone específico
// @Tags        telefones
// @Accept      json
// @Produce     json
// @Param       id path int true "ID do telefone"
// @Param       telefone body domain.Telefone true "Dados do telefone"
// @Success     200 {object} domain.Telefone
// @Failure     501 {object} map[string]string
// @Router      /telefones/{id} [put]
func (h *Handler) UpdateTelefone(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"erro": "Método não implementado"})
}

// @Summary     Deletar telefone
// @Description Remove um telefone do sistema
// @Tags        telefones
// @Accept      json
// @Produce     json
// @Param       id path int true "ID do telefone"
// @Success     200 {object} map[string]string
// @Failure     501 {object} map[string]string
// @Router      /telefones/{id} [delete]
func (h *Handler) DeleteTelefone(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"erro": "Método não implementado"})
}

// @Summary     Listar contextos
// @Description Retorna a lista de todos os contextos cadastrados
// @Tags        contextos
// @Accept      json
// @Produce     json
// @Success     200 {array} domain.Contexto
// @Failure     501 {object} map[string]string
// @Router      /contextos [get]
func (h *Handler) ListContextos(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"erro": "Método não implementado"})
}

// @Summary     Criar contexto
// @Description Cria um novo contexto no sistema
// @Tags        contextos
// @Accept      json
// @Produce     json
// @Param       contexto body domain.Contexto true "Dados do contexto"
// @Success     201 {object} domain.Contexto
// @Failure     501 {object} map[string]string
// @Router      /contextos [post]
func (h *Handler) CreateContexto(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"erro": "Método não implementado"})
}

// @Summary     Buscar contexto
// @Description Retorna os dados de um contexto específico
// @Tags        contextos
// @Accept      json
// @Produce     json
// @Param       id path int true "ID do contexto"
// @Success     200 {object} domain.Contexto
// @Failure     501 {object} map[string]string
// @Router      /contextos/{id} [get]
func (h *Handler) GetContexto(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"erro": "Método não implementado"})
}

// @Summary     Atualizar contexto
// @Description Atualiza os dados de um contexto específico
// @Tags        contextos
// @Accept      json
// @Produce     json
// @Param       id path int true "ID do contexto"
// @Param       contexto body domain.Contexto true "Dados do contexto"
// @Success     200 {object} domain.Contexto
// @Failure     501 {object} map[string]string
// @Router      /contextos/{id} [put]
func (h *Handler) UpdateContexto(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"erro": "Método não implementado"})
}

// @Summary     Deletar contexto
// @Description Remove um contexto do sistema
// @Tags        contextos
// @Accept      json
// @Produce     json
// @Param       id path int true "ID do contexto"
// @Success     200 {object} map[string]string
// @Failure     501 {object} map[string]string
// @Router      /contextos/{id} [delete]
func (h *Handler) DeleteContexto(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"erro": "Método não implementado"})
}

// @Summary     Listar prompts
// @Description Retorna a lista de todos os prompts cadastrados
// @Tags        prompts
// @Accept      json
// @Produce     json
// @Success     200 {array} domain.Prompt
// @Failure     501 {object} map[string]string
// @Router      /prompts [get]
func (h *Handler) ListPrompts(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"erro": "Método não implementado"})
}

// @Summary     Criar prompt
// @Description Cria um novo prompt no sistema
// @Tags        prompts
// @Accept      json
// @Produce     json
// @Param       prompt body domain.Prompt true "Dados do prompt"
// @Success     201 {object} domain.Prompt
// @Failure     501 {object} map[string]string
// @Router      /prompts [post]
func (h *Handler) CreatePrompt(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"erro": "Método não implementado"})
}

// @Summary     Buscar prompt
// @Description Retorna os dados de um prompt específico
// @Tags        prompts
// @Accept      json
// @Produce     json
// @Param       id path int true "ID do prompt"
// @Success     200 {object} domain.Prompt
// @Failure     501 {object} map[string]string
// @Router      /prompts/{id} [get]
func (h *Handler) GetPrompt(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"erro": "Método não implementado"})
}

// @Summary     Atualizar prompt
// @Description Atualiza os dados de um prompt específico
// @Tags        prompts
// @Accept      json
// @Produce     json
// @Param       id path int true "ID do prompt"
// @Param       prompt body domain.Prompt true "Dados do prompt"
// @Success     200 {object} domain.Prompt
// @Failure     501 {object} map[string]string
// @Router      /prompts/{id} [put]
func (h *Handler) UpdatePrompt(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"erro": "Método não implementado"})
}

// @Summary     Deletar prompt
// @Description Remove um prompt do sistema
// @Tags        prompts
// @Accept      json
// @Produce     json
// @Param       id path int true "ID do prompt"
// @Success     200 {object} map[string]string
// @Failure     501 {object} map[string]string
// @Router      /prompts/{id} [delete]
func (h *Handler) DeletePrompt(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"erro": "Método não implementado"})
}
