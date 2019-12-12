package utils

import (
	"crypto/md5"
	"fmt"
	"log"
	"time"
	"user/models"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func Token(userID, scope string) (token models.Token, err error) {
	key := []byte("a40449d2a209f1ba700c20da616b01a2f360b39f97152aa384e01f54ecab17571c5311e5f83108bc57fc94ddcc2ba12530edc2db5f6a57458c8d330d6317307e")

	duration, _ := time.ParseDuration("1h")
	expDate := time.Now().Add(duration)
	expAt := time.Until(expDate)
	claims := &Claims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: expAt.Milliseconds(),
			Subject:   scope,
		},
	}
	unsigned := jwt.NewWithClaims(jwt.GetSigningMethod("HS512"), claims)
	signedStr, err := unsigned.SignedString(key)

	if err != nil {
		token = models.Token("")
		log.Println(err)

		return
	}

	token = models.Token(signedStr)

	return
}

func Hash(password string) (hashed string) {
	mac := md5.New()
	mac.Write([]byte(password))
	hashed = fmt.Sprintf("%x", mac.Sum(nil))

	return
}
