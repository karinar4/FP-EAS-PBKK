package common

import (
	"time"

	"github.com/google/uuid"
)

type BaseModels struct {
	ID        uuid.UUID `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	// DeletedAt gorm.DeletedAt
}
