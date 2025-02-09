package server

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Session struct {
	UserId   string `bson:"user_id"`
	Username string `bson:"username"`
}

// TODO: create TTL index
type SessionsStore struct {
	coll *mongo.Collection
}

func (s *SessionsStore) Upsert(userId, username string) error {
	session := Session{
		UserId:   userId,
		Username: username,
	}
	filter := bson.D{{"user_id", userId}}
	update := bson.D{{"$set", session}}

	opts := options.UpdateOne().SetUpsert(true)
	op, err := s.coll.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		return err
	}
	fmt.Println("Upserted", op.UpsertedID)

	return nil
}

func NewSessionsStore(mongo *MongoClient) *SessionsStore {
	coll := mongo.Db.Collection("sessions")

	return &SessionsStore{
		coll: coll,
	}
}
