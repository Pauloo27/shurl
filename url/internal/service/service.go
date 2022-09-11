package service

type URLService struct {
	Http struct {
		Port int `json:"port"`
	} `json:"http"`
	PG struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Username string `json:"username"`
		Password string `json:"password"`
		Database string `json:"database"`
	} `json:"pg"`
}
