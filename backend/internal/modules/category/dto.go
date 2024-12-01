package category

import "github.com/google/uuid"

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

type CreateCategoryResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type GetCategoryResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type GetAllCategoryResponse struct {
	Categories []GetCategoryResponse `json:"categories"`
}

type UpdateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateCategoryResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type DeleteCategoryResponse struct {
	Message string `json:"message"`
}
