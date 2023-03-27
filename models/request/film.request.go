package request

type FilmCreateRequest struct {
	Name 		string `json:"name" validate:"required"`
	JenisFilm 	string `json:"jenis_film" validate:"required"`
	KodeFilm	string `json:"kode_film" `
}

type FilmUpdateRequest struct {
	Name 		string `json:"name" validate:"required"`
	JenisFilm 	string `json:"jenis_film" validate:"required"`
	Produser	string `json:"produser"`
	KodeFilm	string `json:"kode_film"`
	Sutradara	string `json:"sutradara"`
	Penulis		string `json:"penulis"`
	Produksi	string `json:"produksi"`
	Casts		string `json:"casts"`
	Sinopsis	string `json:"sinopsis"`
}