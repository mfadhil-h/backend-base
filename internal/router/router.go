package router

import (
	"backend-base/internal/handler"
	"backend-base/internal/middleware"
	"backend-base/internal/repository"
	"backend-base/internal/service"
	"backend-base/internal/util"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB, rmq *amqp.Connection) {
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "pong"})
	})

	e.GET("/testjwt", func(c echo.Context) error {
		token, err := util.GenerateJWT(1, "test@example.com")
		if err != nil {
			return c.JSON(500, echo.Map{"error": err.Error()})
		}
		return c.JSON(200, echo.Map{"token": token})
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

	// Auth routes
	authService := service.NewAuthService(ur)
	authHandler := handler.NewAuthHandler(authService)
	e.POST("/auth/register", authHandler.Register)
	e.POST("/auth/login", authHandler.Login)

	authGroup := e.Group("/users")
	authGroup.Use(middleware.JWTAuth)

	// example protected route
	authGroup.GET("/me", func(c echo.Context) error {
		userID := c.Get("user_id").(uint)
		email := c.Get("email").(string)
		return c.JSON(http.StatusOK, echo.Map{
			"user_id": userID,
			"email":   email,
		})
	})
}
