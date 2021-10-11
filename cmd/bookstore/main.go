package main

import (
	"fmt"
	"github.com/abhinav812/cloudy-bookstore/internal/config"
	"github.com/abhinav812/cloudy-bookstore/internal/router"
	"log"
	"net/http"
)

func main() {
	appConf := config.AppConfig()

	appRouter := router.New()

	address := fmt.Sprintf(":%d", appConf.Server.Port)

	log.Printf("Starting server %s\n", address)

	s := &http.Server{
		Addr:         address,
		Handler:      appRouter,
		ReadTimeout:  appConf.Server.TimeoutRead,
		WriteTimeout: appConf.Server.TimeoutWrite,
		IdleTimeout:  appConf.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}
