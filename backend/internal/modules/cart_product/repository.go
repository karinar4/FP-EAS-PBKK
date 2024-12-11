package cart_product

import (
	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/pkg/e"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/cart"
	"gorm.io/gorm"
)

type ICartProductRepository interface {
	CreateCartProduct(*CartProductModel) (*CartProductModel, e.ApiError)
	GetAllCartProducts() ([]CartProductModel, e.ApiError)
	GetCartProductsByID(uuid.UUID) ([]CartProductModel, e.ApiError)
	UpdateCartProduct(uuid.UUID, uuid.UUID, map[string]interface{}) (*CartProductModel, e.ApiError)
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
	if err := r.db.Preload("Product").Find(&cartProducts).Error; err != nil {
		return nil, e.NewApiError(e.ErrDatabaseFetchFailed, err.Error())
	}
	return cartProducts, nil
}

func (r *cartProductRepository) GetCartProductsByID(cartID uuid.UUID) ([]CartProductModel, e.ApiError) {
	var cartProducts []CartProductModel
	if err := r.db.Preload("Product").Where("cart_id = ?", cartID).Find(&cartProducts).Error; err != nil {
		return nil, e.NewApiError(e.ErrDatabaseFetchFailed, err.Error())
	}
	return cartProducts, nil
}

func (r *cartProductRepository) UpdateCartProduct(cart_id uuid.UUID, product_id uuid.UUID, updatedFields map[string]interface{}) (*CartProductModel, e.ApiError) {
	if err := r.db.Model(&CartProductModel{}).Where("cart_id = ? AND product_id = ?", cart_id, product_id).Updates(updatedFields).Error; err != nil {
		return nil, e.NewApiError(e.ErrDatabaseUpdateFailed, err.Error())
	}

	var updatedCartProduct CartProductModel
	if err := r.db.First(&updatedCartProduct, "cart_id = ? AND product_id = ?", cart_id, product_id).Error; err != nil {
		return nil, e.NewApiError(e.ErrDatabaseFetchFailed, err.Error())
	}

	var cartProducts []CartProductModel
    if err := r.db.Where("cart_id = ?", cart_id).Find(&cartProducts).Error; err != nil {
        return nil, e.NewApiError(e.ErrDatabaseFetchFailed, err.Error())
    }

	totalQuantity := 0
    totalPrice := float64(0)
    for _, cartProduct := range cartProducts {
        totalQuantity += cartProduct.Quantity
        totalPrice += cartProduct.Price
    }

	if err := r.db.Model(&cart.CartModel{}).Where("id = ?", cart_id).Updates(map[string]interface{}{
        "total_quantity": totalQuantity,
        "total_price":    totalPrice,
    }).Error; err != nil {
        return nil, e.NewApiError(e.ErrDatabaseUpdateFailed, err.Error())
    }

	return &updatedCartProduct, nil
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
