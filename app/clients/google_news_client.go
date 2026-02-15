package clients

import (
	"fmt"
	"net/http"
	"net/url"
	"rm/app/data/google"
	"rm/app/utils"
	"time"

	"github.com/goravel/framework/facades"
)

type GoogleNewsClient struct {
	httpClient http.Client
	baseUrl    string
}

func NewGoogleNewsClient() *GoogleNewsClient {
	url := facades.Config().GetString("client.google_news.host")
	return &GoogleNewsClient{
		httpClient: http.Client{
			Timeout: 15 * time.Second,
		},
		baseUrl: url,
	}
}

func (c *GoogleNewsClient) GetNewsRss(
	query string,
) (*google.RSS, error) {
	encodedQuery := url.QueryEscape(query)
	resp, err := c.httpClient.Get(fmt.Sprintf("%s/rss/search?q=%s&hl=id&gl=ID&ceid=ID:id", c.baseUrl, encodedQuery))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch google news (%s)", resp.Status)
	}

	return utils.ParseXmlResponse[google.RSS](resp)
}
