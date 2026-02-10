package traits

import (
	"time"

	"github.com/goravel/framework/contracts/http"
)

type DefaultResponse struct {
	Timestamp string      `json:"timestamp"`
	Status    int         `json:"status"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Total     *int64      `json:"total,omitempty"`
}

type ApiResponse struct{}

func (a *ApiResponse) JSON(ctx http.Context, status int, message string, data interface{}) http.Response {
	return ctx.Response().Json(status, DefaultResponse{
		Timestamp: time.Now().Format(time.RFC3339),
		Status:    status,
		Message:   message,
		Data:      data,
	})
}

func (a *ApiResponse) Success(ctx http.Context, data interface{}, message string) http.Response {
	return a.JSON(ctx, 200, message, data)
}

func (a *ApiResponse) Error(ctx http.Context, code int, message string) http.Response {
	return a.JSON(ctx, code, message, nil)
}

func (a *ApiResponse) SuccessList(ctx http.Context, data interface{}, total *int64, message string) http.Response {
	return ctx.Response().Json(200, DefaultResponse{
		Timestamp: time.Now().Format(time.RFC3339),
		Status:    200,
		Message:   message,
		Data:      data,
		Total:     total,
	})
}

func (a *ApiResponse) Unauthorized(ctx http.Context) {
	ctx.Request().AbortWithStatus(401)
}
