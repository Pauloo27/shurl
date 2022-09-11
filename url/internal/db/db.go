package db

import (
	"database/sql"
	"fmt"

	"github.com/Pauloo27/shurl/url/internal/service"
	_ "github.com/lib/pq"
)

func Connect(config *service.Config) (*sql.DB, error) {
	//connStr := "postgres://postgres:password@localhost/DB_1?sslmode=disable"
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.PG.User, config.PG.Password, config.PG.Host, config.PG.Port, config.PG.Database,
	)
	return sql.Open("postgres", connStr)
}
