package dao

import (
	"context"
	"github.com/sblausten/time-tracker/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type SessionDaoInterface interface {
	BuildSessionIndexes()
	InsertSession(session models.Session) error
	GetAllSessions(userId string) ([]models.Session, error)
}

type SessionDao struct {
	Collection *mongo.Collection
}

func (d SessionDao) BuildSessionIndexes() {
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)

	indexModels := []mongo.IndexModel{
		{
			Keys:    bson.M{"start": 1},
			Options: nil,
		},
		{
			Keys:    bson.M{"userId": 1},
			Options: nil,
		},
	}

	indexes, err := d.Collection.Indexes().CreateMany(ctx, indexModels)
	if err != nil {
		log.Println("BuildSessionIndexes - Error creating indexes:", err)
	} else {
		log.Printf("BuildSessionIndexes - Created indexes %v on collection %s \n", indexes, d.Collection.Name())
	}
}

func (d SessionDao) InsertSession(digest models.Session) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	data, err := bson.Marshal(digest)
	if err != nil {
		return err
	}

	_, err = d.Collection.InsertOne(ctx, data)
	if err != nil {
		log.Printf("InsertSession - insert failed with error: %e", err)
	} else {
		log.Printf("InsertSession - inserted new digest record for user: %+v", digest)
	}

	return err
}

func (d SessionDao) GetAllSessions(userId string) ([]models.Session, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{"userId", userId}}
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"requestedAt", 1}})

	var results []models.Session

	cur, err := d.Collection.Find(ctx, filter, findOptions)
	if err != nil {
		log.Printf("GetAllSessions - lookup failed with Find error: %e", err)
		return nil, err
	}

	for cur.Next(ctx) {
		var session models.Session
		err := cur.Decode(&session)
		if err != nil {
			log.Printf("GetAllSessions - lookup failed with Decode error: %e", err)
			return nil, err
		}

		results = append(results, session)
	}

	if err := cur.Err(); err != nil {
		log.Printf("GetAllSessions - lookup failed with cursor error: %e", err)
	}

	return results, err
}
