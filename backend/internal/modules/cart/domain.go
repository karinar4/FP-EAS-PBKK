package cart

import (
	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/auth"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/common"
)

type Cart struct {
	ID            uuid.UUID `json:"id"`
	UserID        uuid.UUID `json:"user_id"`
	ProductID     uuid.UUID `json:"product_id"`
	TotalQuantity int       `json:"total_quantity"`
	TotalPrice    float64   `json:"total_price"`
}

type CartModel struct {
	common.BaseModels
	UserID uuid.UUID
	User   auth.UserModel `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	// ProductID     uuid.UUID
	TotalQuantity int     `gorm:"not null;default:0"`
	TotalPrice    float64 `gorm:"type:decimal(10,2);not null;default:0"`
	// Product       product.ProductModel `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (CartModel) TableName() string {
	return "carts"
}
