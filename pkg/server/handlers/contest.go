package handlers

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/labstack/echo/v4"
	"github.com/mct-joken/kojs5-backend/pkg/server/controller"
	"github.com/mct-joken/kojs5-backend/pkg/server/controller/model"
	"github.com/mct-joken/kojs5-backend/pkg/server/responses"
)

type ContestHandlers struct {
	controller controller.ContestController
	logger     *zap.Logger
}

func NewContestHandlers(controller controller.ContestController, logger *zap.Logger) *ContestHandlers {
	return &ContestHandlers{controller: controller, logger: logger}
}

func (h *ContestHandlers) CreateContest(c echo.Context) error {
	req := model.CreateContestRequestJSON{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, responses.InvalidRequestErrorResponseJSON)
	}
	res, err := h.controller.CreateContest(req)
	if err != nil {
		// ToDo: エラーの種類を判別する
		// e.g: タイトルの長さが正しくありません
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponseJSON)
	}
	return c.JSON(http.StatusCreated, res)
}

func (h *ContestHandlers) FindContestByID(c echo.Context) error {
	id := c.Param("id")
	res, err := h.controller.FindContestByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponseJSON)
	}
	return c.JSON(http.StatusOK, res)
}

func (h *ContestHandlers) FindContest(c echo.Context) error {
	res, err := h.controller.FindContest()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponseJSON)
	}
	return c.JSON(http.StatusOK, res)
}
