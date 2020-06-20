package store

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"testing"
	"time"
)

func TestDB(t *testing.T, mongoDBURL string) (*Store, func(...string)) {
	t.Helper()
	store, err := New(mongoDBURL, "testing2")
	if err != nil {
		log.Fatal(err)
	}
	// teardown func
	return store, func(collNames ...string) {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		for _, collName := range collNames {
			_, _ = store.db.Collection(collName).DeleteMany(ctx, bson.M{})
		}
		if err := store.Close(); err != nil {
			log.Fatal(err)
		}
	}
}
