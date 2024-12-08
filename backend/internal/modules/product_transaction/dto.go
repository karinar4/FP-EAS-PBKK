package product_transaction

import (
	"time"

	"github.com/google/uuid"
)

type CreateProductTransactionRequest struct {
	ProductID		uuid.UUID	`json:"product_id" binding:"required"`
	TransactionID   uuid.UUID	`json:"transaction_id"`
	RentStartDate 	time.Time	`json:"rent_start_date" binding:"required"`
	RentEndDate     time.Time   `json:"rent_end_date" binding:"required"`
	Quantity       	int			`json:"quantity" binding:"required"`
	Price  			float64		`json:"price" binding:"required"`
}

type CreateProductTransactionResponse struct {
	ProductID		uuid.UUID	`json:"product_id"`
	TransactionID   uuid.UUID	`json:"transaction_id"`
	RentStartDate 	time.Time	`json:"rent_start_date"`
	RentEndDate     time.Time   `json:"rent_end_date"`
	Quantity       	int			`json:"quantity"`
	Price  			float64		`json:"price"`
}

type GetProductTransactionResponse struct {
	ProductID		uuid.UUID	`json:"product_id"`
	TransactionID   uuid.UUID	`json:"transaction_id"`
	RentStartDate 	time.Time	`json:"rent_start_date"`
	RentEndDate     time.Time   `json:"rent_end_date"`
	Quantity       	int			`json:"quantity"`
	Price  			float64		`json:"price"`
}

type GetAllProductTransactionsResponse struct {
	ProductTransactions []GetProductTransactionResponse `json:"product_transactions"`
}

// type UpdateProductRequest struct {
// 	Name        string    `json:"name"`
// 	Description string    `json:"description"`
// 	Price       float64   `json:"price"`
// 	Stock       int       `json:"stock"`
// 	CategoryID  uuid.UUID `json:"category_id"`
// 	BrandID     uuid.UUID `json:"brand_id"`
// }

// type UpdateProductResponse struct {
// 	ID          uuid.UUID `json:"id"`
// 	Name        string    `json:"name"`
// 	Description string    `json:"description"`
// 	Price       float64   `json:"price"`
// 	Stock       int       `json:"stock"`
// 	CategoryID  uuid.UUID `json:"category_id"`
// 	BrandID     uuid.UUID `json:"brand_id"`
// }

type DeleteProductTransactionResponse struct {
	Message string `json:"message"`
}
