package transaction

import (
	"time"

	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/auth"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/product_transaction"
)

type CreateTransactionRequest struct {
	TotalQuantity 	int		  `json:"total_quantity" binding:"required"` 
	TotalPrice		float64   `json:"total_price" binding:"required"`
	UserID 			uuid.UUID `json:"user_id" binding:"required"`
	ProductTransactions 	[]product_transaction.CreateProductTransactionRequest `json:"product_transactions" binding:"required"`
}

type CreateTransactionResponse struct {
	ID   			uuid.UUID `json:"id"`
	TransactionDate time.Time `json:"transaction_date"`
	TotalQuantity 	int		  `json:"total_quantity"`
	TotalPrice		float64   `json:"total_price"`
	Status			string	  `json:"status"`
}

type GetTransactionResponse struct {
	ID   			uuid.UUID 		`json:"id"`
	TransactionDate time.Time 		`json:"transaction_date"`
	TotalQuantity 	int		  		`json:"total_quantity"`
	TotalPrice		float64   		`json:"total_price"`
	Status			string	  		`json:"status"`
	User 			auth.GetUser 	`json:"user"`
}

type GetTransactionByUserIDResponse struct {
	ID   			uuid.UUID 		`json:"id"`
	TransactionDate time.Time 		`json:"transaction_date"`
	TotalQuantity 	int		  		`json:"total_quantity"`
	TotalPrice		float64   		`json:"total_price"`
	Status			string	  		`json:"status"`
}

type GetAllTransactionResponse struct {
	Transactions []GetTransactionResponse `json:"transactions"`
}

type GetAllTransactionByUserIDResponse struct {
	Transactions []GetTransactionByUserIDResponse `json:"transactions"`
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
}

type DeleteTransactionResponse struct {
	Message string `json:"message"`
}
