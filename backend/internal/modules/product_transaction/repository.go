package product_transaction

import (
	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/pkg/e"
	"gorm.io/gorm"
)

type IProductTransactionRepository interface {
	CreateProductTransaction(*ProductTransactionModel) (*ProductTransactionModel, e.ApiError)
	GetAllProductTransactions() ([]ProductTransactionModel, e.ApiError)
	// GetProductByID(uuid.UUID) (*ProductTransactionModel, e.ApiError)
	// UpdateProduct(id uuid.UUID, updatedFields map[string]interface{}) (*ProductTransactionModel, e.ApiError)
	DeleteProductTransaction(uuid.UUID, uuid.UUID) e.ApiError
}

type productTransactionRepository struct {
	db *gorm.DB
}

func NewProductTransactionRepository(db *gorm.DB) *productTransactionRepository {
	return &productTransactionRepository{db: db}
}

func (r *productTransactionRepository) CreateProductTransaction(data *ProductTransactionModel) (*ProductTransactionModel, e.ApiError) {
	if err := r.db.Create(data).Error; err != nil {
		return nil, e.NewApiError(e.ErrDatabaseCreateFailed, err.Error())
	}
	return data, nil
}

func (r *productTransactionRepository) GetAllProductTransactions() ([]ProductTransactionModel, e.ApiError) {
	var product_transactions []ProductTransactionModel
	if err := r.db.Find(&product_transactions).Error; err != nil {
		return nil, e.NewApiError(e.ErrDatabaseFetchFailed, err.Error())
	}
	return product_transactions, nil
}

// func (r *productTransactionRepository) GetProductByID(id uuid.UUID) (*ProductTransactionModel, e.ApiError) {
// 	var product ProductTransactionModel
// 	if err := r.db.Preload("Category").Preload("Brand").First(&product, "id = ?", id).Error; err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return nil, e.NewApiError(e.ErrNotFound, "Product not found")
// 		}
// 		return nil, e.NewApiError(e.ErrDatabaseFetchFailed, err.Error())
// 	}
// 	return &product, nil
// }

// func (r *productTransactionRepository) UpdateProduct(id uuid.UUID, updatedFields map[string]interface{}) (*ProductTransactionModel, e.ApiError) {
// 	if err := r.db.Model(&ProductTransactionModel{}).Where("id = ?", id).Updates(updatedFields).Error; err != nil {
// 		return nil, e.NewApiError(e.ErrDatabaseUpdateFailed, err.Error())
// 	}
// 	var updatedProduct ProductTransactionModel
// 	if err := r.db.Preload("Category").Preload("Brand").First(&updatedProduct, "id = ?", id).Error; err != nil {
// 		return nil, e.NewApiError(e.ErrDatabaseFetchFailed, err.Error())
// 	}
// 	return &updatedProduct, nil
// }

func (r *productTransactionRepository) DeleteProductTransaction(product_id uuid.UUID, transaction_id uuid.UUID) e.ApiError {
	result := r.db.Where("transaction_id = ? AND product_id = ?", transaction_id, product_id).Delete(&ProductTransactionModel{})
	if result.Error != nil {
		return e.NewApiError(e.ErrDatabaseDeleteFailed, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return e.NewApiError(e.ErrNotFound, "Product-transaction not found")
	}
	return nil
}
