package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
	"math/rand"
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

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPassword, dbHost, dbName)

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = executeSQLFile(DB, "migrations/user.sql")
	if err != nil {
		log.Fatal("Error creating users table, ", err)
	}
	err = executeSQLFile(DB, "migrations/organization.sql")
	if err != nil {
		log.Fatal("Error creating organizations table, ", err)
	}
	err = executeSQLFile(DB, "migrations/user-organization.sql")
	if err != nil {
		log.Fatal("Error creating user_organiation_membership table, ", err)
	}
	//mocデータの挿入
	insertMockData(DB)

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


func insertMockData(db *sql.DB) {
    rand.Seed(time.Now().UnixNano())

    for i := 1; i <= 10; i++ {
        name := fmt.Sprintf("User%d", i)
		id := fmt.Sprintf("id%d", i)
		email := fmt.Sprintf("email%d", i)
        score := rand.Intn(100)
        _, err := db.Exec("INSERT INTO users (email,username,password,github_id, contributions) VALUES (?,?,?,?,?)",email, name,"mocmoc",id, score)
        if err != nil {
            log.Fatal(err)
        }
    }

    for i := 1; i <= 10; i++ {
        name := fmt.Sprintf("Organization%d", i)
        _, err := db.Exec("INSERT INTO organizations (name) VALUES (?)", name)
        if err != nil {
            log.Fatal(err)
        }
    }

    for i := 1; i <= 10; i++ {
        userID := rand.Intn(10) + 1
        organizationID := rand.Intn(10) + 1
        _, err := db.Exec("INSERT INTO user_organiation_membership (user_id, organization_id) VALUES (?, ?)", userID, organizationID)
        if err != nil {
            log.Fatal(err)
        }
    }
}