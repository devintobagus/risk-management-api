package rm

import "rm/app/data/yahoo"

type ChartDataRequest struct {
	Stock    string          `json:"stock" validate:"required"`
	Interval yahoo.RangeData `json:"interval" validate:"required"`
}

type ChartDataResponse struct {
	Close     []float64
	Timestamp []string
}
