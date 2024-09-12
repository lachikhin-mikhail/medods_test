package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func VerifyUser(uid string) (bool, error) {
	if db == nil {
		err := fmt.Errorf("no db connection")
		log.Println(err)
		return false, err
	}

	var res any

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

	var res any

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
