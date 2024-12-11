package product_transaction

import (
	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/pkg/e"
	"gorm.io/gorm"
)

type IProductTransactionRepository interface {
	CreateProductTransaction(*ProductTransactionModel) (*ProductTransactionModel, e.ApiError)
	GetAllProductTransactions() ([]ProductTransactionModel, e.ApiError)
	GetProductTransactionsByID(uuid.UUID) ([]ProductTransactionModel, e.ApiError)
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
	if err := r.db.Preload("Product").Find(&product_transactions).Error; err != nil {
		return nil, e.NewApiError(e.ErrDatabaseFetchFailed, err.Error())
	}
	return product_transactions, nil
}

func (r *productTransactionRepository) GetProductTransactionsByID(transactionId uuid.UUID) ([]ProductTransactionModel, e.ApiError) {
	var product_transactions []ProductTransactionModel
	if err := r.db.Preload("Product").Where("transaction_id = ?", transactionId).Find(&product_transactions).Error; err != nil {
		return nil, e.NewApiError(e.ErrDatabaseFetchFailed, err.Error())
	}
	return product_transactions, nil
}

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
