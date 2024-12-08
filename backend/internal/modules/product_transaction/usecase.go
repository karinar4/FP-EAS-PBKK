package product_transaction

import (
	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/pkg/e"
)

type IProductTransactionUseCase interface {
	CreateProductTransaction(*CreateProductTransactionRequest) (*CreateProductTransactionResponse, e.ApiError)
	GetAllProductTransactions() (*GetAllProductTransactionsResponse, e.ApiError)
	// GetByTransactionID(uuid.UUID) (*GetProductTransactionResponse, e.ApiError)
	// UpdateProductTransaction(uuid.UUID, *UpdateProductTransactionRequest) (*UpdateProductTransactionResponse, e.ApiError)
	DeleteProductTransaction(uuid.UUID, uuid.UUID) e.ApiError
}

type productTransactionUseCase struct {
	repo IProductTransactionRepository
}

func NewProductTransactionUseCase(repo IProductTransactionRepository) IProductTransactionUseCase {
	return &productTransactionUseCase{repo: repo}
}

func (uc *productTransactionUseCase) CreateProductTransaction(req *CreateProductTransactionRequest) (*CreateProductTransactionResponse, e.ApiError) {
	product_transaction := &ProductTransactionModel{
		ProductID:        req.ProductID,
		TransactionID: req.TransactionID,
		RentStartDate:       req.RentStartDate,
		RentEndDate:       req.RentEndDate,
		Quantity:  req.Quantity,
		Price:     req.Price,
	}

	result, err := uc.repo.CreateProductTransaction(product_transaction)
	if err != nil {
		return nil, e.NewApiError(e.ErrDatabaseCreateFailed, err.Error())
	}

	return &CreateProductTransactionResponse{
		ProductID:        result.ProductID,
		TransactionID: result.TransactionID,
		RentStartDate:       result.RentStartDate,
		RentEndDate:       result.RentEndDate,
		Quantity:  result.Quantity,
		Price:     result.Price,
	}, nil
}

func (uc *productTransactionUseCase) GetAllProductTransactions() (*GetAllProductTransactionsResponse, e.ApiError) {
	product_transactions, err := uc.repo.GetAllProductTransactions()
	if err != nil {
		return nil, e.NewApiError(e.ErrDatabaseFetchFailed, err.Error())
	}

	var response []GetProductTransactionResponse
	for _, product_transaction := range product_transactions {
		response = append(response, GetProductTransactionResponse{
			ProductID:        product_transaction.ProductID,
			TransactionID: product_transaction.TransactionID,
			RentStartDate:       product_transaction.RentStartDate,
			RentEndDate:       product_transaction.RentEndDate,
			Quantity:  product_transaction.Quantity,
			Price:     product_transaction.Price,
		})
	}

	return &GetAllProductTransactionsResponse{ProductTransactions: response}, nil
}

// func (uc *productTransactionUseCase) GetByTransactionID(id uuid.UUID) (*GetProductTransactionResponse, e.ApiError) {
// 	product_transactions, err := uc.repo.GetByTransactionID(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &GetProductTransactionResponse{
// 		ProductID:        product_transaction.ProductID,
// 		TransactionID: product_transaction.TransactionID,
// 			RentStartDate:       product_transaction.RentStartDate,
// 			RentEndDate:       product_transaction.RentEndDate,
// 			Quantity:  product_transaction.Quantity,
// 			Price:     product_transaction.Price,
// 	}, nil
// }

func (uc *productTransactionUseCase) DeleteProductTransaction(product_id uuid.UUID, transaction_id uuid.UUID) e.ApiError {
	return uc.repo.DeleteProductTransaction(product_id, transaction_id)
}