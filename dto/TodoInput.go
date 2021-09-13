package dto

type TodoInput struct {
	Description string `json:"description" binding:"required"`
}
