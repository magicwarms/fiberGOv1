package helpers

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/magicwarms/fiberGOv1/config"
	"golang.org/x/crypto/bcrypt"
)

// JwtClaim adds email as a claim to the token
type JwtClaim struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 13)
	if err != nil {
		panic(err.Error)
	}
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateToken generates a jwt token
func GenerateToken(id, email string) (signedToken string) {
	claims := &JwtClaim{
		ID:    id,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(72)).Unix(),
			Issuer:    "LoginService",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	signedToken, err := token.SignedString([]byte(config.GoDotEnvVariable("SECRET_KEY")))
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}

	return signedToken
}
