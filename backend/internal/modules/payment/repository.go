package payment

import (
	"errors"

	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/pkg/e"
	"gorm.io/gorm"
)

type IPaymentRepository interface {
	CreatePayment(*PaymentModel) (*PaymentModel, e.ApiError)
	GetAllPayment() ([]PaymentModel, e.ApiError)
	GetPaymentByID(uuid.UUID) (*PaymentModel, e.ApiError)
	UpdatePayment(*PaymentModel) (*PaymentModel, e.ApiError)
	DeletePayment(uuid.UUID) e.ApiError
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *paymentRepository {
	return &paymentRepository{db}
}

func (r *paymentRepository) CreatePayment(data *PaymentModel) (*PaymentModel, e.ApiError) {
	if err := r.db.Create(data).Error; err != nil {
		return nil, e.NewApiError(e.ErrDatabaseCreateFailed, err.Error())
	}
	return data, nil
}

func (r *paymentRepository) GetAllPayment() ([]PaymentModel, e.ApiError) {
	var payments []PaymentModel
	if err := r.db.Find(&payments).Error; err != nil {
		return nil, e.NewApiError(e.ErrDatabaseFetchFailed, err.Error())
	}
	return payments, nil
}

func (r *paymentRepository) GetPaymentByID(id uuid.UUID) (*PaymentModel, e.ApiError) {
	payment := &PaymentModel{}
	if err := r.db.Where("id = ?", id).First(payment).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, e.NewApiError(e.ErrNotFound, "Payment not found")
		}
		return nil, e.NewApiError(e.ErrDatabaseFetchFailed, err.Error())
	}
	return payment, nil
}

func (r *paymentRepository) UpdatePayment(data *PaymentModel) (*PaymentModel, e.ApiError) {
	if err := r.db.Save(data).Error; err != nil {
		return nil, e.NewApiError(e.ErrDatabaseUpdateFailed, err.Error())
	}
	return data, nil
}

func (r *paymentRepository) DeletePayment(id uuid.UUID) e.ApiError {
	result := r.db.Delete(&PaymentModel{}, id)
	if result.Error != nil {
		return e.NewApiError(e.ErrDatabaseDeleteFailed, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return e.NewApiError(e.ErrNotFound, "Payment not found")
	}
	return nil
}
