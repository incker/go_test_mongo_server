package store

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"testing"
	"time"
)

func TestStore(t *testing.T, mongoDBURL string) (*Store, func(...string)) {
	t.Helper()
	config := NewConfig()
	config.MongoDBURL = mongoDBURL
	s := New(config)
	if err := s.Open(); err != nil {
		t.Fatal(err)
	}

	// teardown func
	return s, func(collNames ...string) {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		for _, collName := range collNames {
			_, _ = s.db.Collection(collName).DeleteMany(ctx, bson.M{})
		}
		if err := s.Close(); err != nil {
			log.Fatal(err)
		}
	}
}

func TestDB(t *testing.T, mongoDBURL string) (*Store, func(...string)) {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoDBURL))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	db = client.Database("testing2")
	return nil

	t.Helper()
	config := NewConfig()
	config.MongoDBURL = mongoDBURL
	s := New(config)
	if err := s.Open(); err != nil {
		t.Fatal(err)
	}

	// teardown func
	return s, func(collNames ...string) {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		for _, collName := range collNames {
			_, _ = s.db.Collection(collName).DeleteMany(ctx, bson.M{})
		}
		if err := s.Close(); err != nil {
			log.Fatal(err)
		}
	}
}
