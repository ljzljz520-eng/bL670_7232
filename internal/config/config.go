package config

import "os"

type Config struct {
	DatabasePath string
	Address      string
}

func Default() Config { return Config{DatabasePath: "training-review.db", Address: ":8080"} }

func FromEnv() Config {
	cfg := Default()
	if value := os.Getenv("TRAINING_REVIEW_DB"); value != "" {
		cfg.DatabasePath = value
	}
	if value := os.Getenv("TRAINING_REVIEW_ADDR"); value != "" {
		cfg.Address = value
	}
	return cfg
}
