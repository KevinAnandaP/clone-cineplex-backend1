package migration

import (
	"fmt"
	"gofiber-api-gorm/database"
	"gofiber-api-gorm/models/entity"
	"log"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.Film{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Database successfully migrated")
}