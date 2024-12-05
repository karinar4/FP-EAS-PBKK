package transaction

import (
	"time"

	"github.com/google/uuid"
)

type CreateTransactionRequest struct {
	TotalQuantity 	int		  `json:"total_quantity" binding:"required"` 
	TotalPrice		float64   `json:"total_price" binding:"required"`
	UserID 			uuid.UUID `json:"user_id" binding:"required"`
}

type CreateTransactionResponse struct {
	ID   			uuid.UUID `json:"id"`
	TransactionDate time.Time `json:"transaction_date"`
	TotalQuantity 	int		  `json:"total_quantity"`
	TotalPrice		float64   `json:"total_price"`
	Status			string	  `json:"status"`
	UserID 			uuid.UUID `json:"user_id"`
}

type GetTransactionResponse struct {
	ID   			uuid.UUID `json:"id"`
	TransactionDate time.Time `json:"transaction_date"`
	TotalQuantity 	int		  `json:"total_quantity"`
	TotalPrice		float64   `json:"total_price"`
	Status			string	  `json:"status"`
	UserID 			uuid.UUID `json:"user_id"`
}

type GetAllTransactionResponse struct {
	Transactions []GetTransactionResponse `json:"transactions"`
}

type UpdateTransactionRequest struct {
	Status			string	  `json:"status" binding:"required"`
}

type UpdateTransactionResponse struct {
	ID   			uuid.UUID `json:"id"`
	TransactionDate time.Time `json:"transaction_date"`
	TotalQuantity 	int		  `json:"total_quantity"`
	TotalPrice		float64   `json:"total_price"`
	Status			string	  `json:"status"`
	UserID 			uuid.UUID `json:"user_id"`
}

type DeleteTransactionResponse struct {
	Message string `json:"message"`
}
