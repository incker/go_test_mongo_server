package store_test

import (
	"os"
	"testing"
)

var (
	mongoDBURL string
)

func TestMain(m *testing.M) {
	mongoDBURL = os.Getenv("MONGODB_URL")
	if mongoDBURL == "" {
		mongoDBURL = "mongodb://localhost:27017"
	}

	os.Exit(m.Run())
}
