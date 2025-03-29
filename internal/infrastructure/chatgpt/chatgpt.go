package chatgpt

import (
	"context"
	"os"

	"github.com/sashabaranov/go-openai"
)

type ChatGPTService struct {
	client *openai.Client
}

func NewChatGPTService() *ChatGPTService {
	apiKey := os.Getenv("OPENAI_API_KEY")
	client := openai.NewClient(apiKey)
	return &ChatGPTService{client: client}
}

func (s *ChatGPTService) GenerateResponse(ctx context.Context, prompt string) (string, error) {
	resp, err := s.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
