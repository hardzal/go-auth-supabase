package configs

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type ConfigDB struct {
	DBHost     string `env:"SUPABASE_URL"`
	DBPort     string `env:"SUPABASE_PORT"`
	DBUser     string `env:"SUPABASE_USER"`
	DBName     string `env:"SUPABASE_NAME"`
	DBPassword string `env:"SUPABASE_PASSWORD"`
}

func LoadEnv() (*ConfigDB, error) {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}

	cfg := &ConfigDB{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	log.Println(cfg)

	log.Println("Config loaded successfully")
	return cfg, nil
}
