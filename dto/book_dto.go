package dto

type BooksDTO struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Quantity    int        `json:"qty"`
	Cover       string     `json:"cover"`
	AuthorId    int        `json:"author_id"`
	Authors     AuthorsDTO `json:"authors"`
}

type BookPayload struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Quantity    int    `json:"qty" binding:"required"`
	Cover       string `json:"cover"`
	AuthorId    int    `json:"author_id" binding:"required"`
}

type BookResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Quantity    int    `json:"qty"`
	Cover       string `json:"cover"`
}
