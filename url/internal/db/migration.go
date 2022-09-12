package db

import (
	"database/sql"
	"github.com/rubenv/sql-migrate"
)

var (
	migrations = migrate.FileMigrationSource{
		Dir: "migration/url",
	}
)

func Migrate(db *sql.DB) (int, error) {
	return migrate.Exec(db, "postgres", migrations, migrate.Up)
}
