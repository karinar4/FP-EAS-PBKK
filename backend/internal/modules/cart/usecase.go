package cart

import (
	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/common"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/pkg/e"
)

type ICartUseCase interface {
	CreateCart(*CreateCartRequest) (*CartResponse, e.ApiError)
	GetCartByUser(uuid.UUID) (*CartResponse, e.ApiError)
	UpdateCart(uuid.UUID, *UpdateCartRequest) (*CartResponse, e.ApiError)
	DeleteCart(uuid.UUID) e.ApiError
}

type cartUseCase struct {
	repo ICartRepository
}

func NewCartUseCase(repo ICartRepository) ICartUseCase {
	return &cartUseCase{repo}
}

func (uc *cartUseCase) CreateCart(req *CreateCartRequest) (*CartResponse, e.ApiError) {
	cart := &CartModel{
		BaseModels: common.BaseModels{ID: uuid.New()},
		UserID:     req.UserID,
		// ProductID:     req.ProductID,
		TotalQuantity: req.TotalQuantity,
		TotalPrice:    req.TotalPrice,
	}

	createdCart, err := uc.repo.CreateCart(cart)
	if err != nil {
		return nil, err
	}

	return toCartResponse(createdCart), nil
}

func (uc *cartUseCase) GetCartByUser(userID uuid.UUID) (*CartResponse, e.ApiError) {
	cart, err := uc.repo.GetCartByUserID(userID)
	if err != nil {
		return nil, err
	}

	return toCartResponse(cart), nil
}

func (uc *cartUseCase) UpdateCart(id uuid.UUID, req *UpdateCartRequest) (*CartResponse, e.ApiError) {
	cart, err := uc.repo.GetCartByUserID(req.UserID)
	if err != nil {
		return nil, err
	}

	cart.TotalQuantity = req.TotalQuantity
	cart.TotalPrice = req.TotalPrice

	updatedCart, updateErr := uc.repo.UpdateCart(cart)
	if updateErr != nil {
		return nil, updateErr
	}

	return toCartResponse(updatedCart), nil
}

func (uc *cartUseCase) DeleteCart(id uuid.UUID) e.ApiError {
	return uc.repo.DeleteCart(id)
}
