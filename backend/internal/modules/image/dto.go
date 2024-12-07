package image

import "github.com/google/uuid"

type CreateImageRequest struct {
	ProductID uuid.UUID `json:"product_id" binding:"required"`
	URL       string    `json:"url" binding:"required"`
}

type CreateImageResponse struct {
	ID        uuid.UUID `json:"id"`
	ProductID uuid.UUID `json:"product_id"`
	URL       string    `json:"url"`
}

type GetImageResponse struct {
	ID        uuid.UUID `json:"id"`
	ProductID uuid.UUID `json:"product_id"`
	URL       string    `json:"url"`
}

type GetAllImagesResponse struct {
	Images []GetImageResponse `json:"images"`
}

type DeleteImageResponse struct {
	Message string `json:"message"`
}
