package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mct-joken/kojs5-backend/pkg/server/controller"
	"github.com/mct-joken/kojs5-backend/pkg/server/controller/model"
	"github.com/mct-joken/kojs5-backend/pkg/server/responses"
)

type ProblemHandlers struct {
	controller controller.ProblemController
}

func NewProblemHandlers(controller controller.ProblemController) *ProblemHandlers {
	return &ProblemHandlers{controller: controller}
}

func (h *ProblemHandlers) CreateProblem(c echo.Context) error {
	req := model.CreateProblemRequestJSON{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, responses.InvalidRequestErrorResponseJSON)
	}

	res, err := h.controller.CreateProblem(req)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponseJSON)
	}
	return c.JSON(http.StatusCreated, res)
}

func (h *ProblemHandlers) FindByID(c echo.Context) error {
	id := c.Param("id")
	res, err := h.controller.FindByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponseJSON)
	}

	return c.JSON(http.StatusOK, res)
}
