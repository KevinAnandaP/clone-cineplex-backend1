package handler

import (
	"gofiber-api-gorm/database"
	"gofiber-api-gorm/models/entity"
	"gofiber-api-gorm/models/request"
	"log"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func TheaterHandlerRead(ctx *fiber.Ctx) error {
	var Theater []entity.Theater
	result := database.DB.Find(&Theater)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(Theater)
}

func TheaterHandlerGetById(ctx *fiber.Ctx) error {
	TheaterId := ctx.Params("kota")

	var Theater entity.Theater
	err := database.DB.First(&Theater, "kota = ?", TheaterId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "theater does not exist",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "succes",
		"data":    Theater,
	})
}

func TheaterHandlerGetDetails(ctx *fiber.Ctx) error {
	theaterid := ctx.QueryInt("theaterid")
    var theater []entity.TheaterDetails
    err := database.DB.Raw(`
		SELECT theaters.id, theaters.kota, theaters.theater, theaters.phone, films.id, films.name, films.jenis_film, films.produser, films.sutradara, films.penulis, films.produksi, films.casts, films.sinopsis,
		FROM theaters, films
		INNER JOIN theater_lists l ON l.film_id = films.id
		WHERE theaters.id = ?`, theaterid).Scan(&theater)

    if err.Error != nil{
        log.Println(err.Error)
    }

	var film []entity.TheaterId

    return ctx.JSON(fiber.Map{
		"message": "successfully",
		"data": film,
	})
}

func TheaterHandlerCreate(ctx *fiber.Ctx) error {
	Theater := new(request.TheaterCreateRequest)
	if err := ctx.BodyParser(Theater); err != nil {
		return err
	}

	// VALIDASI REQUEST
	validate := validator.New()
	errValidate := validate.Struct(Theater)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed to validate",
			"error":   errValidate.Error(),
		})
	}

	newTheater := entity.Theater{
		Kota:    Theater.Kota,
		Theater: Theater.Theater,
		Phone:   Theater.Phone,
	}

	errCreateTheater := database.DB.Create(&newTheater).Error
	if errCreateTheater != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to create Theater",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "successfully",
		"data":    newTheater,
	})
}

func TheaterHandlerUpdate(ctx *fiber.Ctx) error {
	theaterRequest := new(request.TheaterUpdateRequest)
	if err := ctx.BodyParser(theaterRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var Theater entity.Theater

	TheaterId := ctx.Params("id")
	// CHECK AVAILABLE FOR Theater
	err := database.DB.First(&Theater, "id =?", TheaterId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "theater not found",
		})
	}

	// UPDATE Theater
	if theaterRequest.Kota != "" {
		Theater.Kota = theaterRequest.Kota
	}
	Theater.Kota = theaterRequest.Kota
	Theater.Theater = theaterRequest.Theater
	Theater.Phone = theaterRequest.Phone

	errUpdate := database.DB.Save(&Theater).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "succes",
		"data":    Theater,
	})
}

func TheaterHandlerDelete(ctx *fiber.Ctx) error {
	TheaterId := ctx.Params("id")

	var Theater entity.Theater
	err := database.DB.First(&Theater, "id =?", TheaterId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Theater not found",
		})
	}

	errDelete := database.DB.Delete(&Theater).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "succes",
	})
}
