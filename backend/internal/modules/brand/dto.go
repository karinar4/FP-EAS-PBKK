package brand

import "github.com/google/uuid"

type CreateBrandRequest struct {
	Name string `json:"name" binding:"required"`
}

type CreateBrandResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type GetBrandResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type GetAllBrandResponse struct {
	Brands []GetBrandResponse `json:"brands"`
}

type UpdateBrandRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateBrandResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type DeleteBrandResponse struct {
	Message string `json:"message"`
}
