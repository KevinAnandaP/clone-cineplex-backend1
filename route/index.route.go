package route

import (
	"gofiber-api-gorm/config"
	"gofiber-api-gorm/handler"
	"gofiber-api-gorm/middleware"

	"github.com/gofiber/fiber/v2"
)


func RouteInit(r *fiber.App) {
	r.Static("/public", config.ProjectRootPath+"/public/asset")

	r.Post("/login", handler.LoginHandler)

	r.Get("/film", handler.FilmHandlerRead)
	r.Post("/film", handler.FilmHandlerCreate)
	r.Get("/film/:id", handler.FilmHandlerGetById)
	r.Put("/film/:id", handler.FilmHandlerUpdate)
	r.Delete("/film/:id", handler.FilmHandlerDelete)
	r.Get("theater/theaterlist", handler.FilmHandlerGetByTheaterId)
	r.Get("/comments", middleware.Auth, handler.CommentHandlerGetComments)
	r.Post("/comment",middleware.Auth, handler.CommentHandlerCreate)
	r.Delete("/comment/:id",middleware.Auth, handler.CommentHandlerDelete)

	r.Get("/user", middleware.Auth, handler.UserHandlerRead)
	r.Get("/user/:id", handler.UserHandlerGetById)
	r.Post("/user", handler.UserHandlerCreate)
	r.Put("/user/:id", handler.UserHandlerUpdate)
	r.Put("/user/:id/update-email", handler.UserHandlerUpdateEmail)
	r.Delete("/user/:id", handler.UserHandlerDelete)

	r.Get("/theater", handler.TheaterHandlerRead)
	r.Post("/theater", handler.TheaterHandlerCreate)
	r.Post("/theaterlist", handler.TheaterHandlerCreateTheaterList)
	r.Get("/theater/:kota", handler.TheaterHandlerGetById)
	r.Put("/theater/:id", handler.TheaterHandlerUpdate)
	r.Delete("/theater/:id", handler.TheaterHandlerDelete)
	r.Get("/theaterdetails", handler.TheaterHandlerGetDetails)
}