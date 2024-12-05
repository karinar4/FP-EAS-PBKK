package payment

import (
	"time"

	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/common"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/transaction"
)

type Payment struct {
	ID   uuid.UUID `json:"id"`
	PaymentDate time.Time `json:"payment_date"`
	AccountNumber 	string		  `json:"account_number"`
	Status			string	  `json:"status"`
	TransactionID 			uuid.UUID `json:"transaction_id"`
}

type PaymentModel struct {
	common.BaseModels
	PaymentDate time.Time `gorm:"type:datetime;not null"`
	AccountNumber string `gorm:"type:varchar(20);not null"`
	Status string `gorm:"type:enum('pending', 'completed', 'failed', 'cancelled');not null;default:'pending'"`
	TransactionID uuid.UUID `gorm:"type:char(36)"`
	Transaction transaction.TransactionModel `gorm:"foreignKey:TransactionID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (PaymentModel) TableName() string {
	return "payments"
}
