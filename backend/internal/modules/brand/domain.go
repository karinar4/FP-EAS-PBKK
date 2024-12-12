package brand

import (
	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/common"
)

type Brand struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Origin string    `json:"origin"`
}

type BrandModel struct {
	common.BaseModels
	Name string `gorm:"type:varchar(100);not null;unique"`
	Origin string `gorm:"type:varchar(100);not null"`
}

func (BrandModel) TableName() string {
	return "brands"
}
