// This file contains types that are used in the repository layer.
package repository

import (
	"time"

	"github.com/google/uuid"
)

type GetTestByIdInput struct {
	Id string
}

type GetTestByIdOutput struct {
	Name string
}

type CreateEstateInput struct {
	ID     uuid.UUID
	Length int
	Width  int
}

type GetEstateByIdInput struct {
	ID uuid.UUID
}

type GetEstateByIdOutput struct {
	ID        uuid.UUID
	Length    int
	Width     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateTreeInput struct {
	ID       uuid.UUID
	EstateID uuid.UUID
	Height   int
	X        int
	Y        int
}
