package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/willroberts/databases"
)

const (
	fileDB     = "db.sql"
	inMemoryDB = ":memory:"
)

func main() {
	// Use SQLite3 as a file-backed SQL database.
	db, err := sql.Open("sqlite3", fileDB) // Use inMemoryDB for in-memory mode.
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test the database with a ping.
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// Perform some queries.
	_, err = db.Exec(databases.CreateTable)
	if err != nil {
		log.Fatal("db.create:", err)
	}

	_, err = db.Exec(databases.Insert, 1, "test record")
	if err != nil {
		log.Fatal("db.insert:", err)
	}

	rows, err := db.Query(databases.Select, nil)
	if err != nil {
		log.Fatal("db.select:", err)
	}
	for rows.Next() {
		var id int
		var value string
		if err := rows.Scan(&id, &value); err != nil {
			log.Fatal("db.select.scan:", err)
		}
		log.Println("db.select.id:", id)
		log.Println("db.select.value:", value)
	}

	// Print connection statistics.
	databases.PrintStats("sqlite3", db.Stats())
}
