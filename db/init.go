package db

import (
	"context"
	"final_project/config"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Client ... instance of mongo
var Client *mongo.Client

const (
	// DbName ... Database's name
	DbName = "golang"
	// ColName ... Collection's name
	ColName = "student"
)

func init() {
	connect()
}

func connect() {
	fmt.Println("start: ", config.Config.MongoDB.Uri)
	client, err := mongo.NewClient(options.Client().ApplyURI(config.Config.MongoDB.Uri))
	if err != nil {
		log.Fatalf("connect error: %v", err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("error connect to DB: %v\n", err)
	}
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("error ping to DB server: %v \n", err)
	}
	Client = client
}

func InsertNumber() (interface{}, error) {
	fmt.Println("Client: ", Client)
	collection := Client.Database("test1").Collection("numbers")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	// res, _ := collection.InsertOne(ctx, bson.M{"name", "Phat", "age": 25})
	res, _ := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14})
	// if err != nil {
	// 	log.Fatalf("error to test insert: %v", err)
	// 	return nil, err
	// }
	id := res.InsertedID
	fmt.Println("insert")
	return id, nil
}
