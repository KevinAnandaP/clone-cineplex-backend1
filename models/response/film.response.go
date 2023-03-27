package response

type FilmResponse struct {
	ID			uint   `json:"id" gorm:"primary key"`
	Name 		string `json:"name" validate:"required"`
	JenisFilm 	string `json:"jenis_film" validate:"required"`
	Produser	string `json:"produser" validate:"required"`
	Sutradara	string `json:"sutradara"`
}