package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/abhinav812/cloudy-bookstore/migrations"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"

	"github.com/abhinav812/cloudy-bookstore/internal/dao/postgres"

	"github.com/abhinav812/cloudy-bookstore/internal/app"

	"github.com/abhinav812/cloudy-bookstore/internal/config"
	"github.com/abhinav812/cloudy-bookstore/internal/router"
	lr "github.com/abhinav812/cloudy-bookstore/internal/util/logger"
)

func main() {
	// Get the app config
	appConf, err := config.AppTomlConfig()
	if err != nil {
		panic(err)
	}

	// create logger
	log := lr.New(appConf.Logging.Debug)

	// connect to database
	dbStore, err := postgres.NewDBStore(appConf)
	if err != nil {
		log.Panic().Err(err)
		panic(err)
	}
	code, _ := dbStore.ConnStatusWithContext(context.TODO())

	// apply DB migration
	if err := migrations.ApplyDBMigration(dbStore.DbRef()); err != nil {
		log.Panic().Err(err).Msg("DB migrations failed")
	}

	// create GORM for connected database
	gormDB, err := createGorm(code, appConf, dbStore)
	if err != nil {
		log.Panic().Err(err)
	}

	// create application
	s := createApplication(log, gormDB, appConf)

	// Start application
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("Server startup failed")
	}
}

func createGorm(code codes.Code, appConf *config.TomlConfig, dbStore *postgres.DBStore) (*gorm.DB, error) {
	var gormDB *gorm.DB = nil
	var gormErr error = nil
	if code == codes.Internal {
		// Try getting gorm using conf
		gormDB, gormErr = postgres.GormFromConf(appConf)
	} else {
		gormDB, gormErr = postgres.GormFromDbConn(dbStore.DbRef(), appConf.Logging.Debug)
	}
	if gormErr != nil {
		return nil, gormErr
	}

	return gormDB, nil
}

func createApplication(log *lr.Logger, gormDB *gorm.DB, appConf *config.TomlConfig) *http.Server {
	application := app.New(log, gormDB)

	appRouter := router.New(application)

	address := fmt.Sprintf(":%d", appConf.Server.Port)

	log.Info().Msgf("Starting server %s\n", address)

	s := &http.Server{
		Addr:         address,
		Handler:      appRouter,
		ReadTimeout:  appConf.Server.ReadTimeout.Duration,
		WriteTimeout: appConf.Server.WriteTimeout.Duration,
		IdleTimeout:  appConf.Server.IdleTimeout.Duration,
	}

	return s
}
