package repo

import "go.mongodb.org/mongo-driver/mongo"

type Repo struct {
	mongo *mongo.Database
}

func NewRepo(mongo *mongo.Database) *Repo {
	return &Repo{
		mongo: mongo,
	}
}
