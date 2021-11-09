package migrations

import (
	"database/sql"
	"embed"
	"github.com/pressly/goose/v3"
)

//go:embed *.sql
var embedMigrations embed.FS

//ApplyDBMigration - apply DB migration on supplied db connection
func ApplyDBMigration(dbStore *sql.DB) error {
	goose.SetBaseFS(embedMigrations)
	if err := goose.Up(dbStore, "."); err != nil {
		return err
	}
	return nil
}
