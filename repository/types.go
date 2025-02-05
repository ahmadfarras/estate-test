// This file contains types that are used in the repository layer.
package repository

import (
	"github.com/SawitProRecruitment/UserService/model"
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
	model.Estate
}

type CreateTreeInput struct {
	ID       uuid.UUID
	EstateID uuid.UUID
	Height   int
	X        int
	Y        int
}

type GetAllTreesByEstateIDInput struct {
	EstateID uuid.UUID
}

type GetAllTreesByEstateIDOutput struct {
	Trees []model.Tree
	Total int
}

type GetTreeByEstateIDAndCoordinateInput struct {
	EstateID uuid.UUID
	X        int
	Y        int
}

type GetTreeByEstateIDAndCoordinateOutput struct {
	Tree model.Tree
}
