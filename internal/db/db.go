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

func ConnectDB() *pgx.Conn {

	host = os.Getenv("DB_HOST")
	port = os.Getenv("DB_PORT")
	user = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname = os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?options", user, password, host, port, dbname)

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Println(err)
	}
	db = conn
	return conn
}

func VerifyUser(uid string) (bool, error) {
	if db == nil {
		err := fmt.Errorf("no db connection")
		log.Println(err)
		return false, err
	}

	var res map[string]any

	row := db.QueryRow(context.Background(), "SELECT uid FROM users WHERE uid=@uid", pgx.NamedArgs{"uid": uid})
	err := row.Scan(res)
	return err != pgx.ErrNoRows, err

}

func VerifyRefreshToken(uid string, token string) (bool, error) {
	if db == nil {
		err := fmt.Errorf("no db connection")
		log.Println(err)
		return false, err
	}

	var res map[string]any

	is_active := true
	row := db.QueryRow(context.Background(), "SELECT * FROM refresh_tokens WHERE user_id = @uid AND token=@token AND is_active=@is_active",
		pgx.NamedArgs{"uid": uid, "token": token, "is_active": is_active})

	err := row.Scan(res)
	return err != pgx.ErrNoRows, err
}

func UpdateRefreshToken(uid string, token string) error {
	if db == nil {
		err := fmt.Errorf("no db connection")
		log.Println(err)
		return err
	}

	_, err := db.Exec(context.Background(), "UPDATE users SET is_active=0 WHERE uid=@uid", pgx.NamedArgs{"uid": uid})
	if err != nil {
		return err
	}

	is_active := true
	res, err := db.Exec(context.Background(), "INSERT INTO refresh_tokens(user_id, token, is_active) VALUES (@uid, @token, @is_active)",
		pgx.NamedArgs{"uid": uid, "token": token, "is_active": is_active})
	if err != nil {
		return err
	}
	if res.RowsAffected() != 1 {
		err = fmt.Errorf("rows affected incorrectly")
		return err
	}
	return nil
}
