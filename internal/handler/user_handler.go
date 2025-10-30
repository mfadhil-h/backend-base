package handler

import (
	"backend-base/internal/model"
	"backend-base/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	Service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{Service: s}
}

func (h *UserHandler) GetUsers(c echo.Context) error {
	users, err := h.Service.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if err := h.Service.CreateUser(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, user)
}
