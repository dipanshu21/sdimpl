package lib

import (
	"database/sql"
	"fmt"
)

func ExecuteMySqlQueryWithNewConnection(query string, args ...interface{}) (sql.Result, error) {
	// Create a new connection to the database
	db, err := CreateNewConnectionToDB()
	if err != nil {
		return nil, fmt.Errorf("error creating new connection to database: %w", err)
	}
	defer db.Close()

	// Execute the query
	result, err := db.Exec(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}

	return result, nil
}

func ExecuteMySqlQuery(db *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {
	// Execute the query
	result, err := db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}

	return result, nil
}
