package config

import (
	"errors"
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

// Config holds the user-defined settings for running Imperio.
type Config struct {
	Database       string            `yaml:"database"`        // e.g., mysql, postgres
	SchemaFile     string            `yaml:"schema_file"`     // path to schema file
	Operations     []string          `yaml:"operations"`      // e.g., ["write", "update"]
	OpsPerSecond   int               `yaml:"ops_per_second"`  // operations per second
	Concurrency    int               `yaml:"concurrency"`     // number of concurrent workers
	RunnerDuration time.Duration     `yaml:"runner_duration"` // duration to run the simulation, negative means "run forever"
	DBConfig       map[string]string `yaml:"db_config"`       // e.g., host, port, user, password
}

// Error constants for config validation
const (
	ErrEmptyConfigPath      = "config path is required"
	ErrMissingDatabase      = "database type must be specified"
	ErrMissingSchemaFile    = "schema_file must be specified"
	ErrEmptyOperations      = "at least one operation must be specified"
	ErrInvalidOpsPerSecond  = "ops_per_second must be greater than zero"
	ErrInvalidConcurrency   = "concurrency must be greater than zero"
	ErrMissingDBConfig      = "db_config must be provided"
	ErrReadConfigFileFailed = "failed to read config file: %w"
	ErrUnmarshalYAMLFailed  = "failed to unmarshal YAML config: %w"
)

// Load reads a YAML configuration file and returns a validated Config instance.
func Load(path string) (*Config, error) {
	if path == "" {
		return nil, errors.New(ErrEmptyConfigPath)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf(ErrReadConfigFileFailed, err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf(ErrUnmarshalYAMLFailed, err)
	}

	if err := validate(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// validate checks if all required fields in the Config are set correctly.
func validate(cfg *Config) error {
	switch {
	case cfg.Database == "":
		return errors.New(ErrMissingDatabase)
	case cfg.SchemaFile == "":
		return errors.New(ErrMissingSchemaFile)
	case len(cfg.Operations) == 0:
		return errors.New(ErrEmptyOperations)
	case cfg.OpsPerSecond <= 0:
		return errors.New(ErrInvalidOpsPerSecond)
	case cfg.Concurrency <= 0:
		return errors.New(ErrInvalidConcurrency)
	case cfg.DBConfig == nil || len(cfg.DBConfig) == 0:
		return errors.New(ErrMissingDBConfig)
	}
	return nil
}
