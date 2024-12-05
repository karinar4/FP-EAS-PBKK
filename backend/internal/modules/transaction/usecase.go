package transaction

import (
	"time"

	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/common"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/pkg/e"
)

type ITransactionUseCase interface {
	CreateTransaction(*CreateTransactionRequest) (*CreateTransactionResponse, e.ApiError)
	GetAllTransaction() (*GetAllTransactionResponse, e.ApiError)
	GetTransactionByID(uuid.UUID) (*GetTransactionResponse, e.ApiError)
	GetAllTransactionByUserID(uuid.UUID) (*GetAllTransactionResponse, e.ApiError)
	UpdateTransaction(uuid.UUID, *UpdateTransactionRequest) (*UpdateTransactionResponse, e.ApiError)
	DeleteTransaction(uuid.UUID) e.ApiError
}

type transactionUseCase struct {
	repo ITransactionRepository
}

func NewTransactionUseCase(repo ITransactionRepository) ITransactionUseCase {
	return &transactionUseCase{repo: repo}
}

func (uc *transactionUseCase) CreateTransaction(req *CreateTransactionRequest) (*CreateTransactionResponse, e.ApiError) {
	transaction := &TransactionModel{
		BaseModels: common.BaseModels{
			ID: uuid.New(),
		},
		TransactionDate: time.Now(),
		TotalQuantity: req.TotalQuantity,
		TotalPrice: req.TotalPrice,
		Status: "pending",
		UserID: req.UserID,
	}

	result, err := uc.repo.CreateTransaction(transaction)
	if err != nil {
		return nil, e.NewApiError(e.ErrDatabaseCreateFailed, err.Error())
	}

	return &CreateTransactionResponse{
		ID:              result.ID,
		TransactionDate: result.TransactionDate,
		TotalQuantity:   result.TotalQuantity,
		TotalPrice:      result.TotalPrice,
		Status:          result.Status,
		UserID:          result.UserID,
	}, nil
}

func (uc *transactionUseCase) GetAllTransaction() (*GetAllTransactionResponse, e.ApiError) {
	transactions, err := uc.repo.GetAllTransaction()
	if err != nil {
		return nil, e.NewApiError(e.ErrDatabaseFetchFailed, err.Error())
	}

	var response []GetTransactionResponse
	for _, transaction := range transactions {
		response = append(response, GetTransactionResponse{
			ID:              transaction.ID,
			TransactionDate: transaction.TransactionDate,
			TotalQuantity:   transaction.TotalQuantity,
			TotalPrice:      transaction.TotalPrice,
			Status:          transaction.Status,
			UserID:          transaction.UserID,
		})
	}

	return &GetAllTransactionResponse{Transactions: response}, nil
}

func (uc *transactionUseCase) GetTransactionByID(id uuid.UUID) (*GetTransactionResponse, e.ApiError) {
	transaction, err := uc.repo.GetTransactionByID(id)
	if err != nil {
		return nil, err
	}

	return &GetTransactionResponse{
		ID: transaction.ID,
		TransactionDate: transaction.TransactionDate,
		TotalQuantity: transaction.TotalQuantity,
		TotalPrice: transaction.TotalPrice,
		Status: transaction.Status,
		UserID: transaction.UserID,
	}, nil
}

func (uc *transactionUseCase) GetAllTransactionByUserID(userID uuid.UUID) (*GetAllTransactionResponse, e.ApiError) {
	transactions, err := uc.repo.GetAllTransactionByUserID(userID)
	if err != nil {
		return nil, e.NewApiError(e.ErrDatabaseFetchFailed, err.Error())
	}

	var response []GetTransactionResponse
	for _, transaction := range transactions {
		response = append(response, GetTransactionResponse{
			ID:              transaction.ID,
			TransactionDate: transaction.TransactionDate,
			TotalQuantity:   transaction.TotalQuantity,
			TotalPrice:      transaction.TotalPrice,
			Status:          transaction.Status,
			UserID:          transaction.UserID,
		})
	}

	return &GetAllTransactionResponse{Transactions: response}, nil
}

func (uc *transactionUseCase) UpdateTransaction(id uuid.UUID, req *UpdateTransactionRequest) (*UpdateTransactionResponse, e.ApiError) {
	transaction, err := uc.repo.GetTransactionByID(id)
	if err != nil {
		return nil, err
	}

	transaction.Status = req.Status

	updatedTransaction, updateErr := uc.repo.UpdateTransaction(transaction)
	if updateErr != nil {
		return nil, updateErr
	}

	return &UpdateTransactionResponse{
		ID: updatedTransaction.ID,
		TransactionDate: updatedTransaction.TransactionDate,
		TotalQuantity: updatedTransaction.TotalQuantity,
		TotalPrice: updatedTransaction.TotalPrice,
		Status: updatedTransaction.Status,
		UserID: updatedTransaction.UserID,
	}, nil
}

func (uc *transactionUseCase) DeleteTransaction(id uuid.UUID) e.ApiError {
	return uc.repo.DeleteTransaction(id)
}
