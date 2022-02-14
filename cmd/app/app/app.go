package app

import (
	"gitlab.com/odleral/geoportal-go/util/logger"
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/gorm"
)

const (
	appErrDataAccessFailure   = "data access failure"
	appErrJsonCreationFailure = "json creation failure"
	appErrDataCreationFailure = "data creation failure"
	appErrFormDecodingFailure = "form decoding failure"
	appErrDataUpdateFailure   = "data update failure"
)

type App struct {
	logger *logger.Logger
	db     *gorm.DB
	validator *validator.Validate
}

func New(logger *logger.Logger, db *gorm.DB, validator *validator.Validate) *App {
	return &App{
		logger: logger,
		db:     db,
		validator: validator,
	}
}

func (app *App) Logger() *logger.Logger {
	return app.logger
}
