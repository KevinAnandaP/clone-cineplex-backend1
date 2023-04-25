package usecase

import (
	"context"
	"gofiber-api-gorm/models/entity"
	"gofiber-api-gorm/models/request"
)

type Usecase interface {
	CreateBook(ctx context.Context, req *request.FilmCreateRequest) error
	GetBookByID(ctx context.Context, id int64) (*entity.Film, error)
	GetAllBook(ctx context.Context, req *request.GetAllFilmRequest) (*entity.GetAllFilmResponse, error)
	UpdateBook(ctx context.Context, req *request.FilmUpdateRequest) error
	DeleteBook(ctx context.Context, id int64) error
}