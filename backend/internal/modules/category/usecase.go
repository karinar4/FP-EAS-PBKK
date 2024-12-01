package category

import (
	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/common"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/pkg/e"
)

type ICategoryUseCase interface {
	CreateCategory(*CreateCategoryRequest) (*CreateCategoryResponse, e.ApiError)
	GetAllCategory() (*GetAllCategoryResponse, e.ApiError)
	GetCategoryByID(uuid.UUID) (*GetCategoryResponse, e.ApiError)
	UpdateCategory(uuid.UUID, *UpdateCategoryRequest) (*UpdateCategoryResponse, e.ApiError)
	DeleteCategory(uuid.UUID) e.ApiError
}

type categoryUseCase struct {
	repo ICategoryRepository
}

func NewCategoryUseCase(repo ICategoryRepository) ICategoryUseCase {
	return &categoryUseCase{repo: repo}
}

func (uc *categoryUseCase) CreateCategory(req *CreateCategoryRequest) (*CreateCategoryResponse, e.ApiError) {
	category := &CategoryModel{
		BaseModels: common.BaseModels{
			ID: uuid.New(),
		},
		Name: req.Name,
	}

	result, err := uc.repo.CreateCategory(category)
	if err != nil {
		return nil, e.NewApiError(e.ErrDatabaseCreateFailed, err.Error())
	}

	return &CreateCategoryResponse{
		ID:   result.ID,
		Name: result.Name,
	}, nil
}

func (uc *categoryUseCase) GetAllCategory() (*GetAllCategoryResponse, e.ApiError) {
	categories, err := uc.repo.GetAllCategory()
	if err != nil {
		return nil, e.NewApiError(e.ErrDatabaseFetchFailed, err.Error())
	}

	var response []GetCategoryResponse
	for _, category := range categories {
		response = append(response, GetCategoryResponse{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	return &GetAllCategoryResponse{Categories: response}, nil
}

func (uc *categoryUseCase) GetCategoryByID(id uuid.UUID) (*GetCategoryResponse, e.ApiError) {
	category, err := uc.repo.GetCategoryByID(id)
	if err != nil {
		return nil, err
	}

	return &GetCategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}, nil
}

func (uc *categoryUseCase) UpdateCategory(id uuid.UUID, req *UpdateCategoryRequest) (*UpdateCategoryResponse, e.ApiError) {
	category, err := uc.repo.GetCategoryByID(id)
	if err != nil {
		return nil, err
	}

	category.Name = req.Name

	updatedCategory, updateErr := uc.repo.UpdateCategory(category)
	if updateErr != nil {
		return nil, updateErr
	}

	return &UpdateCategoryResponse{
		ID:   updatedCategory.ID,
		Name: updatedCategory.Name,
	}, nil
}

func (uc *categoryUseCase) DeleteCategory(id uuid.UUID) e.ApiError {
	return uc.repo.DeleteCategory(id)
}
