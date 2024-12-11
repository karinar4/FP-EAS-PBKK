package payment

import (
	"time"

	"github.com/google/uuid"
)

type CreatePaymentRequest struct {
	AccountNumber 	string		  `json:"account_number" binding:"required"`
	TransactionID 			uuid.UUID `json:"transaction_id" binding:"required"`
}

type CreatePaymentResponse struct {
	ID   uuid.UUID `json:"id"`
	PaymentDate time.Time `json:"payment_date"`
	AccountNumber 	string		  `json:"account_number"`
	Status			string	  `json:"status"`
	TransactionID 			uuid.UUID `json:"transaction_id"`
}

type GetPaymentResponse struct {
	ID   uuid.UUID `json:"id"`
	PaymentDate time.Time `json:"payment_date"`
	AccountNumber 	string		  `json:"account_number"`
	Status			string	  `json:"status"`
	TransactionID 			uuid.UUID `json:"transaction_id"`
}

type GetAllPaymentResponse struct {
	Payments []GetPaymentResponse `json:"payments"`
}

type UpdatePaymentRequest struct {
	Status			string	  `json:"status"`
}

type UpdatePaymentResponse struct {
	ID   uuid.UUID `json:"id"`
	PaymentDate time.Time `json:"payment_date"`
	AccountNumber 	string		  `json:"account_number"`
	Status			string	  `json:"status"`
	TransactionID 			uuid.UUID `json:"transaction_id"`
}

type DeletePaymentResponse struct {
	Message string `json:"message"`
}
