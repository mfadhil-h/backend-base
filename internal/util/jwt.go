package util

import (
	"crypto/rsa"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

func LoadKeys() error {
	privPath := viper.GetString("JWT_PRIVATE_KEY_PATH")
	pubPath := viper.GetString("JWT_PUBLIC_KEY_PATH")

	privBytes, err := os.ReadFile(privPath)
	if err != nil {
		return err
	}
	pubBytes, err := os.ReadFile(pubPath)
	if err != nil {
		return err
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privBytes)
	if err != nil {
		return err
	}
	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(pubBytes)
	if err != nil {
		return err
	}

	return nil
}

func GenerateJWT(userID uint, email string) (string, error) {
	expiration := time.Now().Add(time.Hour * time.Duration(viper.GetInt("JWT_EXPIRE_HOURS")))
	claims := jwt.MapClaims{
		"sub": userID,
		"email": email,
		"exp": expiration.Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(privateKey)
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("invalid signing method")
		}
		return publicKey, nil
	})
}
