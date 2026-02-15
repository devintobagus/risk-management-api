package controllers

import (
	"rm/app/data/rm"
	"rm/app/services"
	"rm/app/traits"

	"github.com/goravel/framework/contracts/http"
)

type AiController struct {
	traits.ApiResponse
	*traits.RequestValidator
	aiService services.AIService
}

func NewAiController() *AiController {
	return &AiController{
		aiService:        *services.NewAIService(),
		RequestValidator: traits.NewRequestValidator(),
	}
}

func (r *AiController) NewsKeyword(ctx http.Context) http.Response {
	var req rm.NewsKeywordRequest
	if err := r.BindAndValidate(ctx, &req); err != nil {
		return r.Error(ctx, 422, err.Error())
	}

	resp, err := r.aiService.NewsKeyword(req.Query)

	if err != nil {
		return r.Error(ctx, 400, err.Error())
	}

	return r.Success(ctx, resp, "success")
}
