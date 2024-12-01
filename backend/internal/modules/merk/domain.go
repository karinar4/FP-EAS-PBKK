package merk

import (
	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/common"
)

type Merk struct {
	ID   uuid.UUID `json:"id"`
	Nama string    `json:"nama"`
}

type MerkModel struct {
	common.BaseModels
	Nama string `gorm:"type:varchar(100);not null;unique"`
}

func (MerkModel) TableName() string {
	return "merk"
}
