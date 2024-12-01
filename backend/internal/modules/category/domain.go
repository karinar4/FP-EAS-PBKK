package category

import (
	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/common"
)

type Category struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type CategoryModel struct {
	common.BaseModels
	Name string `gorm:"type:varchar(100);not null;unique"`
}

func (CategoryModel) TableName() string {
	return "categories"
}
