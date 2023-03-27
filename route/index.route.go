package route

import (
	"gofiber-api-gorm/config"
	"gofiber-api-gorm/handler"
	"gofiber-api-gorm/middleware"

	"github.com/gofiber/fiber/v2"
)


func RouteInit(r *fiber.App) {
	r.Get("/", func(ctx *fiber.Ctx) error {

		return ctx.JSON(fiber.Map{
			"hello": "world",
		})
	})
	r.Static("/public", config.ProjectRootPath+"/public/asset")

	r.Post("/login", handler.LoginHandler)

	r.Get("/film", middleware.Auth, handler.FilmHandlerRead)
	r.Post("/film", handler.FilmHandlerCreate)
	r.Get("/film/:id", handler.FilmHandlerGetById)
	r.Put("/film/:id", handler.FilmHandlerUpdate)
	r.Delete("/film/:id", handler.FilmHandlerDelete)
}