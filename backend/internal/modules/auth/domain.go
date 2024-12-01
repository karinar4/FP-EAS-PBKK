package auth

import (
	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/common"
)

type (
	RegisterUserDomain struct {
		Id       uuid.UUID
		Name     string
		Email    string
		Password string
	}

	UserModel struct {
		common.BaseModels
		Name    string  `gorm:"type:varchar(255);not null"`
		Email    string  `gorm:"type:varchar(255);not null;unique"`
		Password string  `gorm:"type:varchar(255);not null"`
		Telephone  *string `gorm:"type:varchar(15)"`
		Address   *string `gorm:"type:varchar(255)"`
		Role     string  `gorm:"type:enum('user', 'admin');not null;default:'user'"`
	}

	PayloadToken struct {
		ID   uuid.UUID
		Role string
	}
)

func (UserModel) TableName() string {
	return "users"
}
