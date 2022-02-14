package main

import (
	"fmt"
	gr "gitlab.com/odleral/geoportal-go/adapter/gorm"
	"gitlab.com/odleral/geoportal-go/cmd/app/app"
	"gitlab.com/odleral/geoportal-go/cmd/app/router"
	"gitlab.com/odleral/geoportal-go/config"
	lr "gitlab.com/odleral/geoportal-go/util/logger"
	vr "gitlab.com/odleral/geoportal-go/util/validator"
	"net/http" //nolint:gci
	"time"
)

func main() {

	appConf := config.AppConfig()

	logger := lr.New(appConf.Debug)

	logger.Info().Msgf(`Sleep %v seconds, waiting postgres`, appConf.Server.DBWait)
	time.Sleep(appConf.Server.DBWait)

	validator := vr.New()

	db, err := gr.New(appConf)
	if err != nil {
		logger.Fatal().Err(err).Msg("Fatal error *gorm.DB")

		return
	}


	application := app.New(logger, db, validator)

	appRouter := router.New(application)

	address := fmt.Sprintf(":%d", appConf.Server.Port)

	logger.Info().Msgf("Starting server %v", address)

	s := &http.Server{
		Addr:         address,
		Handler:      appRouter,
		ReadTimeout:  appConf.Server.TimeoutRead,
		WriteTimeout: appConf.Server.TimeoutWrite,
		IdleTimeout:  appConf.Server.TimeoutIdel,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal().Err(err).Msg("Server startup failed")
	}
}
