package initializers

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectToDatabase() {
	var err error
	DB, err = sql.Open("postgres", os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	fmt.Println("Connected to database")
}
