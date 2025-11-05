package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"backend-base/internal/config"
	"backend-base/internal/database"
	"backend-base/internal/queue"
	"backend-base/internal/router"
	"backend-base/internal/util"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func main() {
	// Initialize logger
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Load config
	config.Load()

	if err := util.LoadKeys(); err != nil {
		log.Fatal().Err(err).Msg("Failed to load RSA keys")
	}

	db, errDB := database.InitPostgres()
	if errDB != nil {
		panic("failed to connect database: " + errDB.Error())
	}
	rmq := queue.InitRabbitMQ()
	defer rmq.Close()

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	router.SetupRoutes(e, db, rmq)

	port := viper.GetString("APP_PORT")
	if port == "" {
		port = "8080"
	}

	go func() {
		log.Info().Msg("🔥 Live reload test successful")
		log.Info().Msg("🔥 Air fully working!")
		log.Info().Msgf("Starting server on port %s", port)
		if err := e.Start(":" + port); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("Shutting down server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Info().Msg("Gracefully shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("Server forced to shutdown")
	}
}
