package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

type HealthHandler struct {
	DB  *gorm.DB
	RMQ *amqp.Connection
}

func NewHealthHandler(db *gorm.DB, rmq *amqp.Connection) *HealthHandler {
	return &HealthHandler{DB: db, RMQ: rmq}
}

func (h *HealthHandler) HealthCheck(c echo.Context) error {
	sqlDB, err := h.DB.DB()
	dbStatus := "ok"
	if err != nil || sqlDB.Ping() != nil {
		dbStatus = "error"
	}

	rmqStatus := "ok"
	if h.RMQ.IsClosed() {
		rmqStatus = "error"
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":   "healthy",
		"database": dbStatus,
		"rabbitmq": rmqStatus,
	})
}
