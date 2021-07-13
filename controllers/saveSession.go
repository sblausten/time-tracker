package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type SaveResponse struct {
	UserId      string `json: userId`
	SessionName string `json: name`
	Duration    string `json: duration`
}

type Session struct {
	UserId      string `json: userId`
	SessionName string `json: name`
	Start       string `json: start`
	End         string `json: end`
}

func SaveSession(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var session Session

	err := json.NewDecoder(r.Body).Decode(&session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Error decoding request, %v", err)
		return
	}

	saved := SaveResponse{
		UserId:      vars["userId"],
		SessionName: session.SessionName,
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(saved)
}
