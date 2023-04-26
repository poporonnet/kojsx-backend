package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mct-joken/kojs5-backend/pkg/server/controller"
	"github.com/mct-joken/kojs5-backend/pkg/server/controller/model"
	"github.com/mct-joken/kojs5-backend/pkg/server/responses"
)

type UserHandlers struct {
	controller.UserController
}

func NewUserHandlers(userController controller.UserController) *UserHandlers {
	return &UserHandlers{userController}
}

func (h *UserHandlers) CreateUser(c echo.Context) error {
	req := model.CreateUserRequestJSON{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, responses.InvalidRequestErrorResponseJSON)
	}

	res, err := h.UserController.Create(req)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponseJSON)
	}

	return c.JSON(http.StatusCreated, res)
}
