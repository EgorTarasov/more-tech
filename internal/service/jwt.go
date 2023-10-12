package service

import (
	"errors"
	"more-tech/internal/config"
	"more-tech/internal/model"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type jwtClaims struct {
	Email         string
	RegisteredClaims jwt.RegisteredClaims
}

func CreateAccessToken(email string) (string, error) {
	key := []byte(config.Cfg.SecretKey)

	jwtClaims := jwtClaims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			Subject:   email,
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims.RegisteredClaims)
	return accessToken.SignedString(key)
}

func VerifyAccessToken(accessTokenString string) (*jwt.Token, error) {
	accessToken, err := jwt.Parse(accessTokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(config.Cfg.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func VerifyPassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func AuthenticateUser(userData model.UserLoginRequest, hashedPassword string) (model.AuthResponse, error) {
	if err := VerifyPassword(hashedPassword, userData.Password); err != nil {
		return model.AuthResponse{}, err
	}

	accessToken, err := CreateAccessToken(userData.Email)
	if err != nil {
		return model.AuthResponse{}, err
	}

	return model.AuthResponse{
		AccessToken: accessToken,
		Type:        "bearer",
	}, nil
}