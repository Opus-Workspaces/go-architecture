package mongo

import (
	"context"
	"fmt"
	config "go-architecture/app/config/database"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

type MongoDB struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func ConnectMongo(dbCfg config.DatabaseType) *MongoDB {
	mongoCfg := dbCfg.MongoDB

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoCfg.Timeout)*time.Second)

	defer cancel()

	enableParams := "?retryWrites=true&retryReads=true&w=majority&authSource=admin"

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", mongoCfg.User, mongoCfg.Password, mongoCfg.Host, mongoCfg.Port, enableParams)

	if mongoCfg.User == "" && mongoCfg.Password == "" {
		// uri = "mongodb://localhost:27017"
		uri = fmt.Sprintf("mongodb://%s:%d/%s", mongoCfg.Host, mongoCfg.Port, enableParams)
	}

	clientOpts := options.Client().ApplyURI(uri)

	log.Println("Connecting to MongoDB...")
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		panic("DB.Mongo.Connect: " + err.Error())
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic("DB.Mongo.Ping: " + err.Error())
	}

	db := client.Database(mongoCfg.DatabaseName)

	log.Println("Connected to MongoDB!")
	return &MongoDB{
		Client: client,
		DB:     db,
	}
}

func DisconnectMongo(c *mongo.Client) error {
	return c.Disconnect(context.Background())
}
