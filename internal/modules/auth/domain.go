package auth

import (
	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/internal/modules/common"
)

type (
	RegisterUserDomain struct {
		Id       uuid.UUID
		Nama     string
		Email    string
		Password string
	}

	UserModel struct {
		common.BaseModels
		Nama     string  `gorm:"type:varchar(255);not null"`
		Email    string  `gorm:"type:varchar(255);not null;unique"`
		Password string  `gorm:"type:varchar(255);not null"`
		Telepon  *string `gorm:"type:varchar(15)"`
		Alamat   *string `gorm:"type:varchar(255)"`
		Role     string  `gorm:"type:varchar(10);not null;default:'user'"`
	}

	PayloadToken struct {
		ID   uuid.UUID
		Role string
	}
)

func (UserModel) TableName() string {
	return "users"
}
