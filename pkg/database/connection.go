package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const STR_CONN = `host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`

func OpenConnection() (*sql.DB, error) {

	conn, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			STR_CONN,
			os.Getenv("DATABASE_HOST"),
			os.Getenv("DATABASE_PORT"),
			os.Getenv("DATABASE_USER"),
			os.Getenv("DATABASE_PASSWD"),
			os.Getenv("DATABASE_NAME")),
	)

	if err != nil {
		log.Fatalf("Error to open connection: %v", err)
		return nil, err
	}

	err = conn.Ping()

	if err != nil {
		log.Fatalf("Error to ping database : %v", err)
		conn.Close()
		return nil, err
	}

	return conn, err
}
