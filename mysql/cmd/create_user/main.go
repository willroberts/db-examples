// cmd/create_user/main.go
package main

import (
	"context"
	"database/sql"
	"fmt"
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

	rootPassword, err := credMgr.GetPassword("root")
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("mysql", mysqltest.DSN{
		Host:     "127.0.0.1",
		Port:     3306,
		Database: "",
		Username: "root",
		Password: rootPassword,
	}.ToString())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database.")

	db.SetConnMaxLifetime(1 * time.Minute)
	db.SetMaxOpenConns(2)
	db.SetMaxIdleConns(2)

	ctx := context.Background()
	dropQuery := "DROP USER IF EXISTS 'testuser'@'%'"
	if _, err := db.ExecContext(ctx, dropQuery); err != nil {
		log.Fatal(err)
	}
	log.Println("Dropped any existing 'testuser'@'%'")

	userPassword, err := credMgr.GetPassword("testuser")
	if err != nil {
		log.Fatal(err)
	}
	createQuery := fmt.Sprintf("CREATE USER 'testuser'@'%%' IDENTIFIED BY '%s'", userPassword)
	if _, err := db.ExecContext(ctx, createQuery); err != nil {
		log.Fatal(err)
	}
	log.Println("Created user")

	grantQuery := "GRANT SELECT, INSERT, UPDATE, DELETE ON test.* TO 'testuser'@'%'"
	if _, err := db.ExecContext(ctx, grantQuery); err != nil {
		log.Fatal(err)
	}
	log.Println("Assigned privileges")
}
