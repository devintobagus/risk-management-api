package routes

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"

	"rm/app/http/controllers"
)

func Api() {
	aiController := controllers.NewAiController()
	newsController := controllers.NewNewsController()
	stockController := controllers.NewStockController()
	facades.Route().Prefix("api").Group(func(router route.Router) {
		router.Prefix("v1").Group(func(router route.Router) {
			router.Prefix("stock").Group(func(router route.Router) {
				router.Post("chart", stockController.Chart)
				// router.Get("chart", stockController.List)
			})
			router.Prefix("news").Group(func(router route.Router) {
				router.Get("list", newsController.News)
			})
			router.Prefix("ai").Group(func(router route.Router) {
				router.Get("news-keyword", aiController.NewsKeyword)
			})
		})
	})
}
