package services

import "rm/app/clients"

type AIService struct {
	geminiClient clients.GeminiClient
}

func NewAIService() *AIService {
	return &AIService{
		geminiClient: *clients.NewGeminiClient(),
	}
}

// TODO: need more research here
func (s *AIService) Reason(
	prompt string,
) {

}
