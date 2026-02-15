package services

import (
	"rm/app/clients"
	"rm/app/constants"
	"rm/app/data/rm"
	"rm/app/utils"

	"github.com/gage-technologies/mistral-go"
)

type AIService struct {
	geminiClient  clients.GeminiClient
	mistralClient clients.MistralClient
}

func NewAIService() *AIService {
	return &AIService{
		geminiClient:  *clients.NewGeminiClient(),
		mistralClient: *clients.NewMistralClient(),
	}
}

func (s *AIService) NewsKeyword(
	prompt string,
) (*rm.NewsKeyword, error) {
	resp, err := s.mistralClient.ChatCompletions([]mistral.ChatMessage{
		{
			Role:    mistral.RoleSystem,
			Content: constants.NEWS_KEYWORD_PROMPT,
		},
		{
			Role:    mistral.RoleUser,
			Content: prompt,
		},
	})
	if err != nil {
		return nil, err
	}

	jsonStr := utils.ExtractJSON(resp.Choices[0].Message.Content)

	return utils.ParseFromBytes[rm.NewsKeyword]([]byte(jsonStr))
}
