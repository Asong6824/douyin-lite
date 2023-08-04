package service

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"fmt"
)

func (svc *Service) GenerateToken(userID uint32) (string, error) {
	// Define the expiration time for the token.
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create the claims for the token, including the userID.
	claims := jwt.MapClaims{
		"id":  userID,
		"exp": expirationTime.Unix(),
	}

	// Create the JWT token with the claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with a secret key and get the complete token string.
	tokenString, err := token.SignedString([]byte("secret_key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (svc *Service) GetUserIDFromToken(tokenString string) (uint32, error) {
	// Parse the JWT token string using the secret key.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verify the signing method used in the token.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret_key"), nil
	})
	if err != nil {
		return 0, err
	}

	// Extract the claims from the token.
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	// Extract the user ID from the claims.
	userID, ok := claims["id"].(uint32)
	if !ok {
		return 0, fmt.Errorf("invalid user ID")
	}

	return userID, nil
}