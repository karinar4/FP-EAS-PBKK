package category

import (
	"errors"

	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/pkg/e"
	"gorm.io/gorm"
)

type ICategoryRepository interface {
	CreateCategory(*CategoryModel) (*CategoryModel, e.ApiError)
	GetAllCategory() ([]CategoryModel, e.ApiError)
	GetCategoryByID(uuid.UUID) (*CategoryModel, e.ApiError)
	UpdateCategory(*CategoryModel) (*CategoryModel, e.ApiError)
	DeleteCategory(uuid.UUID) e.ApiError
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) CreateCategory(data *CategoryModel) (*CategoryModel, e.ApiError) {
	if err := r.db.Create(data).Error; err != nil {
		return nil, e.NewApiError(e.ErrDatabaseCreateFailed, err.Error())
	}
	return data, nil
}

func (r *categoryRepository) GetAllCategory() ([]CategoryModel, e.ApiError) {
	var categories []CategoryModel
	if err := r.db.Find(&categories).Error; err != nil {
		return nil, e.NewApiError(e.ErrDatabaseFetchFailed, err.Error())
	}
	return categories, nil
}

func (r *categoryRepository) GetCategoryByID(id uuid.UUID) (*CategoryModel, e.ApiError) {
	category := &CategoryModel{}
	if err := r.db.Where("id = ?", id).First(category).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, e.NewApiError(e.ErrNotFound, "Category not found")
		}
		return nil, e.NewApiError(e.ErrDatabaseFetchFailed, err.Error())
	}
	return category, nil
}

func (r *categoryRepository) UpdateCategory(data *CategoryModel) (*CategoryModel, e.ApiError) {
	if err := r.db.Save(data).Error; err != nil {
		return nil, e.NewApiError(e.ErrDatabaseUpdateFailed, err.Error())
	}
	return data, nil
}

func (r *categoryRepository) DeleteCategory(id uuid.UUID) e.ApiError {
	result := r.db.Delete(&CategoryModel{}, id)
	if result.Error != nil {
		return e.NewApiError(e.ErrDatabaseDeleteFailed, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return e.NewApiError(e.ErrNotFound, "Category not found")
	}
	return nil
}
