package handlers

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/labstack/echo/v4"
	"github.com/poporonnet/kojsx-backend/pkg/server/controller"
	"github.com/poporonnet/kojsx-backend/pkg/server/controller/model"
	"github.com/poporonnet/kojsx-backend/pkg/server/responses"
)

type UserHandlers struct {
	controller controller.UserController
	auth       controller.AuthController
	logger     *zap.Logger
}

func NewUserHandlers(
	userController controller.UserController,
	authController controller.AuthController,
	logger *zap.Logger,
) *UserHandlers {
	return &UserHandlers{userController, authController, logger}
}

func (h *UserHandlers) CreateUser(c echo.Context) error {
	req := model.CreateUserRequestJSON{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, responses.InvalidRequestErrorResponseJSON)
	}

	res, err := h.controller.Create(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponseJSON)
	}

	return c.JSON(http.StatusCreated, res)
}

func (h *UserHandlers) FindByID(c echo.Context) error {
	i := c.Param("id")
	res, err := h.controller.FindByID(i)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponseJSON)
	}
	return c.JSON(http.StatusOK, res)
}

func (h *UserHandlers) FindAllUser(c echo.Context) error {
	res, err := h.controller.FindAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponseJSON)
	}

	return c.JSON(http.StatusOK, res)
}

func (h *UserHandlers) Login(c echo.Context) error {
	req := model.LoginRequestJSON{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, responses.InvalidRequestErrorResponseJSON)
	}

	res, err := h.auth.Login(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponseJSON)
	}

	return c.JSON(http.StatusCreated, res)
}

func (h *UserHandlers) Verify(c echo.Context) error {
	t := c.Param("token")
	ok, err := h.auth.Verify(t)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.InvalidRequestErrorResponseJSON)
	}
	if !ok {
		return c.JSON(http.StatusBadRequest, responses.InvalidRequestErrorResponseJSON)
	}
	return c.NoContent(http.StatusOK)
}
