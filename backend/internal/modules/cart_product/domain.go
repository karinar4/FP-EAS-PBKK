package cart_product

import (
	"time"

	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/cart"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/product"
)

type CartProduct struct {
	CartID        uuid.UUID `json:"cart_id"`
	ProductID     uuid.UUID `json:"product_id"`
	RentStartDate time.Time `json:"rent_start_date"`
	RentEndDate   time.Time `json:"rent_end_date"`
	Quantity      int       `json:"quantity"`
	Price         float64   `json:"price"`
}

type CartProductModel struct {
	CreatedAt     time.Time            `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time            `gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	CartID        uuid.UUID            `gorm:"type:char(36)"`
	ProductID     uuid.UUID            `gorm:"type:char(36)"`
	RentStartDate time.Time            `gorm:"type:datetime;not null"`
	RentEndDate   time.Time            `gorm:"type:datetime;not null"`
	Quantity      int                  `gorm:"not null"`
	Price         float64              `gorm:"type:decimal(10,2);not null"`
	Cart          cart.CartModel       `gorm:"foreignKey:CartID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Product       product.ProductModel `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (CartProductModel) TableName() string {
	return "carts_products"
}

func (CartProductModel) PrimaryKey() []string {
	return []string{"cart_id", "product_id"}
}
