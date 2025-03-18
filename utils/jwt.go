package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/robertantonyjaikumar/hangover-common/config"
	"hangover/models"
	"time"
)

var jwtKey = []byte(config.CFG.V.Get("access_secret").(string))

// Define the Claims struct for JWT
type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func (*Claims) Valid() error {
	return nil
}

// Function to generate JWT token
func GenerateJWT(user *models.User) (string, error) {
	claims := &Claims{
		UserID: user.UUID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Minute)), // Set expiration time (1 day)
			Issuer:    config.CFG.V.Get("issuer").(string),
		},
	}

	// Create the token with claims and sign it
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Function to validate JWT token
func ValidateJWT(tokenString string) (*Claims, error) {
	// Parse and validate the token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Make sure the token method is expected
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the key for signing method
		return jwtKey, nil
	})

	// Return the claims if valid
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
