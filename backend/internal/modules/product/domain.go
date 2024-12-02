package product

import (
	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/brand"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/category"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/common"
)

type Product struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	CategoryID  uuid.UUID `json:"category_id"`
	BrandID     uuid.UUID `json:"brand_id"`
}

type ProductModel struct {
	common.BaseModels
	Name        string  `gorm:"type:varchar(255);not null"`
	Description string  `gorm:"type:text"`
	Price       float64 `gorm:"type:decimal(10,2);not null"`
	Stock       int     `gorm:"not null"`
	CategoryID  uuid.UUID
	BrandID     uuid.UUID
	Category    category.CategoryModel `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Brand       brand.BrandModel       `gorm:"foreignKey:BrandID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (ProductModel) TableName() string {
	return "products"
}
