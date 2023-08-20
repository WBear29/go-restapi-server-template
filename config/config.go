package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App
		Log
		Server
		DB
		Tracing
	}

	App struct {
		Name string `yaml:"name" env:"APP_NAME"`
		Env  string `yaml:"env" env:"APP_ENV" env-default:"development"`
	}

	Server struct {
		Port    int    `yaml:"port" env:"SERVER_PORT" env-default:"3000"`
		GinMode string `yaml:"gin_mode" env:"GIN_MODE" env-default:"release"`
	}

	Log struct {
		Level string `yaml:"level" env:"LOG_LEVEL" env-default:"debug"`
	}

	DB struct {
		Host         string `yaml:"host" env:"DB_HOST" env-default:"0.0.0.0"`
		Port         string `yaml:"port" env:"DB_PORT" env-default:"3306"`
		User         string `yaml:"user" env:"DB_USER" env-default:"root"`
		Pass         string `yaml:"pass" env:"DB_PASS" env-default:"password"`
		Name         string `yaml:"name" env:"DB_NAME" env-default:"sample"`
		Retry        uint   `yaml:"retry" env:"DB_RETRY" env-default:"5"`
		MigrationDir string `yaml:"migration_dir" env:"DB_MIGRATION_DIR" env-default:"./database/migrations"`
	}

	Tracing struct {
		Endpoint string `yaml:"endpoint" env:"TRACING_ENDPOINT" env-default:"http://localhost:14268/api/traces"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	if err := cleanenv.ReadConfig("./config/config.yml", cfg); err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, fmt.Errorf("failed to read env: %w", err)
	}

	return cfg, nil
}
