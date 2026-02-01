package clients

import (
	"fmt"
	"net/http"
	"rm/app/data/yahoo"
	"rm/app/utils"
	"time"

	"github.com/goravel/framework/facades"
)

type YahooClient struct {
	baseUrl    string
	httpClient http.Client
}

func NewYahooClient() *YahooClient {
	return &YahooClient{
		baseUrl: facades.Config().GetString("client.yahoo.host"),
		httpClient: http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

func (c *YahooClient) deriveRange(interval yahoo.RangeData) yahoo.RangeData {
	switch interval {

	case yahoo.ONE_DAY:
		return yahoo.ONE_DAY

	case yahoo.FIVE_DAYS:
		return yahoo.ONE_DAY

	case yahoo.ONE_MONTH:
		return yahoo.FIVE_DAYS

	case yahoo.THREE_MONTH:
		return yahoo.ONE_MONTH

	case yahoo.SIX_MONTH:
		return yahoo.THREE_MONTH

	case yahoo.ONE_YEAR:
		return yahoo.ONE_MONTH

	case yahoo.TWO_YEAR:
		return yahoo.THREE_MONTH

	case yahoo.FIVE_YEAR:
		return yahoo.SIX_MONTH

	case yahoo.TEN_YEAR:
		return yahoo.ONE_YEAR

	case yahoo.YTD:
		return yahoo.ONE_MONTH

	case yahoo.MAX:
		return yahoo.ONE_YEAR

	default:
		return yahoo.ONE_MONTH // safe fallback
	}
}

func (c *YahooClient) GetChart(
	stock string,
	interval yahoo.RangeData,
) (*yahoo.YahooChartResponse, error) {
	range_ := c.deriveRange(interval)
	url := fmt.Sprintf("%s/v8/finance/chart/%s.JK?interval=%s&range=%s", c.baseUrl, stock, interval, range_)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get chart (%s)", resp.Status)
	}

	return utils.ParseResponse[yahoo.YahooChartResponse](resp)
}
