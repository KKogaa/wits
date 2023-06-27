package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_NAME                   string
	JWT_SECRET                string
	PORT                      string
	VECTORIZEME_SERVICE       string
	MINIO_ENDPOINT            string
	MINIO_ACCESS_KEY_ID       string
	MINIO_SECRET_KEY          string
	MINIO_BUCKET_NAME         string
	MINIO_LOCATION            string
	ELASTIC_SEARCH_ENDPOINT   string
	ELASTIC_SEARCH_INDEX_NAME string
	QDRANT_COLLECTION_NAME    string
	QDRANT_ENDPOINT           string
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		panic("failed to load .env file")
	}

	config := Config{
		DB_NAME:                   os.Getenv("DB_NAME"),
		JWT_SECRET:                os.Getenv("JWT_SECRET"),
		PORT:                      os.Getenv("PORT"),
		VECTORIZEME_SERVICE:       os.Getenv("VECTORIZEME_SERVICE"),
		MINIO_ENDPOINT:            os.Getenv("MINIO_ENDPOINT"),
		MINIO_ACCESS_KEY_ID:       os.Getenv("MINIO_ACCESS_KEY_ID"),
		MINIO_SECRET_KEY:          os.Getenv("MINIO_SECRET_KEY"),
		MINIO_BUCKET_NAME:         os.Getenv("MINIO_BUCKET_NAME"),
		MINIO_LOCATION:            os.Getenv("MINIO_LOCATION"),
		ELASTIC_SEARCH_ENDPOINT:   os.Getenv("ELASTIC_SEARCH_ENDPOINT"),
		ELASTIC_SEARCH_INDEX_NAME: os.Getenv("ELASTIC_SEARCH_INDEX_NAME"),
		QDRANT_COLLECTION_NAME:    os.Getenv("QDRANT_COLLECTION_NAME"),
		QDRANT_ENDPOINT:           os.Getenv("QDRANT_ENDPOINT"),
	}

	return &config
}
