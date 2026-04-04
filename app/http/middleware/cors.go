package middleware

import (
	"github.com/goravel/framework/contracts/http"
)

func Cors() http.Middleware {
	return func(ctx http.Context) {
		ctx.Response().Header("Access-Control-Allow-Origin", "*")
		ctx.Response().Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		ctx.Response().Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept")

		// Handle preflight OPTIONS request
		if ctx.Request().Method() == "OPTIONS" {
			ctx.Request().AbortWithStatus(204)
			return
		}

		ctx.Request().Next()
	}
}
