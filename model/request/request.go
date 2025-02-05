package request

import (
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/go-playground/validator/v10"
)

type CreateEstate struct {
	Length int `json:"length" validate:"required,numeric,min=1,max=50000"`
	Width  int `json:"width" validate:"required,numeric,min=1,max=50000"`
}

func ValidateCreateEstate(req generated.PostEstateJSONRequestBody, validator *validator.Validate) error {
	createEstate := CreateEstate{
		Length: req.Length,
		Width:  req.Width,
	}

	return validator.Struct(createEstate)
}

type CreateTree struct {
	Height int `json:"height" validate:"required,numeric,min=1,max=30"`
	X      int `json:"x" validate:"required,numeric,min=1"`
	Y      int `json:"y" validate:"required,numeric,min=1"`
}

func ValidateCreateTree(req generated.PostEstateIdTreeJSONRequestBody, validator *validator.Validate) error {
	createEstate := CreateTree{
		Height: req.Height,
		X:      req.X,
		Y:      req.Y,
	}

	return validator.Struct(createEstate)
}
