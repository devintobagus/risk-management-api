package services

import (
	"rm/app/clients"
	"rm/app/data/rm"
	"time"
)

type StocksService struct {
	yahooClient *clients.YahooClient
}

func NewStocksServices() *StocksService {
	return &StocksService{
		yahooClient: &clients.YahooClient{},
	}
}

func (s *StocksService) GetChart(
	request rm.ChartDataRequest,
) (*rm.ChartDataResponse, error) {
	stock := request.Stock
	interval := request.Interval

	resp, err := s.yahooClient.GetChart(
		stock,
		interval,
	)

	if err != nil {
		return nil, err
	}

	var data rm.ChartDataResponse
	for _, ts := range resp.Chart.Result[0].Timestamp {
		t := time.Unix(ts, 0).In(time.FixedZone("WIB", 7*3600))

		data.Timestamp = append(data.Timestamp, t.String())
	}

	for _, close := range resp.Chart.Result[0].Indicators.Quote[0].Close {
		data.Close = append(data.Close, close)
	}

	return &data, nil
}
