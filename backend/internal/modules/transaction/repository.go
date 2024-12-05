package transaction

import (
	"errors"

	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/pkg/e"
	"gorm.io/gorm"
)

type ITransactionRepository interface {
	CreateTransaction(*TransactionModel) (*TransactionModel, e.ApiError)
	GetAllTransaction() ([]TransactionModel, e.ApiError)
	GetTransactionByID(uuid.UUID) (*TransactionModel, e.ApiError)
	GetAllTransactionByUserID(uuid.UUID) ([]TransactionModel, e.ApiError)
	UpdateTransaction(*TransactionModel) (*TransactionModel, e.ApiError)
	DeleteTransaction(uuid.UUID) e.ApiError
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *transactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) CreateTransaction(data *TransactionModel) (*TransactionModel, e.ApiError) {
	if err := r.db.Create(data).Error; err != nil {
		return nil, e.NewApiError(e.ErrDatabaseCreateFailed, err.Error())
	}
	return data, nil
}

func (r *transactionRepository) GetAllTransaction() ([]TransactionModel, e.ApiError) {
	var transactions []TransactionModel
	if err := r.db.Find(&transactions).Error; err != nil {
		return nil, e.NewApiError(e.ErrDatabaseFetchFailed, err.Error())
	}
	return transactions, nil
}

func (r *transactionRepository) GetTransactionByID(id uuid.UUID) (*TransactionModel, e.ApiError) {
	transaction := &TransactionModel{}
	if err := r.db.Where("id = ?", id).First(transaction).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, e.NewApiError(e.ErrNotFound, "Transaction not found")
		}
		return nil, e.NewApiError(e.ErrDatabaseFetchFailed, err.Error())
	}
	return transaction, nil
}

func (r *transactionRepository) GetAllTransactionByUserID(userID uuid.UUID) ([]TransactionModel, e.ApiError) {
	var transactions []TransactionModel
	if err := r.db.Where("user_id = ?", userID).Find(&transactions).Error; err != nil {
		return nil, e.NewApiError(e.ErrDatabaseFetchFailed, err.Error())
	}
	return transactions, nil
}

func (r *transactionRepository) UpdateTransaction(data *TransactionModel) (*TransactionModel, e.ApiError) {
	if err := r.db.Save(data).Error; err != nil {
		return nil, e.NewApiError(e.ErrDatabaseUpdateFailed, err.Error())
	}
	return data, nil
}

func (r *transactionRepository) DeleteTransaction(id uuid.UUID) e.ApiError {
	result := r.db.Delete(&TransactionModel{}, id)
	if result.Error != nil {
		return e.NewApiError(e.ErrDatabaseDeleteFailed, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return e.NewApiError(e.ErrNotFound, "Transaction not found")
	}
	return nil
}
