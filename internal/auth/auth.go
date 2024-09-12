package auth

import (
	"crypto/sha256"
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

	tokenSha256Hash := sha256.Sum256([]byte(signedToken))
	bcryptHash, err := bcrypt.GenerateFromPassword(tokenSha256Hash[:], bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	hashStr := string(bcryptHash)

	db.UpdateRefreshToken(uid, hashStr)

	return signedToken, nil

}
