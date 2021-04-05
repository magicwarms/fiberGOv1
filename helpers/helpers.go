package helpers

import (
	"errors"
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

// JwtWrapper wraps the signing key and the issuer
type JwtWrapper struct {
	SecretKey string
	Issuer    string
	ExpiresAt int64
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
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(72)).Unix(), // expires at 3 days
			Issuer:    "AuthService",
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

//ValidateToken validates the jwt token
func (j *JwtWrapper) ValidateToken(signedToken string) (claims *JwtClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.SecretKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JwtClaim)
	if !ok {
		err = errors.New("ERROR-PARSE-TOKEN")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("TOKEN-EXPIRED")
		return
	}

	return
}
