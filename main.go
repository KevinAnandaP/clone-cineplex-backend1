package main

import (
	"gofiber-api-gorm/database"
	"gofiber-api-gorm/migration"
	"gofiber-api-gorm/route"
	//"database/sql"
	//"fmt"
	//"github.com/lib/pq"

	"github.com/gofiber/fiber/v2"
)

func main() {

	database.DatabaseInit()
	migration.RunMigration()
	app := fiber.New()
	
	// INITAL ROUTE
	route.RouteInit(app)

	app.Listen(":8080")
}
