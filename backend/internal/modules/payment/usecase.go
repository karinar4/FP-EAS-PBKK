package payment

import (
	"time"

	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/common"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/pkg/e"
)

type IPaymentUseCase interface {
	CreatePayment(*CreatePaymentRequest) (*CreatePaymentResponse, e.ApiError)
	GetAllPayment() (*GetAllPaymentResponse, e.ApiError)
	GetPaymentByID(uuid.UUID) (*GetPaymentResponse, e.ApiError)
	UpdatePayment(uuid.UUID, *UpdatePaymentRequest) (*UpdatePaymentResponse, e.ApiError)
	DeletePayment(uuid.UUID) e.ApiError
}

type paymentUseCase struct {
	repo IPaymentRepository
}

func NewPaymentUseCase(repo IPaymentRepository) IPaymentUseCase {
	return &paymentUseCase{repo: repo}
}

func (uc *paymentUseCase) CreatePayment(req *CreatePaymentRequest) (*CreatePaymentResponse, e.ApiError) {
	payment := &PaymentModel{
		BaseModels: common.BaseModels{
			ID: uuid.New(),
		},
		PaymentDate: time.Now(),
		AccountNumber: req.AccountNumber,
		Status: "pending",
		TransactionID: req.TransactionID,
	}

	result, err := uc.repo.CreatePayment(payment)
	if err != nil {
		return nil, e.NewApiError(e.ErrDatabaseCreateFailed, err.Error())
	}

	return &CreatePaymentResponse{
		ID:   result.ID,
		PaymentDate: result.PaymentDate,
		AccountNumber: result.AccountNumber,
		Status: result.Status,
		TransactionID: result.TransactionID,
	}, nil
}

func (uc *paymentUseCase) GetAllPayment() (*GetAllPaymentResponse, e.ApiError) {
	payments, err := uc.repo.GetAllPayment()
	if err != nil {
		return nil, e.NewApiError(e.ErrDatabaseFetchFailed, err.Error())
	}

	var response []GetPaymentResponse
	for _, payment := range payments {
		response = append(response, GetPaymentResponse{
			ID:   payment.ID,
			PaymentDate: payment.PaymentDate,
			AccountNumber: payment.AccountNumber,
			Status: payment.Status,
			TransactionID: payment.TransactionID,
		})
	}

	return &GetAllPaymentResponse{Payments: response}, nil
}

func (uc *paymentUseCase) GetPaymentByID(id uuid.UUID) (*GetPaymentResponse, e.ApiError) {
	payment, err := uc.repo.GetPaymentByID(id)
	if err != nil {
		return nil, err
	}

	return &GetPaymentResponse{
		ID:   payment.ID,
		PaymentDate: payment.PaymentDate,
		AccountNumber: payment.AccountNumber,
		Status: payment.Status,
		TransactionID: payment.TransactionID,
	}, nil
}

func (uc *paymentUseCase) UpdatePayment(id uuid.UUID, req *UpdatePaymentRequest) (*UpdatePaymentResponse, e.ApiError) {
	payment, err := uc.repo.GetPaymentByID(id)
	if err != nil {
		return nil, err
	}

	payment.Status = req.Status

	updatedPayment, updateErr := uc.repo.UpdatePayment(payment)
	if updateErr != nil {
		return nil, updateErr
	}

	return &UpdatePaymentResponse{
		ID:   updatedPayment.ID,
		PaymentDate: updatedPayment.PaymentDate,
		AccountNumber: updatedPayment.AccountNumber,
		Status: updatedPayment.Status,
		TransactionID: updatedPayment.TransactionID,
	}, nil
}

func (uc *paymentUseCase) DeletePayment(id uuid.UUID) e.ApiError {
	return uc.repo.DeletePayment(id)
}
