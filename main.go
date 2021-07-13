package main

import (
	"context"
	"github.com/gorilla/handlers"
	"github.com/sblausten/time-tracker/config"
	"github.com/sblausten/time-tracker/controllers"
	"github.com/sblausten/time-tracker/dao"

	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	config := config.Config{}.Build()

	dbClient := dao.BuildClient(config, ctx)
	defer dbClient.Disconnect(ctx)

	srv := newServer()
	go func() {
		log.Print("Listening on port 8080...")
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	ch := make(chan os.Signal, 1)

	signal.Notify(ch, os.Interrupt)
	<-ch

	srv.Shutdown(ctx)

	log.Println("Shutting down...")
	os.Exit(0)
}

func newServer() *http.Server {
	r := newRouter()
	corsObj:=handlers.AllowedOrigins([]string{"http://localhost:3000"})

	return &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      handlers.CORS(corsObj)(r),
	}
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/v1/users/{userId}/session", logRequest(controllers.SaveSession)).Methods("POST")
	return r
}

func logRequest(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Recieved request to %s %s from %s", r.Method, r.URL.Path, r.Referer())
		f(w, r)
	}
}

