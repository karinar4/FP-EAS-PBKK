package cart

import (
	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/pkg/e"
	"gorm.io/gorm"
)

type ICartRepository interface {
	CreateCart(*CartModel) (*CartModel, e.ApiError)
	GetCartByUserID(uuid.UUID) (*CartModel, e.ApiError)
	UpdateCart(*CartModel) (*CartModel, e.ApiError)
	DeleteCart(uuid.UUID) e.ApiError
}

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) ICartRepository {
	return &cartRepository{db}
}

func (r *cartRepository) CreateCart(data *CartModel) (*CartModel, e.ApiError) {
	if err := r.db.Create(data).Error; err != nil {
		return nil, e.NewApiError(e.ErrDatabaseCreateFailed, err.Error())
	}
	return data, nil
}

func (r *cartRepository) GetCartByUserID(userID uuid.UUID) (*CartModel, e.ApiError) {
	cart := &CartModel{}
	if err := r.db.Where("user_id = ?", userID).First(&cart).Error; err != nil {
		return nil, e.NewApiError(e.ErrDatabaseFetchFailed, err.Error())
	}
	return cart, nil
}

func (r *cartRepository) UpdateCart(data *CartModel) (*CartModel, e.ApiError) {
	if err := r.db.Save(data).Error; err != nil {
		return nil, e.NewApiError(e.ErrDatabaseUpdateFailed, err.Error())
	}
	return data, nil
}

func (r *cartRepository) DeleteCart(id uuid.UUID) e.ApiError {
	result := r.db.Delete(&CartModel{}, id)
	if result.Error != nil {
		return e.NewApiError(e.ErrDatabaseDeleteFailed, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return e.NewApiError(e.ErrNotFound, "Cart not found")
	}
	return nil
}
