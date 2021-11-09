package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/abhinav812/cloudy-bookstore/internal/config"
	lr "github.com/abhinav812/cloudy-bookstore/internal/util/logger"
	"log"
	// Mandatory to load postgres sql library
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
)

//DBStore - struct encapsulating database and logger
type DBStore struct {
	db     *sql.DB
	logger *lr.Logger
}

//NewDBStore - Initializes database based on the config.Conf and returns new instance of DBStore
func NewDBStore(conf *config.TomlConfig) (*DBStore, error) {
	logger := lr.New(conf.Logging.Debug)

	defer func() {
		if err := recover(); err != nil {
			//log.Fatalf("Error while connecting to DB: %v", err )
			logger.Error().Err(err.(error))
		}
	}()

	dbStore := &DBStore{}

	log.Printf("dbConfig hostname : %v, username: %v\n", conf.DB.Host, conf.DB.User)

	dbConnString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		conf.DB.Host,
		conf.DB.Port,
		conf.DB.User,
		conf.DB.DbName,
		conf.DB.Password)

	//log.Printf("Connecting to [%s]", dbConnString)
	logger.Info().Msgf("Connecting to [%s]", dbConnString)

	db, err := sql.Open("postgres", dbConnString)
	if err != nil {
		return dbStore, errors.Errorf("unable to connect to db, %v", err)
	}

	dbStore.db = db
	dbStore.logger = logger
	//log.Println("Initialized postgres database")
	logger.Info().Msg("Initialized postgres database")

	return dbStore, nil

}

//ConnStatusWithContext - Returns the connection status of database.
// Times out if the passed context have a timeout specified.
func (d *DBStore) ConnStatusWithContext(ctx context.Context) (codes.Code, error) {
	resp := codes.OK
	status := "Up"

	if err := d.db.PingContext(ctx); err != nil {
		status = "Down"
		resp = codes.Internal
	}

	d.logger.Info().Msgf("DB connection status %s", status)
	return resp, nil
}

//ConnStatus - Returns the connection status of database.
func (d *DBStore) ConnStatus() (codes.Code, error) {
	resp := codes.OK
	status := "Up"

	if err := d.db.Ping(); err != nil {
		status = "Down"
		resp = codes.Internal
	}

	d.logger.Info().Msgf("DB connection status %s", status)
	return resp, nil
}

//Close - Closes the database connection.
func (d *DBStore) Close() error {
	return d.db.Close()
}

//DbRef - returns the encapsulated database reference.
func (d *DBStore) DbRef() *sql.DB {
	return d.db
}
