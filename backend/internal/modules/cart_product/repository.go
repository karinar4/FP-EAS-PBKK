package cart_product

import (
	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/pkg/e"
	"gorm.io/gorm"
)

type ICartProductRepository interface {
	CreateCartProduct(*CartProductModel) (*CartProductModel, e.ApiError)
	GetAllCartProducts() ([]CartProductModel, e.ApiError)
	DeleteCartProduct(uuid.UUID, uuid.UUID) e.ApiError
}

type cartProductRepository struct {
	db *gorm.DB
}

func NewCartProductRepository(db *gorm.DB) *cartProductRepository {
	return &cartProductRepository{db: db}
}

func (r *cartProductRepository) CreateCartProduct(data *CartProductModel) (*CartProductModel, e.ApiError) {
	if err := r.db.Create(data).Error; err != nil {
		return nil, e.NewApiError(e.ErrDatabaseCreateFailed, err.Error())
	}
	return data, nil
}

func (r *cartProductRepository) GetAllCartProducts() ([]CartProductModel, e.ApiError) {
	var cartProducts []CartProductModel
	if err := r.db.Find(&cartProducts).Error; err != nil {
		return nil, e.NewApiError(e.ErrDatabaseFetchFailed, err.Error())
	}
	return cartProducts, nil
}

func (r *cartProductRepository) DeleteCartProduct(cartID uuid.UUID, productID uuid.UUID) e.ApiError {
	result := r.db.Where("cart_id = ? AND product_id = ?", cartID, productID).Delete(&CartProductModel{})
	if result.Error != nil {
		return e.NewApiError(e.ErrDatabaseDeleteFailed, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return e.NewApiError(e.ErrNotFound, "Cart-product not found")
	}
	return nil
}
