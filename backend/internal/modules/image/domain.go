package image

import (
	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/common"
)

type Image struct {
	ID        uuid.UUID `json:"id"`
	ProductID uuid.UUID `json:"product_id"`
	URL       string    `json:"url"`
}

type ImageModel struct {
	common.BaseModels
	ProductID uuid.UUID `gorm:"type:char(36);not null"`
	URL       string    `gorm:"type:varchar(255);not null"`
}

func (ImageModel) TableName() string {
	return "product_images"
}
