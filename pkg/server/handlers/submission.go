package handlers

import (
	"net/http"

	"github.com/mct-joken/kojs5-backend/pkg/utils/id"

	"github.com/labstack/echo/v4"
	"github.com/mct-joken/kojs5-backend/pkg/server/controller"
	"github.com/mct-joken/kojs5-backend/pkg/server/controller/model"
	"github.com/mct-joken/kojs5-backend/pkg/server/responses"
	"go.uber.org/zap"
)

type SubmissionHandlers struct {
	controller controller.SubmissionController
	logger     *zap.Logger
}

func NewSubmissionHandlers(controller controller.SubmissionController, logger *zap.Logger) *SubmissionHandlers {
	return &SubmissionHandlers{
		controller: controller,
		logger:     logger,
	}
}

func (h SubmissionHandlers) CreateSubmission(c echo.Context) error {
	req := model.CreateSubmissionRequestJSON{}
	if err := c.Bind(&req); err != nil {
		h.logger.Sugar().Errorf("%s", err)
		return c.JSON(http.StatusBadRequest, responses.InvalidRequestErrorResponseJSON)
	}
	// ToDo: 提出したユーザー(コンテスタント)IDを渡す
	res, err := h.controller.CreateSubmission("", req)
	if err != nil {
		h.logger.Sugar().Errorf("%s", err)
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponseJSON)
	}

	return c.JSON(http.StatusCreated, res)
}

func (h SubmissionHandlers) FindByID(c echo.Context) error {
	i := c.Param("submissionId")
	res, err := h.controller.FindByID(id.SnowFlakeID(i))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponseJSON)
	}
	return c.JSON(http.StatusOK, res)
}

func (h SubmissionHandlers) GetTask(c echo.Context) error {
	res, err := h.controller.FindTask()
	if err != nil {
		// ToDo: うまくunwrap出来ていない問題を修正する
		if err.Error() == "failed to find task: not found" {
			h.logger.Sugar().Infof("no judge task: %s", err)
			return c.NoContent(http.StatusNoContent)
		}
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponseJSON)
	}

	return c.JSON(http.StatusOK, res)
}

func (h SubmissionHandlers) CreateSubmissionResult(c echo.Context) error {
	req := model.CreateSubmissionResultRequestJSON{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponseJSON)
	}

	err := h.controller.CreateSubmissionResult(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponseJSON)
	}
	return c.NoContent(http.StatusNoContent)
}

func (h SubmissionHandlers) FindSubmissionByContestID(c echo.Context) error {
	i := c.Param("id")
	res, err := h.controller.FindByContestID(i)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponseJSON)
	}
	return c.JSON(http.StatusOK, res)
}
