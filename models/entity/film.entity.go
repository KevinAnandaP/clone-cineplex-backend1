package entity

type Film struct {
	ID			uint   `json:"id" gorm:"primary key"`
	Name 		string `json:"name" validate:"required"`
	JenisFilm 	string `json:"jenis_film" validate:"required"`
	Produser	string `json:"produser"`
	KodeFilm	string `json:"-" gorm:"column:password" validate:"required"`
	Sutradara	string `json:"sutradara"`
	Penulis		string `json:"penulis"`
	Produksi	string `json:"produksi"`
	Casts		string `json:"casts"`
	Sinopsis	string `json:"sinopsis"`
}