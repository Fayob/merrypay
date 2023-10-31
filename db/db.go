package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func openDB(DNS string) (*sql.DB, error) { 
	db, err := sql.Open("postgres", DNS)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Database Disconnected")
		return nil, err
	}

	fmt.Println("Database connected successfully...")
	return db, nil
}

func ConnectToDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading env file")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	
	// connStr := "postgres://postgres:password@localhost/merrypay?sslmode=disable"
	DNS := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)
	conn, err := openDB(DNS)
	if err != nil {
		log.Fatal(err)
	}

	return conn
}