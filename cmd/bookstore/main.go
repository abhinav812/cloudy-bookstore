package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
	"net/http"

	"github.com/abhinav812/cloudy-bookstore/internal/dao/postgres"

	"github.com/abhinav812/cloudy-bookstore/internal/app"

	"github.com/abhinav812/cloudy-bookstore/internal/config"
	"github.com/abhinav812/cloudy-bookstore/internal/router"
	lr "github.com/abhinav812/cloudy-bookstore/internal/util/logger"
)

func main() {
	appConf := config.AppConfig()

	log := lr.New(appConf.Debug)

	dbStore, err := postgres.NewDBStore(appConf)
	if err != nil {
		log.Panic().Err(err)
		panic(err)
	}
	code, _ := dbStore.ConnStatusWithContext(context.TODO())

	var gormDB *gorm.DB = nil
	var gormErr error = nil
	if code == codes.Internal {
		// Try getting gorm using conf
		gormDB, gormErr = postgres.GormFromConf(appConf)
	} else {
		gormDB, gormErr = postgres.GormFromDbConn(dbStore.DbRef(), appConf.Debug)
	}
	if gormErr != nil {
		log.Panic().Err(err)
		panic(err)
	}

	application := app.New(log, gormDB)

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
