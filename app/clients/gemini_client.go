package clients

import (
	"context"

	"github.com/goravel/framework/facades"
	"google.golang.org/genai"
)

type GeminiClient struct {
	apiKey string
	client *genai.Client
	ctx    context.Context
}

func NewGeminiClient() *GeminiClient {
	ctx := context.Background()
	client, _ := genai.NewClient(ctx, nil)
	return &GeminiClient{
		apiKey: facades.Config().GetString("client.gemini.api_key"),
		client: client,
		ctx:    ctx,
	}
}

func (c *GeminiClient) Embeddings(prompt string) (*[]float32, error) {
	contents := []*genai.Content{
		genai.NewContentFromText(prompt, genai.RoleUser),
	}

	results, err := c.client.Models.EmbedContent(c.ctx,
		"gemini-embedding-001",
		contents,
		nil,
	)

	if err != nil {
		return nil, err
	}

	return &results.Embeddings[0].Values, nil
}

func (c *GeminiClient) ChatCompletion(prompt string) (*string, error) {
	result, err := c.client.Models.GenerateContent(
		c.ctx,
		"gemini-3-flash-preview",
		genai.Text(prompt),
		nil,
	)

	if err != nil {
		return nil, err
	}
	response := result.Text()

	return &response, nil
}
