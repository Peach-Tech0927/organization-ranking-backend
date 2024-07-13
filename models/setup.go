package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/briandowns/spinner"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func ConnectDatabase() {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Prefix = "Connecting to the database "
	s.Start()
	defer s.Stop()
	time.Sleep(time.Second)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DSN")
	if len(dsn) == 0 {
		log.Fatal("DSN is not set in the environment")
	}

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = executeSQLFile(DB, "migrations/user.sql")
	if err != nil {
		log.Fatal("Error creating users table, ", err)
	}

	log.Print("\n\nConnected to the database successfully!!\n\n")
}

func executeSQLFile(db *sql.DB, filePath string) error {
	query, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("could not read SQL file: %v", err)
	}

	_, err = db.Exec(string(query))
	if err != nil {
		return fmt.Errorf("could not execute SQL file: %v", err)
	}

	return nil
}