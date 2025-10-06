package server

import (
	"github.com/pretorian41/goaggregate/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func New() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())

	app.Use(func(ctx *fiber.Ctx) error {
		if err := ctx.Next(); err != nil {
			return ctx.Status(404).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return nil
	})

	api := app.Group("/api")
	routes.SetupAggRoutes(api)

	return app
}
