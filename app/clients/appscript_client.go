package clients

import (
	"fmt"
	"net/http"
	"rm/app/data/appscript"
	"rm/app/utils"
	"time"

	"github.com/goravel/framework/facades"
)

type AppScriptClient struct {
	httpClient   http.Client
	stockListUrl string
}

func NewAppScriptClient() *AppScriptClient {
	return &AppScriptClient{
		stockListUrl: facades.Config().GetString("client.appscript.stock_list_url"),
		httpClient: http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

func (c *AppScriptClient) GetStockList() (*appscript.StockListResponse, error) {
	resp, err := c.httpClient.Get(c.stockListUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get stock (%s)", resp.Status)
	}

	return utils.ParseResponse[appscript.StockListResponse](resp)
}
