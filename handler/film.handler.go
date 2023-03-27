package handler

import (
	"gofiber-api-gorm/database"
	"gofiber-api-gorm/models/entity"
	"gofiber-api-gorm/models/request"
	//"gofiber-api-gorm/utils"

	//"os/user"

	//"gofiber-api-gorm/models/response"
	"log"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func FilmHandlerRead(ctx *fiber.Ctx) error {
	var film []entity.Film
	result := database.DB.Find(&film)
	if result.Error != nil {
		log.Println(result.Error)
	}

	/* err := database.DB.Find(&film).Error
	if err != nil {
		log.Println(err)
	} */

	return ctx.JSON(film)
}

func FilmHandlerCreate(ctx *fiber.Ctx) error {
	/*var film entity.Film
	err := ctx.BodyParser(&film)
	if err != nil {
		log.Println(err)
	}

	validate := validator.New()
	errValidate := validate.Struct(film)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed to validate",
			"error":   errValidate.Error(),
		})
	}

	hashedPassword, err := utils.HashingPassword(film.KodeFilm)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed",
		})
	}

	newFilm.KodeFilm = hashedPassword

	newFilm := entity.Film{
		Name:		film.Name,
		JenisFilm:	film.JenisFilm,
		KodeFilm:	film.KodeFilm,
	}

	result := database.DB.Create(&film)
	if result.Error != nil {
		log.Println(result.Error)
	}

	return ctx.JSON(film)*/
	Film:= new(request.FilmCreateRequest)
	if err := ctx.BodyParser(Film); err != nil {
		return err
	}

	// VALIDASI REQUEST

	validate := validator.New()
	errValidate := validate.Struct(Film)
	if errValidate != nil {
		return ctx.Status (400).JSON(fiber.Map{ 
			"message": "failed to validate",
			"error": errValidate.Error(),
		})
	}

	/*newFilm := entity.Film{
		Name: 		Film.Name,
		KodeFilm:	Film.KodeFilm,
	}

	hashedPassword, err := utils.HashingPassword(Film.KodeFilm)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError). JSON(fiber.Map{ 
			"message": "internal server error",
		})
	}

	newFilm.KodeFilm = hashedPassword

	errCreateFilm := database.DB.Create(&newFilm). Error 
	if errCreateFilm != nil { 
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to create film",
		})
	}*/

	return ctx.JSON(fiber.Map{ 
		"message": "successfully",
		"data": Film,
	})
}

func FilmHandlerGetById(ctx *fiber.Ctx) error {
	/*var film entity.Film
	  result := database.DB.First(&film, ctx.Params("id"))
	  if result.Error!= nil {
	      log.Println(result.Error)
	  }

	  return ctx.JSON(film)
	*/

	// Yang Diatas Pake tabnine
	// Yang Dibawah Ikut Youtube

	filmId := ctx.Params("id")

	var film entity.Film
	err := database.DB.First(&film, "id = ?", filmId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "film not found",
		})
	}

	/*filmResponse := response.FilmResponse{
		ID:        film.ID,
		Name:      film.Name,
		JenisFilm: film.JenisFilm,
		Produser:  film.Produser,
		Sutradara: film.Sutradara,
	}*/

	return ctx.JSON(fiber.Map{
		"message": "succes",
		"data":    film,
	})
}

func FilmHandlerUpdate(ctx *fiber.Ctx) error {
	filmRequest := new(request.FilmUpdateRequest)
	if err := ctx.BodyParser(filmRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
    }

	var film entity.Film

	filmId := ctx.Params("id")
	// CHECK AVAILABLE FOR FILM
	err := database.DB.First(&film, "id =?", filmId).Error
	if err!= nil {
        return ctx.Status(404).JSON(fiber.Map{
            "message": "film not found",
        })
    }

	// UPDATE FILM
	if filmRequest.Name != ""{
		film.Name = filmRequest.Name
	}
	film.JenisFilm = filmRequest.JenisFilm
	film.KodeFilm = filmRequest.KodeFilm
	film.Sutradara = filmRequest.Sutradara

	errUpdate := database.DB.Save(&film).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "succes",
		"data":    film,
	})
}

func FilmHandlerDelete(ctx *fiber.Ctx) error {
	filmId := ctx.Params("id")
	
	var film entity.Film
	err := database.DB.First(&film, "id =?", filmId).Error
	if err!= nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "film not found",
		})
	}

	errDelete := database.DB.Delete(&film).Error
	if errDelete!= nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "succes",
	})
}