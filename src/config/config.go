package config

import (
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	MongoURI string
	MongoDB  string
}

func GetConfig() *Config {
	return &Config{
		MongoURI: getEnv("MONGO_URI", "mongodb://localhost:27017"),
		MongoDB:  getEnv("MONGO_DB", "blog"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func ConnectDB() (*mongo.Database, error) {
	config := GetConfig()

	clientOptions := options.Client().ApplyURI(config.MongoURI)
	client, err := mongo.Connect(nil, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}
	log.Printf("Successfully connected to MongoDB at %s", config.MongoURI)
	db := client.Database(config.MongoDB)
	return db, nil
}
