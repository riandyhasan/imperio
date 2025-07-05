package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const (
	yamlExt = ".yaml"
	ymlExt  = ".yml"
	jsonExt = ".json"
)

// Schema holds the structure for a database table.
type Schema struct {
	Table  string            `yaml:"table" json:"table"`
	Fields map[string]string `yaml:"fields" json:"fields"`
}

// FileSchemaLoader is the default implementation of SchemaLoader.
// It supports YAML and JSON based on file extension.
type FileSchemaLoader struct{}

// NewFileSchemaLoader returns a new instance of FileSchemaLoader.
func NewFileSchemaLoader() *FileSchemaLoader {
	return &FileSchemaLoader{}
}

// LoadSchema loads a schema from the provided file path.
func (l *FileSchemaLoader) LoadSchema(path string) (*Schema, error) {
	if path == "" {
		return nil, errors.New("schema file path is empty")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read schema file: %w", err)
	}

	switch ext := filepath.Ext(path); ext {
	case yamlExt, ymlExt:
		return loadYAMLSchema(data)
	case jsonExt:
		return loadJSONSchema(data)
	default:
		return nil, fmt.Errorf("unsupported file extension: %s", ext)
	}
}

func loadYAMLSchema(data []byte) (*Schema, error) {
	var schema Schema
	if err := yaml.Unmarshal(data, &schema); err != nil {
		return nil, fmt.Errorf("invalid YAML schema: %w", err)
	}

	return &schema, nil
}

func loadJSONSchema(data []byte) (*Schema, error) {
	var schema Schema
	if err := json.Unmarshal(data, &schema); err != nil {
		return nil, fmt.Errorf("invalid JSON schema: %w", err)
	}

	return &schema, nil
}
