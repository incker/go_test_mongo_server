package server

import (
	"github.com/sirupsen/logrus"
	store2 "go_test_learning/internal/app/store"
)

type APIServer struct {
	logger *logrus.Logger
	Store  *store2.Store
}

func New(config *Config) (*APIServer, error) {
	logger, err := newLogger(config.LogLevel)
	if err != nil {
		return nil, err
	}

	store, err := store2.New(config.MongoDBURL, config.MongoDBName)
	if err != nil {
		return nil, err
	}

	apiServer := APIServer{
		logger: logger,
		Store:  store,
	}
	return &apiServer, nil
}

func newLogger(logLevel string) (*logrus.Logger, error) {
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return nil, err
	}
	logger := logrus.New()
	logger.SetLevel(level)
	return logger, nil
}
