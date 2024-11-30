package merk

import (
	"github.com/google/uuid"
	"github.com/karinar4/FP-EAS-PBKK/internal/modules/common"
	"github.com/karinar4/FP-EAS-PBKK/internal/pkg/e"
)

type IMerkUseCase interface {
	CreateMerk(*CreateMerkRequest) (*CreateMerkResponse, e.ApiError)
	GetAllMerk() (*GetAllMerkResponse, e.ApiError)
	GetMerkByID(uuid.UUID) (*GetMerkResponse, e.ApiError)
	UpdateMerk(uuid.UUID, *UpdateMerkRequest) (*UpdateMerkResponse, e.ApiError)
	DeleteMerk(uuid.UUID) e.ApiError
}

type merkUseCase struct {
	repo IMerkRepository
}

func NewMerkUseCase(repo IMerkRepository) IMerkUseCase {
	return &merkUseCase{repo: repo}
}

func (uc *merkUseCase) CreateMerk(req *CreateMerkRequest) (*CreateMerkResponse, e.ApiError) {
	merk := &MerkModel{
		BaseModels: common.BaseModels{
			ID: uuid.New(),
		},
		Nama: req.Nama,
	}

	result, err := uc.repo.CreateMerk(merk)
	if err != nil {
		return nil, e.NewApiError(e.ErrDatabaseCreateFailed, err.Error())
	}

	return &CreateMerkResponse{
		ID:   result.ID,
		Nama: result.Nama,
	}, nil
}

func (uc *merkUseCase) GetAllMerk() (*GetAllMerkResponse, e.ApiError) {
	merks, err := uc.repo.GetAllMerk()
	if err != nil {
		return nil, e.NewApiError(e.ErrDatabaseFetchFailed, err.Error())
	}

	var response []GetMerkResponse
	for _, merk := range merks {
		response = append(response, GetMerkResponse{
			ID:   merk.ID,
			Nama: merk.Nama,
		})
	}

	return &GetAllMerkResponse{Merks: response}, nil
}

func (uc *merkUseCase) GetMerkByID(id uuid.UUID) (*GetMerkResponse, e.ApiError) {
	merk, err := uc.repo.GetMerkByID(id)
	if err != nil {
		return nil, err
	}

	return &GetMerkResponse{
		ID:   merk.ID,
		Nama: merk.Nama,
	}, nil
}

func (uc *merkUseCase) UpdateMerk(id uuid.UUID, req *UpdateMerkRequest) (*UpdateMerkResponse, e.ApiError) {
	merk, err := uc.repo.GetMerkByID(id)
	if err != nil {
		return nil, err
	}

	merk.Nama = req.Nama

	updatedMerk, updateErr := uc.repo.UpdateMerk(merk)
	if updateErr != nil {
		return nil, updateErr
	}

	return &UpdateMerkResponse{
		ID:   updatedMerk.ID,
		Nama: updatedMerk.Nama,
	}, nil
}

func (uc *merkUseCase) DeleteMerk(id uuid.UUID) e.ApiError {
	return uc.repo.DeleteMerk(id)
}
