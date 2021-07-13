package server

import (
	"github.com/rs/cors"
	"github.com/sblausten/time-tracker/dao"
	"net/http"
	"time"
)

func NewServer(sessionDao dao.SessionDaoInterface) *http.Server {
	r := NewRouter(sessionDao)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowCredentials: false,
	})
	handler := c.Handler(r)

	return &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      handler,
	}
}
