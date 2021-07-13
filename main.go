package main

import (
	"context"
	"github.com/sblausten/time-tracker/config"
	"github.com/sblausten/time-tracker/dao"
	"github.com/sblausten/time-tracker/server"

	"log"
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
	sessionCol := dao.GetCollection(dbClient, config.Db.Name, "session")
	sessionDao := dao.SessionDao{Collection: sessionCol}

	srv := server.NewServer(sessionDao)
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

