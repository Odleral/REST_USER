package app

//func HandleLive(w http.ResponseWriter, _ *http.Request){
//	writeHealthy(w)
//}

//func (app *App) HandleReady(w http.ResponseWriter, r *http.Request){
//	if err := app.db.DB(); err != nil {
//		app.logger.Fatal().Err(err).Msg("")
//		writeUnhelthy(w)
//
//		return
//	}
//	writeHealthy(w)
//}
//
//func writeHealthy(w http.ResponseWriter){
//	w.Header().Set("Content-Type", "text/plain")
//	w.WriteHeader(http.StatusOK)
//	_, err := w.Write([]byte("."))
//	if err != nil {
//		log.Fatal(err)
//
//		return
//	}
//}
//
//func writeUnhelthy(w http.ResponseWriter){
//	w.Header().Set("Content-Type", "text/plain")
//	w.WriteHeader(http.StatusInternalServerError)
//	_, err := w.Write([]byte("."))
//	if err != nil {
//		log.Fatal(err)
//
//		return
//	}
//}

