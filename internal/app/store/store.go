package store

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Store struct {
	config         *Config
	db             *mongo.Database
	userRepository *UserRepository
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(s.config.MongoDBURL))
	if err != nil {
		return err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	s.db = client.Database("testing2")
	return nil
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
