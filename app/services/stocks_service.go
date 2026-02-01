package services

import (
	"rm/app/clients"
	"rm/app/data/appscript"
	"rm/app/data/rm"
)

type StocksService struct {
	yahooClient      clients.YahooClient
	appsscriptClient clients.AppScriptClient
}

func NewStocksServices() *StocksService {
	return &StocksService{
		yahooClient:      *clients.NewYahooClient(),
		appsscriptClient: *clients.NewAppScriptClient(),
	}
}

func (s *StocksService) GetChart(
	request rm.ChartDataRequest,
) (*rm.ChartDataResponse, error) {
	stock := request.Stock
	interval := request.Range

	resp, err := s.yahooClient.GetChart(
		stock,
		interval,
	)

	if err != nil {
		return nil, err
	}

	var data rm.ChartDataResponse
	for _, ts := range resp.Chart.Result[0].Timestamp {

		data.Timestamp = append(data.Timestamp, ts)
	}

	for _, close := range resp.Chart.Result[0].Indicators.Quote[0].Close {
		data.Close = append(data.Close, close)
	}

	return &data, nil
}

func (s *StocksService) GetStockList() (*appscript.StockListResponse, error) {
	resp, err := s.appsscriptClient.GetStockList()

	if err != nil {
		return nil, err
	}

	return resp, nil
}
