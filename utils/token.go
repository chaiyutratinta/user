package utils

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
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
	key, err := ioutil.ReadFile("./signature")

	if err != nil {
		log.Panic(err)
		token = ""

		return
	}

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

func Hash(password string) (hashed string, err error) {
	mac := md5.New()
	mac.Write([]byte(password))
	hashed = fmt.Sprintf("%x", mac.Sum(nil))

	return
}
