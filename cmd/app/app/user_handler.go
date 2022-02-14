package app

import (
	"encoding/json"
	"fmt" //nolint:gci
	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
	"gitlab.com/odleral/geoportal-go/model"
	"gitlab.com/odleral/geoportal-go/repository"
	"gorm.io/gorm"
	"net/http" //nolint:gci
)

func (app *App) HandleListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := repository.ListUser(app.db)
	if err != nil {
		app.logger.Warn().Err(err).Msg("")
		w.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprintf(w, `{"error":"%v"}`, appErrDataAccessFailure)
		if err != nil {
			app.logger.Error().Err(err).Msg("Fprintf error")
		}

		return
	}

	if users == nil {
		_, err := fmt.Fprintf(w, "[]")
		if err != nil {
			app.logger.Error().Err(err).Msg("Fprintf error")
		}

		return
	}

	dtos := users
	if err := json.NewEncoder(w).Encode(dtos); err != nil {
		app.logger.Warn().Err(err).Msg("")

		w.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprintf(w, `{"error":"%v"}`, appErrJsonCreationFailure)
		if err != nil {
			app.logger.Error().Err(err).Msg("Fprintf error")
		}

		return
	}
}

func (app *App) HandleCreateUser(w http.ResponseWriter, r *http.Request) {

	form := &model.UserForm{}

	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		app.logger.Warn().Err(err).Msg("")
		w.WriteHeader(http.StatusUnprocessableEntity)
		_, err := fmt.Fprintf(w, `{"error":"%v"}`, appErrFormDecodingFailure)
		if err != nil {
			app.logger.Error().Err(err).Msg("Fprintf error")
		}

		return
	}

	if err := app.validator.Struct(form); err != nil {
		app.logger.Warn().Err(err).Msg("")
		w.WriteHeader(http.StatusUnprocessableEntity)
		_, err := fmt.Fprintf(w, `{"error": "%v"}`, err.Error())
		if err != nil {
			app.logger.Error().Err(err).Msg("Fprintf error")
		}

		return
	}

	UserModel, err := form.ToModel()
	if err != nil {
		app.logger.Warn().Err(err).Msg("")
		w.WriteHeader(http.StatusUnprocessableEntity)
		_, err := fmt.Fprintf(w, `{"error": "%v"}`, appErrFormDecodingFailure)
		if err != nil {
			app.logger.Error().Err(err).Msg("Fprintf error")
		}

		return
	}

	user, err := repository.CreateUser(app.db, UserModel)
	if err != nil {
		app.logger.Warn().Err(err).Msg("")
		w.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprintf(w, `{"error":"%v"}`, appErrDataCreationFailure)
		if err != nil {
			app.logger.Error().Err(err).Msg("Fprintf error")
		}

		return
	}

	app.logger.Info().Msgf("New user created: %s", user.UUID)
	w.WriteHeader(http.StatusCreated)
	if _, err = w.Write([]byte(user.UUID.String())); err != nil {
		app.logger.Warn().Err(err).Msg("")
	}
}

func (app *App) HandleFindUser(w http.ResponseWriter, r *http.Request) {
	uuid := chi.URLParam(r, "uuid")
	if uuid == "" {
		app.logger.Info().Msgf("can not parse or empty UUID: %v", uuid)
		w.WriteHeader(http.StatusUnprocessableEntity)

		return
	}

	user, err := repository.FindUser(app.db, uuid)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)

			return
		}

		app.logger.Warn().Err(err).Msg("")

		w.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprintf(w, `{"error":"%v"}`, appErrDataAccessFailure)
		if err != nil {
			app.logger.Error().Err(err).Msg("Fprintf error")
		}

		return
	}

	dto := user.ToDto()
	if err := json.NewEncoder(w).Encode(dto); err != nil {
		app.logger.Warn().Err(err).Msg("")

		w.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprintf(w, `{"error":"%v"}`, appErrJsonCreationFailure)
		if err != nil {
			app.logger.Error().Err(err).Msg("Fprintf error")
		}

		return
	}
}

func (app *App) HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "uuid")
	if id == "" {
		app.logger.Warn().Msgf("can not parse or empty UUID: %v", id)
		w.WriteHeader(http.StatusUnprocessableEntity)

		return
	}

	form := &model.UserForm{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		app.logger.Warn().Err(err).Msg("")
		w.WriteHeader(http.StatusUnprocessableEntity)
		_, err := fmt.Fprintf(w, `{"error":"%v"}`, appErrFormDecodingFailure)
		if err != nil {
			app.logger.Warn().Err(err).Msg("Fprintf error")
		}

		return
	}

	userModel, err := form.ToModel()
	if err != nil {
		app.logger.Warn().Err(err).Msg("")
		w.WriteHeader(http.StatusUnprocessableEntity)
		_, err := fmt.Fprintf(w, `{"error":"%v"}`, appErrFormDecodingFailure)
		if err != nil {
			app.logger.Warn().Err(err).Msg("Fprintf error")
		}

		return
	}

	userModel.UUID, err = uuid.FromString(id)
	if err := repository.UpdateUser(app.db, userModel); err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)

			return
		}

		app.logger.Warn().Err(err).Msg("")

		w.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprintf(w, `{"error":"%v"}`, appErrDataUpdateFailure)
		if err != nil {
			app.logger.Warn().Err(err).Msg("Fprintf error")
		}
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func (app *App) HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	uuid := chi.URLParam(r, "uuid")
	if uuid == "" {
		app.logger.Warn().Msgf("can not parse or empty UUID: %v", uuid)
		w.WriteHeader(http.StatusUnprocessableEntity)

		return
	}

	if err := repository.DeleteUser(app.db, uuid); err != nil {
		app.logger.Warn().Err(err).Msg("")
		w.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprintf(w, `{"error":"%v"}`, appErrDataAccessFailure)
		if err != nil {
			app.logger.Info().Err(err).Msg("Fprintf error")
		}

		return
	}

	app.logger.Info().Msgf("User deleted %s", uuid)
	w.WriteHeader(http.StatusAccepted)
}
