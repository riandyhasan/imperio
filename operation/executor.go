package operation

import (
	"errors"
	"fmt"

	"github.com/riandyhasan/imperio/model"
)

// Operation types
const (
	OpWrite  = "write"
	OpUpdate = "update"
	OpDelete = "delete"
)

// Error messages
const (
	ErrUnknownOperation = "unknown operation: %s"
)

// Database interface defines the methods for database operations.
//
//go:generate mockery --name=Database --output=./mocks
type Database interface {
	Write(data model.Schema) error
	Update(data model.Schema) error
	Delete(data model.Schema) error
}

// Executor coordinates execution of database operations using a provided schema.
type Executor struct {
	DB     Database
	Schema *model.Schema
}

// NewExecutor returns a properly constructed Executor instance.
func NewExecutor(db Database, schema *model.Schema) (*Executor, error) {
	if db == nil {
		return nil, errors.New("database instance is required")
	}
	if schema == nil {
		return nil, errors.New("schema definition is required")
	}

	return &Executor{
		DB:     db,
		Schema: schema,
	}, nil
}

// Execute runs the requested operation using the appropriate handler.
func (e *Executor) Execute(op string) error {
	var data model.Schema

	switch op {
	case OpWrite:
		data = generateRandomData(e.Schema)
		return e.DB.Write(data)
	case OpUpdate:
		data = generateUpdateData(e.Schema)
		return e.DB.Update(data)
	case OpDelete:
		data = generateDeleteData(e.Schema)
		return e.DB.Delete(data)
	default:
		return fmt.Errorf(ErrUnknownOperation, op)
	}
}
