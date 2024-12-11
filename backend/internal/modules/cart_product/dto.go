package cart_product

import (
	"time"

	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/product"
)

type CreateCartProductRequest struct {
	CartID        uuid.UUID `json:"cart_id" binding:"required"`
	ProductID     uuid.UUID `json:"product_id" binding:"required"`
	RentStartDate time.Time `json:"rent_start_date" binding:"required"`
	RentEndDate   time.Time `json:"rent_end_date" binding:"required"`
	Quantity      int       `json:"quantity" binding:"required"`
	Price         float64   `json:"price" binding:"required"`
}

type CreateCartProductResponse struct {
	CartID        uuid.UUID `json:"cart_id"`
	ProductID     uuid.UUID `json:"product_id"`
	RentStartDate time.Time `json:"rent_start_date"`
	RentEndDate   time.Time `json:"rent_end_date"`
	Quantity      int       `json:"quantity"`
	Price         float64   `json:"price"`
}

type GetCartProductResponse struct {
	CartID        uuid.UUID `json:"cart_id"`
	// ProductID     uuid.UUID `json:"product_id"`
	Product product.GetProductResponse `json:"product"`
	RentStartDate time.Time `json:"rent_start_date"`
	RentEndDate   time.Time `json:"rent_end_date"`
	Quantity      int       `json:"quantity"`
	Price         float64   `json:"price"`
}

type GetAllCartProductsResponse struct {
	CartProducts []GetCartProductResponse `json:"cart_products"`
}

type UpdateCartProductRequest struct {
	RentStartDate time.Time `json:"rent_start_date"`
	RentEndDate   time.Time `json:"rent_end_date"`
	Quantity      int       `json:"quantity"`
	Price         float64   `json:"price"`
}

type UpdateCartProductResponse struct {
	CartID        uuid.UUID `json:"cart_id"`
	ProductID     uuid.UUID `json:"product_id"`
	RentStartDate time.Time `json:"rent_start_date"`
	RentEndDate   time.Time `json:"rent_end_date"`
	Quantity      int       `json:"quantity"`
	Price         float64   `json:"price"`
}

type DeleteCartProductResponse struct {
	Message string `json:"message"`
}
