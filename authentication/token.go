package authentication

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

func CreateToken(userID int) (string, error) {
	permission := jwt.MapClaims{}
	permission["authorized"] = true
	permission["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permission["userID"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permission)
	return token.SignedString([]byte("Secret"))
}
