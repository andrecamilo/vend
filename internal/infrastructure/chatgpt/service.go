package chatgpt

import (
	"context"
	"vend/internal/domain"

	"github.com/sashabaranov/go-openai"
)

type Service struct {
	client *openai.Client
}

func NewService(apiKey string) *Service {
	client := openai.NewClient(apiKey)
	return &Service{client: client}
}

func (s *Service) GenerateResponse(ctx context.Context, prompt *domain.Prompt) (string, error) {
	resp, err := s.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: prompt.Conteudo,
				},
			},
			Temperature: 0.7,
		},
	)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

func (s *Service) GenerateContextualResponse(ctx context.Context, contexto *domain.Contexto, prompt *domain.Prompt) (string, error) {
	systemMessage := "Contexto: " + contexto.Descricao + "\n"
	systemMessage += "Período: " + contexto.DataInicio.Format("2006-01-02") + " até " + contexto.DataFim.Format("2006-01-02") + "\n"
	systemMessage += "Pessoas envolvidas:\n"
	for _, pessoa := range contexto.Pessoas {
		systemMessage += "- " + pessoa.Nome + " (" + pessoa.Email + ")\n"
	}

	resp, err := s.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: systemMessage,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt.Conteudo,
				},
			},
			Temperature: 0.7,
		},
	)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
