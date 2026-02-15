package controllers

import (
	"rm/app/data/rm"
	"rm/app/services"
	"rm/app/traits"

	"github.com/goravel/framework/contracts/http"
)

type NewsController struct {
	traits.ApiResponse
	*traits.RequestValidator
	newsService services.NewsService
}

func NewNewsController() *NewsController {
	return &NewsController{
		newsService:      *services.NewNewsService(),
		RequestValidator: traits.NewRequestValidator(),
	}
}

func (r *NewsController) News(ctx http.Context) http.Response {
	var req rm.NewsRequest
	if err := r.BindAndValidate(ctx, &req); err != nil {
		return r.Error(ctx, 422, err.Error())
	}

	resp, err := r.newsService.News(req.Query)

	if err != nil {
		return r.Error(ctx, 400, err.Error())
	}

	return r.Success(ctx, resp, "success")
}
