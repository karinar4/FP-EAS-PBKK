package merk

import (
	"errors"

	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/pkg/e"
	"gorm.io/gorm"
)

type IMerkRepository interface {
	CreateMerk(*MerkModel) (*MerkModel, e.ApiError)
	GetAllMerk() ([]MerkModel, e.ApiError)
	GetMerkByID(uuid.UUID) (*MerkModel, e.ApiError)
	UpdateMerk(*MerkModel) (*MerkModel, e.ApiError)
	DeleteMerk(uuid.UUID) e.ApiError
}

type merkRepository struct {
	db *gorm.DB
}

func NewMerkRepository(db *gorm.DB) *merkRepository {
	return &merkRepository{db}
}

func (r *merkRepository) CreateMerk(data *MerkModel) (*MerkModel, e.ApiError) {
	if err := r.db.Create(data).Error; err != nil {
		return nil, e.NewApiError(e.ErrDatabaseCreateFailed, err.Error())
	}
	return data, nil
}

func (r *merkRepository) GetAllMerk() ([]MerkModel, e.ApiError) {
	var merks []MerkModel
	if err := r.db.Find(&merks).Error; err != nil {
		return nil, e.NewApiError(e.ErrDatabaseFetchFailed, err.Error())
	}
	return merks, nil
}

func (r *merkRepository) GetMerkByID(id uuid.UUID) (*MerkModel, e.ApiError) {
	merk := &MerkModel{}
	if err := r.db.Where("id = ?", id).First(merk).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, e.NewApiError(e.ErrNotFound, "Merk not found")
		}
		return nil, e.NewApiError(e.ErrDatabaseFetchFailed, err.Error())
	}
	return merk, nil
}

func (r *merkRepository) UpdateMerk(data *MerkModel) (*MerkModel, e.ApiError) {
	if err := r.db.Save(data).Error; err != nil {
		return nil, e.NewApiError(e.ErrDatabaseUpdateFailed, err.Error())
	}
	return data, nil
}

func (r *merkRepository) DeleteMerk(id uuid.UUID) e.ApiError {
	result := r.db.Delete(&MerkModel{}, id)
	if result.Error != nil {
		return e.NewApiError(e.ErrDatabaseDeleteFailed, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return e.NewApiError(e.ErrNotFound, "Merk not found")
	}
	return nil
}
