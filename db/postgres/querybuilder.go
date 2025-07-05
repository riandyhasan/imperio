package postgres

import (
	"fmt"
	"strings"

	"github.com/riandyhasan/imperio/model"
)

// GenerateInsert builds an INSERT SQL query
func GenerateInsert(schema model.Schema) (string, []interface{}) {
	var cols []string
	var placeholders []string
	var args []interface{}

	i := 1
	for col, val := range schema.Fields {
		cols = append(cols, col)
		placeholders = append(placeholders, fmt.Sprintf("$%d", i))
		args = append(args, val)
		i++
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		schema.Table,
		strings.Join(cols, ", "),
		strings.Join(placeholders, ", "),
	)

	return query, args
}

// GenerateUpdate builds an UPDATE SQL query (based on `id`)
func GenerateUpdate(schema model.Schema) (string, []interface{}) {
	var sets []string
	var args []interface{}
	var id interface{}
	i := 1

	for col, val := range schema.Fields {
		if strings.ToLower(col) == "id" {
			id = val
			continue
		}
		sets = append(sets, fmt.Sprintf("%s = $%d", col, i))
		args = append(args, val)
		i++
	}

	args = append(args, id) // add ID as last argument for WHERE
	query := fmt.Sprintf(
		"UPDATE %s SET %s WHERE id = $%d",
		schema.Table,
		strings.Join(sets, ", "),
		i,
	)

	return query, args
}

// GenerateDelete builds a DELETE SQL query (based on `id`)
func GenerateDelete(schema model.Schema) (string, []interface{}) {
	id, ok := schema.Fields["id"]
	if !ok {
		return "", nil
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", schema.Table)
	return query, []interface{}{id}
}
