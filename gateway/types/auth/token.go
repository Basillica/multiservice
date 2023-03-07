package auth

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Token struct {
	Token string
}

// GenerateToken generates a new jwt token for a given user
//
// It takes a userid and returns a jwt claim object
func (t *Token) GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(24)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("CRYPTO_SECRET")))
}

// ValidateToken validates a given token
//
// It takes a token and returns an error if token is not valid
func (t *Token) ValidateToken() error {
	_, err := jwt.Parse(t.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("CRYPTO_SECRET")), nil
	})
	if err != nil {
		return err
	}
	return nil
}

// ExtractTokenID extracts a token from an IdToken
//
// # It takes a token and returns a tuple of a userid string  and error if token is not valid
//
// # The userid string is extrated from the token if it is valid
func (t *Token) ExtractTokenID() (string, error) {
	token, err := jwt.Parse(t.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("CRYPTO_SECRET")), nil
	})
	if err != nil {
		return "", err
	}
	var uID string
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uID, ok = claims["user_id"].(string)
		if !ok {
			return "", err
		}
	}
	return uID, nil
}
