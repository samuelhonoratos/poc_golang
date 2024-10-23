package dto

type PostResponseDTO struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
