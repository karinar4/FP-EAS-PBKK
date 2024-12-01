package merk

import "github.com/google/uuid"

type CreateMerkRequest struct {
	Nama string `json:"nama" binding:"required"`
}

type CreateMerkResponse struct {
	ID   uuid.UUID `json:"id"`
	Nama string    `json:"nama"`
}

type GetMerkResponse struct {
	ID   uuid.UUID `json:"id"`
	Nama string    `json:"nama"`
}

type GetAllMerkResponse struct {
	Merks []GetMerkResponse `json:"merks"`
}

type UpdateMerkRequest struct {
	Nama string `json:"nama" binding:"required"`
}

type UpdateMerkResponse struct {
	ID   uuid.UUID `json:"id"`
	Nama string    `json:"nama"`
}

type DeleteMerkResponse struct {
	Message string `json:"message"`
}
