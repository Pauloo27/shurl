package service

import "database/sql"

type URLService struct {
	Config
	DB *sql.DB
}

func NewService(config *Config, db *sql.DB) *URLService {
	return &URLService{
		Config: *config,
		DB:     db,
	}
}
