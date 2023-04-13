package entity

import (
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
	FilmID    uint      `json:"film_Id"`
	Film      Film      `gorm:"foreignKey:FilmID"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
