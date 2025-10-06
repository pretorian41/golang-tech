package routes

import (
	"github.com/pretorian41/goaggregate/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupAggRoutes(api fiber.Router) {
	apiProduct := api.Group("/agg")
	apiProduct.Get("/1", controllers.GetSourceFirst)
	apiProduct.Get("/2", controllers.GetSourceSecond)
	apiProduct.Get("/3", controllers.GetSourceThird)
	apiProduct.Get("/fetch/:id", controllers.GetLoadAggregate)
}
