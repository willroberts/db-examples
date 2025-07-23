// cmd/scan/main.go
package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/willroberts/mysqltest"
)

func main() {
	credMgr, err := mysqltest.NewCredentialManager(".credentials")
	if err != nil {
		log.Fatal(err)
	}

	password, err := credMgr.GetPassword("testuser")
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("mysql", mysqltest.DSN{
		Host:     "127.0.0.1",
		Port:     3306,
		Database: "test",
		Username: "testuser",
		Password: password,
	}.ToString())

	db.SetConnMaxLifetime(1 * time.Minute)
	db.SetMaxOpenConns(2)
	db.SetMaxIdleConns(2)

	ctx := context.Background()
	if err := db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}

	rows, err := db.QueryContext(ctx, "select * from mytable")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var id int
		var value int
		err = rows.Scan(&id, &value)
		if err != nil {
			// Will be logged below, after closing the connection.
			break
		}
		log.Println("Value from database:", value)
	}

	if closeErr := rows.Close(); closeErr != nil {
		log.Fatal(closeErr)
	}

	if err != nil {
		// Failed to process a row.
		log.Fatal(err)
	}

	if err := rows.Err(); err != nil {
		// Encountered an error during iteration.
		log.Fatal(err)
	}
}
