package utils

import (
	"time"

	"fmt"

	"github.com/dgrijalva/jwt-go"
)

// SignTestCompletion signs that a user has finished a test at a given point in time
func SignTestCompletion(userID string, currentTime time.Time) (string, error) {
	// Create a new JWT token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the claims
	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = userID
	claims["currentTime"] = currentTime.Unix()

	// Sign the token with a secret key
	tokenString, err := token.SignedString([]byte("secretKey"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyTestCompletion verifies if the tokenString belongs to the specified user
func VerifyTestCompletion(tokenString string, userID string) (bool, error) {
	// Parse the token

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secretKey"), nil
	})

	// Check for parsing errors
	if err != nil {
		return false, err
	}

	// Verify the claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return false, nil
	}

	// Check if the userID matches
	if claims["userID"] != userID {
		return false, nil
	}

	return true, nil
}
