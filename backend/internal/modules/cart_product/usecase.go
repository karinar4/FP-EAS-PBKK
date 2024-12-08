package cart_product

import (
	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/pkg/e"
)

type ICartProductUseCase interface {
	CreateCartProduct(*CreateCartProductRequest) (*CreateCartProductResponse, e.ApiError)
	GetAllCartProducts() (*GetAllCartProductsResponse, e.ApiError)
	DeleteCartProduct(uuid.UUID, uuid.UUID) e.ApiError
}

type cartProductUseCase struct {
	repo ICartProductRepository
}

func NewCartProductUseCase(repo ICartProductRepository) ICartProductUseCase {
	return &cartProductUseCase{repo: repo}
}

func (uc *cartProductUseCase) CreateCartProduct(req *CreateCartProductRequest) (*CreateCartProductResponse, e.ApiError) {
	cartProduct := &CartProductModel{
		CartID:        req.CartID,
		ProductID:     req.ProductID,
		RentStartDate: req.RentStartDate,
		RentEndDate:   req.RentEndDate,
		Quantity:      req.Quantity,
		Price:         req.Price,
	}

	result, err := uc.repo.CreateCartProduct(cartProduct)
	if err != nil {
		return nil, e.NewApiError(e.ErrDatabaseCreateFailed, err.Error())
	}

	return &CreateCartProductResponse{
		CartID:        result.CartID,
		ProductID:     result.ProductID,
		RentStartDate: result.RentStartDate,
		RentEndDate:   result.RentEndDate,
		Quantity:      result.Quantity,
		Price:         result.Price,
	}, nil
}

func (uc *cartProductUseCase) GetAllCartProducts() (*GetAllCartProductsResponse, e.ApiError) {
	cartProducts, err := uc.repo.GetAllCartProducts()
	if err != nil {
		return nil, e.NewApiError(e.ErrDatabaseFetchFailed, err.Error())
	}

	var response []GetCartProductResponse
	for _, cartProduct := range cartProducts {
		response = append(response, GetCartProductResponse{
			CartID:        cartProduct.CartID,
			ProductID:     cartProduct.ProductID,
			RentStartDate: cartProduct.RentStartDate,
			RentEndDate:   cartProduct.RentEndDate,
			Quantity:      cartProduct.Quantity,
			Price:         cartProduct.Price,
		})
	}

	return &GetAllCartProductsResponse{CartProducts: response}, nil
}

func (uc *cartProductUseCase) DeleteCartProduct(cartID uuid.UUID, productID uuid.UUID) e.ApiError {
	return uc.repo.DeleteCartProduct(cartID, productID)
}
