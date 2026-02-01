package controllers

import (
	"rm/app/data/rm"
	"rm/app/services"
	"rm/app/traits"

	"github.com/goravel/framework/contracts/http"
)

type StockController struct {
	traits.ApiResponse
	*traits.RequestValidator
	stockService services.StocksService
}

func NewStockController() *StockController {
	return &StockController{
		stockService:     *services.NewStocksServices(),
		RequestValidator: traits.NewRequestValidator(),
	}
}

func (r *StockController) Chart(ctx http.Context) http.Response {
	var req rm.ChartDataRequest

	if err := r.BindAndValidate(ctx, &req); err != nil {
		return r.Error(ctx, 422, err.Error())
	}

	resp, err := r.stockService.GetChart(
		req,
	)
	if err != nil {
		return r.Error(ctx, 400, err.Error())
	}

	return r.Success(ctx, resp, "success")
}
