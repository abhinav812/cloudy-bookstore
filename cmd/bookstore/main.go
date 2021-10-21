package main

import (
	"fmt"
	"net/http"

	"github.com/abhinav812/cloudy-bookstore/internal/app"

	"github.com/abhinav812/cloudy-bookstore/internal/config"
	"github.com/abhinav812/cloudy-bookstore/internal/router"
	lr "github.com/abhinav812/cloudy-bookstore/internal/util/logger"
)

func main() {
	appConf := config.AppConfig()

	log := lr.New(appConf.Debug)

	application := app.New(log)

	appRouter := router.New(application)

	address := fmt.Sprintf(":%d", appConf.Server.Port)

	log.Info().Msgf("Starting server %s\n", address)

	s := &http.Server{
		Addr:         address,
		Handler:      appRouter,
		ReadTimeout:  appConf.Server.TimeoutRead,
		WriteTimeout: appConf.Server.TimeoutWrite,
		IdleTimeout:  appConf.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("Server startup failed")
	}
}
