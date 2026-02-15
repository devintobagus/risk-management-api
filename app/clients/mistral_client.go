package clients

import (
	"github.com/gage-technologies/mistral-go"
	"github.com/goravel/framework/facades"
)

type MistralClient struct {
	client mistral.MistralClient
}

func NewMistralClient() *MistralClient {
	apiKey := facades.Config().GetString("client.mistral.api_key")
	return &MistralClient{
		client: *mistral.NewMistralClientDefault(apiKey),
	}
}

func (c *MistralClient) ChatCompletions(
	messages []mistral.ChatMessage,
) (*mistral.ChatCompletionResponse, error) {
	return c.client.Chat("mistral-large-latest", messages, nil)
}

func (c *MistralClient) Embeddings(
	messages []string,
) (*mistral.EmbeddingResponse, error) {
	return c.client.Embeddings("mistral-embed", messages)

}
