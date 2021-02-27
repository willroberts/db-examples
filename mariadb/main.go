package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/willroberts/databases"
)

func init() {
	// Perform first-time setup.
	db, err := sql.Open("mysql", "root:secret@/")
	if err != nil {
		log.Fatal("sql.open:", err)
	}
	defer db.Close()
	_, err = db.Exec(databases.CreateDatabase)
	if err != nil {
		log.Fatal("db.create_database:", err)
	}
}

func main() {
	db, err := sql.Open("mysql", "root:secret@/test_database")
	if err != nil {
		log.Fatal("sql.open:", err)
	}
	defer db.Close()

	// Test the database with a ping.
	if err := db.Ping(); err != nil {
		log.Fatal("db.ping:", err)
	}

	// Perform some queries.
	_, err = db.Exec(databases.CreateTable)
	if err != nil {
		log.Fatal("db.create_table:", err)
	}

	_, err = db.Exec(databases.Insert, 1, "test record")
	if err != nil {
		log.Fatal("db.insert:", err)
	}

	r, err := db.Query(databases.Select)
	if err != nil {
		log.Fatal("db.select:", err)
	}
	for r.Next() {
		var id int
		var value string
		if err := r.Scan(&id, &value); err != nil {
			log.Fatal("db.select.scan:", err)
		}
		log.Println("db.select.id:", id)
		log.Println("db.select.value:", value)
	}

	// Print connection statistics.
	databases.PrintStats("mariadb", db.Stats())
}
