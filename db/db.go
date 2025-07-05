package db

import (
	"errors"
	"fmt"

	"github.com/riandyhasan/imperio/db/postgres"
	"github.com/riandyhasan/imperio/model"
)

// DBMS represents the supported database types.
type DBMS string

// String returns the string representation of the DBMS type.
func (d DBMS) String() string {
	return string(d)
}

// Supported database constants.
const (
	MYSQL    DBMS = "mysql"
	POSTGRES DBMS = "postgres"
)

// Error messages.
const (
	ErrUnsupportedDBType = "unsupported database type: %s"
)

// Predefined errors.
var (
	ErrDatabaseNotInitialized = errors.New("database client not initialized")
)

// DatabaseClient defines the interface for all database operations.
//
//go:generate mockery --name=DatabaseClient --output=./mocks
type DatabaseClient interface {
	Write(data model.Schema) error
	Update(data model.Schema) error
	Delete(data model.Schema) error
	Close() error
}

// Database wraps a selected database strategy that implements DatabaseClient.
type Database struct {
	client DatabaseClient
}

// NewDatabase returns a Database wrapper with the correct strategy selected.
func NewDatabase(cfg map[string]string, dbType DBMS) (*Database, error) {
	var client DatabaseClient
	var err error

	switch dbType {
	case POSTGRES:
		client, err = postgres.NewPostgres(cfg)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf(ErrUnsupportedDBType, dbType)
	}

	return &Database{client: client}, nil
}

// Write delegates a write operation to the database client.
func (d *Database) Write(data model.Schema) error {
	if d.client == nil {
		return ErrDatabaseNotInitialized
	}
	return d.client.Write(data)
}

// Update delegates an update operation to the database client.
func (d *Database) Update(data model.Schema) error {
	if d.client == nil {
		return ErrDatabaseNotInitialized
	}
	return d.client.Update(data)
}

// Delete delegates a delete operation to the database client.
func (d *Database) Delete(data model.Schema) error {
	if d.client == nil {
		return ErrDatabaseNotInitialized
	}
	return d.client.Delete(data)
}

// Close shuts down the database connection.
func (d *Database) Close() error {
	if d.client == nil {
		return ErrDatabaseNotInitialized
	}
	return d.client.Close()
}
