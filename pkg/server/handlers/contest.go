package handlers

import (
	"go.uber.org/zap"
	"net/http"

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
		h.logger.Sugar().Errorf("%s", err)
		return c.JSON(http.StatusBadRequest, responses.InvalidRequestErrorResponseJSON)
	}
	res, err := h.controller.CreateContest(req)
	if err != nil {
		// ToDo: エラーの種類を判別する
		// e.g: タイトルの長さが正しくありません
		h.logger.Sugar().Errorf("%s", err)
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponseJSON)
	}
	return c.JSON(http.StatusCreated, res)
}

func (h *ContestHandlers) FindContestByID(c echo.Context) error {
	id := c.Param("id")
	res, err := h.controller.FindContestByID(id)
	if err != nil {
		h.logger.Sugar().Errorf("%s", err)
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponseJSON)
	}
	return c.JSON(http.StatusOK, res)
}
