package request

type LoginRequest struct {
	Name 		string `json:"name" validate:"required"`
	KodeFilm	string `json:"kode_film"`
}