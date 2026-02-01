package rm

import "rm/app/data/yahoo"

type ChartDataRequest struct {
	Stock string          `json:"stock" validate:"required"`
	Range yahoo.RangeData `json:"range" validate:"required"`
}

type ChartDataResponse struct {
	Close     []float64 `json:"close"`
	Timestamp []int64   `json:"timestamp"`
}
