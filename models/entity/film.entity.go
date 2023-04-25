package entity

import (
	"gofiber-api-gorm/models/pg"
	"time"
)

type Film struct {
	ID        uint   `json:"id" gorm:"primary key"`
	Name      string `json:"name"`
	JenisFilm string `json:"jenis_film"`
	Produser  string `json:"produser"`
	Sutradara string `json:"sutradara"`
	Penulis   string `json:"penulis"`
	Produksi  string `json:"produksi"`
	Casts     string `json:"casts"`
	Sinopsis  string `json:"sinopsis"`
	Like      uint   `json:"like"`
	Cover	  string `json:"cover"`
}

type GetAllFilmResponse struct {
	Films		[]Film `json:"films"`
	Pagination  *pg.PaginationResponse `json:"pagination"`
}

type TheaterId struct {
	ID        uint   `json:"id" gorm:"primary key"`
	TheaterId uint   `json:"theater_id"`
	Name      string `json:"name"`
	JenisFilm string `json:"jenis_film"`
	Produser  string `json:"produser"`
	Sutradara string `json:"sutradara"`
	Penulis   string `json:"penulis"`
	Produksi  string `json:"produksi"`
	Casts     string `json:"casts"`
	Sinopsis  string `json:"sinopsis"`
	Like      uint   `json:"like"`
}

type Comment struct {
	ID        uint      `json:"id" gorm:"primary key"`
	FilmID    uint      `json:"film_id"`
	Film      Film      `gorm:"foreignKey:FilmID"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
