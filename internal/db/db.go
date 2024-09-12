package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

var (
	host, port, user, password, dbname string
	db                                 *pgx.Conn
)

func ConnectDB() (*pgx.Conn, error) {

	host = os.Getenv("DB_HOST")
	port = os.Getenv("DB_PORT")
	user = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname = os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?options", user, password, host, port, dbname)

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	db = conn
	return conn, nil
}
