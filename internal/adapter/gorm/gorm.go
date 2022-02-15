package gorm

import (
	"fmt"
	"gitlab.com/odleral/geoportal-go/config"
	"gitlab.com/odleral/geoportal-go/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(conf *config.Conf) (*gorm.DB, error){
	cfg := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", conf.DB.Username, conf.DB.Password, conf.DB.Host, conf.DB.Port, conf.DB.DBName) //nolint:lll
	db, err := gorm.Open(postgres.Open(cfg), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		return nil, err
	} //nolint:errcheck

	return db, err
}
