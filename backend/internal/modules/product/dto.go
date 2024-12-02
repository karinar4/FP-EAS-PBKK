package product

import (
	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/brand"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/category"
)

type CreateProductRequest struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	Price       float64   `json:"price" binding:"required"`
	Stock       int       `json:"stock" binding:"required"`
	CategoryID  uuid.UUID `json:"category_id" binding:"required"`
	BrandID     uuid.UUID `json:"brand_id" binding:"required"`
}

type CreateProductResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
}

type GetProductResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	//CategoryID  uuid.UUID                    `json:"category_id"`
	//BrandID     uuid.UUID                    `json:"brand_id"`
	Category category.GetCategoryResponse `json:"category"`
	Brand    brand.GetBrandResponse       `json:"brand"`
}

type GetAllProductsResponse struct {
	Products []GetProductResponse `json:"products"`
}

type UpdateProductRequest struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	CategoryID  uuid.UUID `json:"category_id"`
	BrandID     uuid.UUID `json:"brand_id"`
}

type UpdateProductResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	CategoryID  uuid.UUID `json:"category_id"`
	BrandID     uuid.UUID `json:"brand_id"`
}

type DeleteProductResponse struct {
	Message string `json:"message"`
}
