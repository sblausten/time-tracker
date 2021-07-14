package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sblausten/time-tracker/dao"
	"log"
	"net/http"
)

type SessionSummary struct {
	SessionName string `json:"name"`
	Start       string `json:"start"`
	End         string `json:"end"`
	Duration    int64 `json:"duration"`
}

type SessionsResponse struct {
	UserId      string `json:"userId"`
	Sessions  	[]SessionSummary `json:"sessions"`
}

func GetSessions(sessionDao dao.SessionDaoInterface) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userId := vars["userId"]
		if userId == "" {
			var errorMessage = "No user id in request url"
			http.Error(w, errorMessage, http.StatusBadRequest)
			log.Printf("SaveSession - %s", errorMessage)
			return
		}

		retrievedSessions, err := sessionDao.GetAllSessions(userId)
		if err != nil {
			var errorMessage = "Failed to find previous sessions for user " + userId
			http.Error(w, errorMessage, http.StatusInternalServerError)
			log.Printf("SaveSession - Error saving session: %s", err.Error())
			return
		}

		var sessionSummaries []SessionSummary
		for _, session := range retrievedSessions {
			newSession := SessionSummary{
				SessionName: session.SessionName,
				Start:       session.Start,
				End:         session.End,
				Duration:    session.Duration,
			}
			sessionSummaries = append(sessionSummaries, newSession)
		}

		sessionsRes := SessionsResponse{
			UserId: vars["userId"],
			Sessions: sessionSummaries,
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(sessionsRes)
	}
}