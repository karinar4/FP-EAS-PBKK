package product_transaction

import (
	"time"

	"github.com/google/uuid"
	// "github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/product"
)

type ProductTransaction struct {
	ProductID		uuid.UUID	`json:"product_id"`
	TransactionID   uuid.UUID	`json:"transaction_id"`
	RentStartDate 	time.Time	`json:"rent_start_date"`
	RentEndDate     time.Time   `json:"rent_end_date"`
	Quantity       	int			`json:"quantity"`
	Price  			float64		`json:"price"`
}

type ProductTransactionModel struct {
	CreatedAt time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	ProductID uuid.UUID `gorm:"type:char(36)"`
	TransactionID uuid.UUID `gorm:"type:char(36)"`
	RentStartDate time.Time `gorm:"type:datetime;not null"`
	RentEndDate time.Time	`gorm:"type:datetime;not null"`
	Quantity int	`gorm:"not null"`
	Price float64	`gorm:"type:decimal(10,2);not null"`

	// Product   product.ProductModel   `gorm:"foreignKey:ProductID"`
	// Transaction transaction.TransactionModel `gorm:"foreignKey:TransactionID"`
}

func (ProductTransactionModel) TableName() string {
	return "product_transactions"
}

func (ProductTransactionModel) PrimaryKey() []string {
	return []string{"product_id", "transaction_id"}
}