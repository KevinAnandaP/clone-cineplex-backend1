package migration

import (
	"fmt"
	"gofiber-api-gorm/database"
	"gofiber-api-gorm/models/entity"
	"log"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.User{}, &entity.Film{}, &entity.Comment{}, &entity.Theater{}, &entity.TheaterList{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Database successfully migrated")
}