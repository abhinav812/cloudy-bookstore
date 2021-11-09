package postgres

import (
	"database/sql"
	"fmt"

	"github.com/abhinav812/cloudy-bookstore/internal/config"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// GormFromConf - Creates new GORM connection based on passed DB configuration
func GormFromConf(conf *config.TomlConfig) (*gorm.DB, error) {
	gormConfig := &gorm.Config{}
	if conf.Logging.Debug {
		gormConfig = &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		}
	}

	dbConnString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		conf.DB.Host,
		conf.DB.Port,
		conf.DB.User,
		conf.DB.DbName,
		conf.DB.Password)

	gormDB, err := gorm.Open(postgres.Open(dbConnString), gormConfig)
	if err != nil {
		return nil, errors.Errorf("unable to connect to db, %v", err)
	}

	return gormDB, nil
}

// GormFromDbConn - Create a new Postgres GORM connection from existing postgres db connection
func GormFromDbConn(db *sql.DB, debug bool) (*gorm.DB, error) {
	gormConfig := &gorm.Config{}
	if debug {
		gormConfig = &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		}
	}
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), gormConfig)
	if err != nil {
		return nil, errors.Errorf("unable to connect to db, %v", err)
	}

	return gormDB, nil
}
