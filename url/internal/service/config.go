package service

type Config struct {
	Http struct {
		Port int `json:"port"`
	} `json:"http"`
	PG struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Database string `json:"database"`
	} `json:"pg"`
}
