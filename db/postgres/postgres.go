package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/riandyhasan/imperio/model"
)

// Postgres implements db.DatabaseClient interface
type Postgres struct {
	conn *sql.DB
}

// Constants
const (
	DriverName     = "postgres"
	ConnStringTmpl = "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s"
	DefaultSSLMode = "disable"
)

// Error messages
var (
	ErrMissingDBConfig = errors.New("missing required database config fields")
	ErrNilConnection   = errors.New("no active PostgreSQL connection")
)

// RequiredKeys defines keys expected in config map
var RequiredKeys = []string{"host", "port", "user", "password", "dbname"}

// NewPostgres constructs a new Postgres instance
func NewPostgres(cfg map[string]string) (*Postgres, error) {
	for _, key := range RequiredKeys {
		if _, ok := cfg[key]; !ok {
			return nil, fmt.Errorf("%w: %s", ErrMissingDBConfig, key)
		}
	}

	sslMode := cfg["sslmode"]
	if sslMode == "" {
		sslMode = DefaultSSLMode
	}

	connStr := fmt.Sprintf(
		ConnStringTmpl,
		cfg["host"],
		cfg["port"],
		cfg["user"],
		cfg["password"],
		cfg["dbname"],
		sslMode,
	)

	db, err := sql.Open(DriverName, connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open postgres connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping postgres: %w", err)
	}

	return &Postgres{conn: db}, nil
}

// Write inserts a new row into the table
func (p *Postgres) Write(schema model.Schema) error {
	if p.conn == nil {
		return ErrNilConnection
	}

	query, args := GenerateInsert(schema)
	_, err := p.conn.Exec(query, args...)
	return err
}

// Update modifies an existing row in the table
func (p *Postgres) Update(schema model.Schema) error {
	if p.conn == nil {
		return ErrNilConnection
	}

	query, args := GenerateUpdate(schema)
	_, err := p.conn.Exec(query, args...)
	return err
}

// Delete removes a row from the table
func (p *Postgres) Delete(schema model.Schema) error {
	if p.conn == nil {
		return ErrNilConnection
	}

	query, args := GenerateDelete(schema)
	_, err := p.conn.Exec(query, args...)
	return err
}

// Close shuts down the database connection
func (p *Postgres) Close() error {
	if p.conn == nil {
		return nil
	}
	return p.conn.Close()
}
