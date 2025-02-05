package handler

import (
	"net/http"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// GetEstateIdDronePlan implements generated.ServerInterface.
func (s *Server) GetEstateIdDronePlan(ctx echo.Context, id uuid.UUID) error {
	var (
		resp    generated.GetDronePlaneResponse
		context = ctx.Request().Context()
	)

	resp, err := s.EstateUsecase.CalculateTravelDistance(context, id)
	if err != nil {
		logrus.WithContext(context).WithError(err).Error("failed to get estate stats")
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	return ctx.JSON(http.StatusOK, resp)
}

// GetEstateIdStats implements generated.ServerInterface.
func (s *Server) GetEstateIdStats(ctx echo.Context, id uuid.UUID) error {
	var (
		resp    generated.GetEstateStatResponse
		context = ctx.Request().Context()
	)

	resp, err := s.EstateUsecase.GetEstateStats(context, id)
	if err != nil {
		logrus.WithContext(context).WithError(err).Error("failed to get estate stats")
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	return ctx.JSON(http.StatusOK, resp)
}

// PostEstate implements generated.ServerInterface.
func (s *Server) PostEstate(ctx echo.Context) error {
	var (
		req     generated.PostEstateJSONRequestBody
		resp    generated.CreateEstateResponse
		context = ctx.Request().Context()
	)

	if err := ctx.Bind(&req); err != nil {
		logrus.WithContext(context).WithError(err).Error("failed to bind request")
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Request"})
	}

	resp, err := s.EstateUsecase.CreateEstate(context, req)
	if err != nil {
		logrus.WithContext(context).WithError(err).Error("failed to create estate")
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	return ctx.JSON(http.StatusOK, resp)
}

// PostEstateIdTree implements generated.ServerInterface.
func (s *Server) PostEstateIdTree(ctx echo.Context, id uuid.UUID) error {
	var (
		req     generated.PostEstateIdTreeJSONRequestBody
		resp    generated.CreateTreeResponse
		context = ctx.Request().Context()
	)

	if err := ctx.Bind(&req); err != nil {
		logrus.WithContext(context).WithError(err).Error("failed to bind request")
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Request"})
	}

	resp, err := s.EstateUsecase.CreateTree(context, id, req)
	if err != nil {
		logrus.WithContext(context).WithError(err).Error("failed to create tree")
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	return ctx.JSON(http.StatusOK, resp)
}
