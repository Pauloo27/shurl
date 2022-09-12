package service

import "database/sql"

const Version = "0.0.1"

type URLService struct {
	Config
	DB      *sql.DB
	Version string
}

func NewService(config *Config, db *sql.DB) *URLService {
	return &URLService{
		Config:  *config,
		DB:      db,
		Version: Version,
	}
}
