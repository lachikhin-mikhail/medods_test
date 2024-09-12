package auth

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/lachikhin-mikhail/medods_test/internal/db"
	"golang.org/x/crypto/bcrypt"
)

var (
	secret []byte
)

func GenerateAccessToken(uid string) (string, error) {
	secret = []byte(os.Getenv("SECRET"))

	const minutesExpire = 1
	claims := jwt.MapClaims{
		"uid": uid,
		"Exp": time.Now().Add(time.Minute * minutesExpire).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	signedToken, err := jwtToken.SignedString(secret)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return signedToken, nil
}

func GenerateRefreshToken(uid string) (string, error) {
	secret = []byte(os.Getenv("SECRET"))

	const minutesExpire = 10
	claims := jwt.MapClaims{
		"uid": uid,
		"Exp": time.Now().Add(time.Minute * minutesExpire).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	signedToken, err := jwtToken.SignedString(secret)
	if err != nil {
		log.Println(err)
		return "", err
	}

	tokenHash, err := bcrypt.GenerateFromPassword([]byte(signedToken), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	tokenHashStr := string(tokenHash)

	db.UpdateRefreshToken(uid, tokenHashStr)

	return signedToken, nil

}
