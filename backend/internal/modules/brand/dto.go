package brand

import "github.com/google/uuid"

type CreateBrandRequest struct {
	Name string `json:"name" binding:"required"`
	Origin string `json:"origin" binding:"required"`
}

type CreateBrandResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Origin string    `json:"origin"`
}

type GetBrandResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Origin string    `json:"origin"`
}

type GetAllBrandResponse struct {
	Brands []GetBrandResponse `json:"brands"`
}

type UpdateBrandRequest struct {
	Name string `json:"name" binding:"required"`
	Origin string `json:"origin" binding:"required"`
}

type UpdateBrandResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Origin string    `json:"origin"`
}

type DeleteBrandResponse struct {
	Message string `json:"message"`
}
