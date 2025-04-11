package database

import (
	"database/sql"
	"log"
)

func TableExists(db *sql.DB, dbName string, tableName string) bool {
	// Here we have to define a count variable via we get that the table exists or not.
	var count int
	query := "SELECT COUNT(*) FROM INFORMATION_SCHEMA.tables WHERE table_schema = ? AND table_name = ?" // In the table schema put the db name and in the table name put the table name.
	err := db.QueryRow(query, dbName, tableName).Scan(&count)
	if err != nil {
		log.Printf("ERROR: checking table: %v", err)
		return false
	}

	return count > 0
}
