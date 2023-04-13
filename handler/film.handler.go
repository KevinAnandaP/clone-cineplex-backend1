package handler

import (
	"gofiber-api-gorm/database"
	"gofiber-api-gorm/models/entity"
	"gofiber-api-gorm/models/request"
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
	return ctx.JSON(film)
}

func FilmHandlerGetByTheaterId(ctx *fiber.Ctx) error {
	theaterid := ctx.QueryInt("theaterid")
    var film []entity.TheaterId
    err := database.DB.Raw(`
        SELECT f.id, f.name, l.theater_id AS theater_id, f.jenis_film, f. produser, f.sutradara, f.penulis, f.produksi, f.casts, f.sinopsis, f.like
        FROM films f
        INNER JOIN theater_lists l ON l.film_id = f.id
        WHERE l.theater_id = ?`, theaterid).Scan(&film)

	if err.Error != nil {
		log.Println(err.Error)
	}
	return ctx.JSON(film)
}

func FilmHandlerGetById(ctx *fiber.Ctx) error {
	filmId := ctx.Params("id")

	var film entity.Film
	err := database.DB.First(&film, "id = ?", filmId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "film not found",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "successfully",
		"data":    film,
	})
}

func FilmHandlerCreate(ctx *fiber.Ctx) error {
	film := new(request.FilmCreateRequest)
	if err := ctx.BodyParser(film); err != nil {
		return err
	}

	// VALIDASI REQUEST
	validate := validator.New()
	errValidate := validate.Struct(film)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed to validate",
			"error":   errValidate.Error(),
		})
	}

	newFilm := entity.Film{
		Name:      film.Name,
		JenisFilm: film.JenisFilm,
		Produser:  film.Produser,
		Sutradara: film.Sutradara,
		Penulis:   film.Penulis,
		Produksi:  film.Produksi,
		Casts:     film.Casts,
		Sinopsis:  film.Sinopsis,
		Like:      film.Like,
	}

	errCreateUser := database.DB.Create(&newFilm).Error
	if errCreateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to create user",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "successfully",
		"data":    newFilm,
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
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "film not found",
		})
	}

	// UPDATE FILM
	if filmRequest.Name != "" {
		film.Name = filmRequest.Name
	}
	film.JenisFilm = filmRequest.JenisFilm
	film.Produksi = filmRequest.Produksi
	film.Sutradara = filmRequest.Sutradara
	film.Penulis = filmRequest.Penulis
	film.Produksi = filmRequest.Produksi
	film.Casts = filmRequest.Casts
	film.Sinopsis = filmRequest.Sinopsis
	film.Like = filmRequest.Like

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
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "film not found",
		})
	}

	errDelete := database.DB.Delete(&film).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "succes",
	})
}

func CommentHandlerGetComments(ctx *fiber.Ctx) error {
	filmId := ctx.QueryInt("filmId")
    var film []entity.Film
    err := database.DB.Raw(`
		SELECT f.id, f.name, f.jenis_film, f. produser, f.sutradara, f.penulis, f.produksi, f.casts, f.sinopsis, f.like, c.comment
		FROM films f
		INNER JOIN comments c ON c.film_id = f.id
		WHERE c.film_id = ?`, filmId).Scan(&film)

    if err.Error != nil{
        log.Println(err.Error)
    }

	return ctx.JSON(film)
}

func CommentHandlerCreate(ctx *fiber.Ctx) error {
	Comment := new(request.CommentCreateRequest)
	if err := ctx.BodyParser(Comment); err != nil {
		return err
	}

	// VALIDASI REQUEST
 	validate := validator.New()
	errValidate := validate.Struct(Comment)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message" : "failed to validate",
			"error" : errValidate.Error(),
		})
	}

	newComment := entity.Comment{
		FilmID:		Comment.FilmID,
		Comment:		Comment.Comment,
	}

	errCreateComment := database.DB.Create(&newComment).Error
	if errCreateComment != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to create comment",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "successfully",
		"data": newComment,
	})
}

func CommentHandlerDelete(ctx *fiber.Ctx) error {
	commentid := ctx.Params("id")
	var comment entity.Comment

	// CHECK AVAILABLE COMMENT
	err := database.DB.Debug().First(&comment, "id=?" ,&commentid).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "comment Not Found",
		})
	}

	errDelete := database.DB.Debug().Delete(&comment).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "comment deleted",
	})
}