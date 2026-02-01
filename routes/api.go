package routes

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"

	"rm/app/http/controllers"
)

func Api() {
	userController := controllers.NewUserController()
	stockController := controllers.NewStockController()
	facades.Route().Prefix("api").Group(func(router route.Router) {
		router.Prefix("v1").Group(func(router route.Router) {
			router.Prefix("stock").Group(func(router route.Router) {
				router.Post("chart", stockController.Chart)
				router.Get("chart", stockController.List)
			})
		})
	})
	facades.Route().Get("/users/{id}", userController.Show)
}
