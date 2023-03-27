package handler

import (
	"gofiber-api-gorm/database"
	"gofiber-api-gorm/models/entity"
	"gofiber-api-gorm/models/request"
	"gofiber-api-gorm/utils"
	"log"
	"time"

	//"github.com/gofiber/fiber/v2/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func LoginHandler(ctx *fiber.Ctx) error {
	loginRequest := new(request.LoginRequest)
	if err := ctx.BodyParser(loginRequest); err != nil {
		return err
    }
	log.Println(loginRequest)
	
	validate := validator.New()
	errValidate := validate.Struct(loginRequest)
	if errValidate!= nil {
        return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error": errValidate.Error(),
		})
    }

	var film entity.Film
	err := database.DB.First(&film, "kode_film = ?", loginRequest.KodeFilm).Error
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credentials",
		})
	}

	//
	isValid := utils.CheckPasswordHash(loginRequest.KodeFilm, film.KodeFilm)
	if !isValid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "message": "wrong credentials",
        })
    }

	//GEN JWT
	claims := jwt.MapClaims{}
	claims["name"] = film.Name
	claims["kode_film"] = film.KodeFilm
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()

	token, errGenerateToken := utils.GenerateToken(&claims)
	if errGenerateToken!= nil {
		log.Println(errGenerateToken)
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credentials",
        })
	}

	return ctx.JSON(fiber.Map{
		"token": token,
    })
}