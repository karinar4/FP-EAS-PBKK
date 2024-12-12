package brand

import (
	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/common"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/pkg/e"
)

type IBrandUseCase interface {
	CreateBrand(*CreateBrandRequest) (*CreateBrandResponse, e.ApiError)
	GetAllBrand() (*GetAllBrandResponse, e.ApiError)
	GetBrandByID(uuid.UUID) (*GetBrandResponse, e.ApiError)
	UpdateBrand(uuid.UUID, *UpdateBrandRequest) (*UpdateBrandResponse, e.ApiError)
	DeleteBrand(uuid.UUID) e.ApiError
}

type brandUseCase struct {
	repo IBrandRepository
}

func NewBrandUseCase(repo IBrandRepository) IBrandUseCase {
	return &brandUseCase{repo: repo}
}

func (uc *brandUseCase) CreateBrand(req *CreateBrandRequest) (*CreateBrandResponse, e.ApiError) {
	brand := &BrandModel{
		BaseModels: common.BaseModels{
			ID: uuid.New(),
		},
		Name: req.Name,
		Origin: req.Origin,
	}

	result, err := uc.repo.CreateBrand(brand)
	if err != nil {
		return nil, e.NewApiError(e.ErrDatabaseCreateFailed, err.Error())
	}

	return &CreateBrandResponse{
		ID:   result.ID,
		Name: result.Name,
		Origin: result.Origin,
	}, nil
}

func (uc *brandUseCase) GetAllBrand() (*GetAllBrandResponse, e.ApiError) {
	brands, err := uc.repo.GetAllBrand()
	if err != nil {
		return nil, e.NewApiError(e.ErrDatabaseFetchFailed, err.Error())
	}

	var response []GetBrandResponse
	for _, brand := range brands {
		response = append(response, GetBrandResponse{
			ID:   brand.ID,
			Name: brand.Name,
			Origin: brand.Origin,
		})
	}

	return &GetAllBrandResponse{Brands: response}, nil
}

func (uc *brandUseCase) GetBrandByID(id uuid.UUID) (*GetBrandResponse, e.ApiError) {
	brand, err := uc.repo.GetBrandByID(id)
	if err != nil {
		return nil, err
	}

	return &GetBrandResponse{
		ID:   brand.ID,
		Name: brand.Name,
		Origin: brand.Origin,
	}, nil
}

func (uc *brandUseCase) UpdateBrand(id uuid.UUID, req *UpdateBrandRequest) (*UpdateBrandResponse, e.ApiError) {
	brand, err := uc.repo.GetBrandByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != ""{
		brand.Name = req.Name
	}
	if req.Origin != ""{
		brand.Origin = req.Origin
	}

	updatedBrand, updateErr := uc.repo.UpdateBrand(brand)
	if updateErr != nil {
		return nil, updateErr
	}

	return &UpdateBrandResponse{
		ID:   updatedBrand.ID,
		Name: updatedBrand.Name,
		Origin: updatedBrand.Origin,
	}, nil
}

func (uc *brandUseCase) DeleteBrand(id uuid.UUID) e.ApiError {
	return uc.repo.DeleteBrand(id)
}
