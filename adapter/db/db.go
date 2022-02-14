package db

import (
	"database/sql"
	"fmt"
	"gitlab.com/odleral/geoportal-go/config"
)

func New(conf *config.Conf) (*sql.DB, error) {
	cfg := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.DB.Host,
		conf.DB.Port,
		conf.DB.Username,
		conf.DB.DBName,
	)
	return sql.Open("postgres", cfg)
}
