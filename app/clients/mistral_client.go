package clients

import (
	"rm/app/data/mistralai"

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
	model mistralai.MistralModel,
	messages []mistral.ChatMessage,
	requestParams *mistral.ChatRequestParams,
) (*mistral.ChatCompletionResponse, error) {
	return c.client.Chat(string(model), messages, requestParams)
}

func (c *MistralClient) Embeddings(
	messages []string,
) (*mistral.EmbeddingResponse, error) {
	return c.client.Embeddings("mistral-embed", messages)

}
