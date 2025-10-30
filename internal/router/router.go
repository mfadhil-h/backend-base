package router

import (
	"backend-base/internal/handler"
	"backend-base/internal/repository"
	"backend-base/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB, rmq *amqp.Connection) {
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "pong"})
	})

	// Health1
	hh := handler.NewHealthHandler(db, rmq)
	e.GET("/health", hh.HealthCheck)

	// Users
	ur := repository.NewUserRepository(db)
	us := service.NewUserService(ur)
	uh := handler.NewUserHandler(us)

	e.GET("/users", uh.GetUsers)
	e.POST("/users", uh.CreateUser)
}
