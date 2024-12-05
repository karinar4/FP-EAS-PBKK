package transaction

import (
	"time"

	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/common"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/auth"
)

type Transaction struct {
	ID   			uuid.UUID `json:"id"`
	TransactionDate time.Time `json:"transaction_date"`
	TotalQuantity 	int		  `json:"total_quantity"`
	TotalPrice		float64   `json:"total_price"`
	Status			string	  `json:"status"`
	UserID 			uuid.UUID `json:"user_id"`
}

type TransactionModel struct {
	common.BaseModels
	TransactionDate time.Time `gorm:"type:datetime;not null"`
	TotalQuantity int `gorm:"not null"`
	TotalPrice float64 `gorm:"type:DECIMAL(10,2);not null"`
	Status string `gorm:"type:enum('pending', 'in_progress','completed','cancelled');not null;default:'pending'"`
	UserID uuid.UUID `gorm:"type:char(36)"`
	User auth.UserModel `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (TransactionModel) TableName() string {
	return "transactions"
}
