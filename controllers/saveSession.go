package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sblausten/time-tracker/dao"
	"github.com/sblausten/time-tracker/models"
	"log"
	"net/http"
)

type SaveResponse struct {
	UserId      string `json: userId`
	SessionName string `json: name`
	Duration    int64 `json: duration`
}

func SaveSession(sessionDao dao.SessionDaoInterface) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if vars["userId"] == "" {
			var errorMessage = "No user id in request url"
			http.Error(w, errorMessage, http.StatusBadRequest)
			log.Printf("SaveSession - %s", errorMessage)
			return
		}

		var session models.Session

		err := json.NewDecoder(r.Body).Decode(&session)
		if err != nil {
			var errorMessage = "Invalid request"
			http.Error(w, errorMessage, http.StatusBadRequest)
			log.Printf("SaveSession - Error decoding request: %s", err.Error())
			return
		}
		session.UserId = vars["userId"]

		err = sessionDao.InsertSession(session)
		if err != nil {
			var errorMessage = "Failed to save session"
			http.Error(w, errorMessage, http.StatusInternalServerError)
			log.Printf("SaveSession - Error saving session: %s", err.Error())
			return
		}

		saved := SaveResponse{
			UserId:      vars["userId"],
			SessionName: session.SessionName,
			Duration:    session.Duration,
		}
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(saved)
	}
}
