package models

import (
	"database/sql"
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

	log.Print("\n\nConnected to the database successfully!!\n\n")
}