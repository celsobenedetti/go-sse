package server

import (
	"context"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Session struct {
	UserId    string    `bson:"user_id,omitempty"`
	Username  string    `bson:"username,omitempty"`
	UpdatedAt time.Time `bson:"updated_at,omitempty"`
}

type SessionsStore struct {
	coll *mongo.Collection
}

func (s *SessionsStore) Upsert(userId, username string) error {
	session := Session{
		UserId:    userId,
		Username:  username,
		UpdatedAt: time.Now(),
	}

	filter := bson.D{{Key: "user_id", Value: userId}}
	update := bson.D{{Key: "$set", Value: session}}
	opts := options.UpdateOne().SetUpsert(true)

	result, err := s.coll.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		return err
	}
	slog.Debug("updated session doc", "session", session, "UpsertedCount", result.UpsertedCount)

	return nil
}

func NewSessionsStore(client *MongoClient) *SessionsStore {
	coll := client.Db.Collection("sessions")

	err := createTTLIndex(coll)
	if err != nil {
		slog.Warn("failed to create TTL index", "err", err)
	}

	return &SessionsStore{
		coll: coll,
	}
}

func createTTLIndex(coll *mongo.Collection) error {
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "updated_at", Value: 1}},
		Options: options.Index().SetName("session_ttl").SetExpireAfterSeconds(10), // TODO: CE-32 get from env
	}

	_, err := coll.Indexes().CreateOne(context.TODO(), indexModel)
	return err
}
