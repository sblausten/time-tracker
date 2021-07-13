package server

import (
	"github.com/gorilla/mux"
	"github.com/sblausten/time-tracker/controllers"
	"github.com/sblausten/time-tracker/dao"
)

func NewRouter(sessionDao dao.SessionDaoInterface) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/v1/users/{userId}/session", logRequest(controllers.SaveSession(sessionDao))).Methods("POST")
	r.HandleFunc("/v1/users/{userId}/sessions", logRequest(controllers.GetSessions(sessionDao))).Methods("GET")
	return r
}
