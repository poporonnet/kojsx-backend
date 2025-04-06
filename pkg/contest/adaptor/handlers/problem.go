package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/poporonnet/kojsx-backend/pkg/contest/adaptor/controller"
	"github.com/poporonnet/kojsx-backend/pkg/contest/adaptor/controller/schema"
	"github.com/poporonnet/kojsx-backend/pkg/utils/id"
	errorSchema "github.com/poporonnet/kojsx-backend/pkg/utils/schema"
	"go.uber.org/zap"
)

type ProblemHandlers struct {
	controller controller.ProblemController
	logger     *zap.Logger
}

func NewProblemHandlers(controller controller.ProblemController, logger *zap.Logger) *ProblemHandlers {
	return &ProblemHandlers{controller: controller, logger: logger}
}

func (h *ProblemHandlers) CreateProblem(c echo.Context) error {
	req := schema.CreateProblemRequestJSON{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, errorSchema.InvalidRequestErrorResponseJSON)
	}

	res, err := h.controller.CreateProblem(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorSchema.InternalServerErrorResponseJSON)
	}
	return c.JSON(http.StatusCreated, res)
}

func (h *ProblemHandlers) FindByID(c echo.Context) error {
	id := c.Param("id")
	res, err := h.controller.FindByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorSchema.InternalServerErrorResponseJSON)
	}

	return c.JSON(http.StatusOK, res)
}

func (h *ProblemHandlers) FindByContestID(c echo.Context) error {
	i := c.Param("id")
	res, err := h.controller.FindByContestID(id.SnowFlakeID(i))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorSchema.InternalServerErrorResponseJSON)
	}

	return c.JSON(http.StatusOK, res)
}
