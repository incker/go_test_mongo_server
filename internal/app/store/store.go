package store

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Store struct {
	db             *mongo.Database
	userRepository *UserRepository
}

func New(mongoDBURL string, mongoDBName string) (*Store, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoDBURL))
	if err != nil {
		return nil, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	st := &Store{
		db: client.Database(mongoDBName),
	}
	return st, nil
}

func (s *Store) Close() error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return s.db.Client().Disconnect(ctx)
}

func (s *Store) User() *UserRepository {
	if s.userRepository == nil {
		s.userRepository = NewUserRepository(s)
	}
	return s.userRepository
}
