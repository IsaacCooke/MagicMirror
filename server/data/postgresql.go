package data

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func getConnectionString() string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	return psqlInfo
}

func Connect() *sql.DB {
	psqlInfo := getConnectionString()

	db, err := sql.Open("postgres", psqlInfo)
	checkError(err)

	err = db.Ping()
	checkError(err)

	fmt.Println("Successfully connected!")

	return db
}

func SetupDB() {
	db := Connect()

	_, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS reminders (
    				id SERIAL PRIMARY KEY,
    				title TEXT,
    				description TEXT,
    				is_done boolean,
    				due_date TIMESTAMP,
					repeat boolean
				);
				CREATE TABLE IF NOT EXISTS flashcards (
    				id SERIAL PRIMARY KEY,
    				term TEXT,
					definition TEXT
				);`)
	checkError(err)
	fmt.Println("Successfully created tables!")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
