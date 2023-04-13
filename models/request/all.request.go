package request

type FilmCreateRequest struct {
	Name 		string `json:"name" validate:"required"` 
	JenisFilm 	string `json:"jenis_film" validate:"required"`
	Produser	string `json:"produser" validate:"required"`
	Sutradara	string `json:"sutradara" validate:"required"`
	Penulis		string `json:"penulis" validate:"required"`
	Produksi	string `json:"produksi" validate:"required"`
	Casts		string `json:"casts" validate:"required"`
	Sinopsis	string `json:"sinopsis" validate:"required"`
	Like		uint `json:"like"`
}

type FilmUpdateRequest struct {
	Name 		string `json:"name" validate:"required"` 
	JenisFilm 	string `json:"jenis_film" validate:"required"`
	Produser	string `json:"produser" validate:"required"`
	Sutradara	string `json:"sutradara" validate:"required"`
	Penulis		string `json:"penulis" validate:"required"`
	Produksi	string `json:"produksi" validate:"required"`
	Casts		string `json:"casts" validate:"required"`
	Sinopsis	string `json:"sinopsis" validate:"required"`
	Like		uint `json:"like"`
}

type UserCreateRequest struct {
	Name 		string `json:"name" validate:"required"` 
	Email 		string 	`json:"email" validate:"required,email"`
	Password 	string 	`json:"password" validate:"required,min=6"`
	Address		string `json:"address"`
	Phone		string `json:"phone"`
}

type UserUpdateRequest struct {
	Name 		string `json:"name" validate:"required"` 
	Email 		string 	`json:"email" validate:"required"`
	Password 	string 	`json:"password" validate:"required"`
}

type UserEmailRequest struct {
	Email 		string 	`json:"email" validate:"required"`
}

type TheaterCreateRequest struct {
	Kota		string `json:"kota" validate:"required"`
	Theater		string `json:"theater" validate:"required"`
	Phone		string `json:"phone" validate:"required"`
}

type TheaterUpdateRequest struct {
	Kota		string `json:"kota"`
	Theater		string `json:"theater"`
	Phone		string `json:"phone"`
}

type TheaterListRequest struct {
	TheaterID	uint `json:"theater_id"`
	FilmID		uint `json:"film_id"`
}

type CommentCreateRequest struct {
	FilmID    uint      `json:"film_id"`
	Comment   string    `json:"comment"`
}