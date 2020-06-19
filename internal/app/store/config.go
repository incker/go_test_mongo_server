package store

type Config struct {
	MongoDBURL string `toml:"mongodb_url"`
}

func NewConfig() *Config {
	return &Config{
		MongoDBURL: "mongodb://localhost:27017",
	}
}
