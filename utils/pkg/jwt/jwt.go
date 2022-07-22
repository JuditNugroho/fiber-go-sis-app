package jwt

import (
	"crypto/rand"
	"crypto/rsa"
	"time"

	"github.com/golang-jwt/jwt/v4"

	formsEntity "github.com/fiber-go-sis-app/internal/entity/forms"
)

var privateKey *rsa.PrivateKey

func GenerateJWT() error {
	var err error
	privateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	return err
}

func GetPrivateKey() *rsa.PrivateKey {
	return privateKey
}

func CreateJWTToken(req formsEntity.JWTRequest) (string, error) {
	// Create the Claims
	claims := jwt.MapClaims{
		"user_id": req.UserID,
		"name":    req.Name,
		"admin":   req.Admin,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString(privateKey)
	if err != nil {
		return t, err
	}
	return t, nil
}
