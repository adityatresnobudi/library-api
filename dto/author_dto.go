package dto

type AuthorsDTO struct {
	ID   int    `json:"-"`
	Name string `json:"name"`
}
