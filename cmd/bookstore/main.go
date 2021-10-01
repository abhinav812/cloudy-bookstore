package main

import (
	"fmt"
	"github.com/abhinav812/cloudy-bookstore/config"
	"log"
	"math/rand"
	"net/http"
)

func main() {
	appConf := config.AppConfig()

	mux := http.NewServeMux()
	mux.HandleFunc("/", Greet)

	address := fmt.Sprintf(":%d", appConf.Server.Port)

	log.Printf("Starting server %s\n", address)

	s := &http.Server{
		Addr:         address,
		Handler:      mux,
		ReadTimeout:  appConf.Server.TimeoutRead,
		WriteTimeout: appConf.Server.TimeoutWrite,
		IdleTimeout:  appConf.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}

func Greet(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w, "Hello There!!\n"+
		"Here is a magic number for you %d", rand.Intn(10000))
	if err != nil {
		panic(err)
	}
}
