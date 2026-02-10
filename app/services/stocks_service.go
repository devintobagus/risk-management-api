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
) (*[]rm.ChartResponseData, error) {
	stock := request.Stock
	interval := request.Range

	resp, err := s.yahooClient.GetChart(
		stock,
		interval,
	)

	if err != nil {
		return nil, err
	}

	var data []rm.ChartResponseData
	for i := range len(resp.Chart.Result[0].Timestamp) {
		data_ := rm.ChartResponseData{
			Close:     resp.Chart.Result[0].Indicators.Quote[0].Close[i],
			Timestamp: resp.Chart.Result[0].Timestamp[i],
		}
		data = append(data, data_)
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
