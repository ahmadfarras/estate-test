package error

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ErrorHandler(ctx echo.Context, err error) error {

	if _, ok := err.(validator.ValidationErrors); ok {
		return ctx.JSON(http.StatusBadRequest,
			map[string]string{"error": err.Error()})
	}

	switch err {
	case ErrEstateNotFound:
		return ctx.JSON(http.StatusNotFound,
			map[string]string{"error": "Resource Not Found"})
	case ErrTreePositionOutOfBoundary,
		ErrTreePositionNegative,
		ErrTreeHeightNegative,
		ErrTreeAlreadyPlanted:
		return ctx.JSON(http.StatusBadRequest,
			map[string]string{"error": err.Error()})
	default:
		return ctx.JSON(http.StatusInternalServerError,
			map[string]string{"error": "Internal Server Error"})
	}
}
