package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	privateKey *ecdsa.PrivateKey
	publicKey *ecdsa.PublicKey
)

// It generates the key pair at once.
func init() {
	var err error
	privateKey, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	publicKey = &privateKey.PublicKey
}

// Method for creating a JWT token.
func createToken(username string) (string, error) {

	// We create the new JWT Token using this method.
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"username": username,                              // Token is associated with which username.
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Expiration Time.
	})

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		fmt.Println("error:", err)
		return "", err
	}

	return tokenString, nil
}

// Method for verifying the token.
func verifyToken(tokenString string) error {
	// Here we use the jwt.Parse function for parsing and verifying the token.
	// In the second parameter we use a callback function to reterive the secret key used for signing the token.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})

	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	if !token.Valid {
		return fmt.Errorf("ERROR: Invalid token")
	}

	fmt.Println("Token is valid")
	return nil
}
