package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/capitan-beto/macbot/api"
	log "github.com/sirupsen/logrus"
)

func Health(w http.ResponseWriter, r *http.Request) {
	res := struct {
		Code int
		Msg  string
	}{
		Code: http.StatusOK,
		Msg:  "Looks fine to me!",
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
