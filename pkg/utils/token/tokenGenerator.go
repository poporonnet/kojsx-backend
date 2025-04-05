package token

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/poporonnet/kojsx-backend/pkg/utils/id"
)

type JWTTokenGenerator struct {
	key string
}

func NewJWTTokenGenerator(key string) *JWTTokenGenerator {
	return &JWTTokenGenerator{key: key}
}

func (g *JWTTokenGenerator) NewAccessToken(uid id.SnowFlakeID) (string, error) {
	c := jwt.MapClaims{
		"sub":  string(uid),
		"exp":  jwt.NewNumericDate(time.Now().Add(time.Minute * 10)),
		"iat":  jwt.NewNumericDate(time.Now()),
		"type": "access",
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	res, err := t.SignedString([]byte(g.key))
	if err != nil {
		return "", fmt.Errorf("TokenGenerator: %w", err)
	}

	return res, nil
}

func (g *JWTTokenGenerator) NewRefreshToken(uid id.SnowFlakeID) (string, error) {
	c := jwt.MapClaims{
		"sub":  string(uid),
		"exp":  jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),
		"iat":  jwt.NewNumericDate(time.Now()),
		"type": "refresh",
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	res, err := t.SignedString([]byte(g.key))
	if err != nil {
		return "", fmt.Errorf("TokenGenerator: %w", err)
	}
	return res, nil
}

func (g *JWTTokenGenerator) NewVerifyToken(uid id.SnowFlakeID) (string, error) {
	c := jwt.MapClaims{
		"sub":  string(uid),
		"exp":  jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
		"iat":  jwt.NewNumericDate(time.Now()),
		"type": "verify",
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	res, err := t.SignedString([]byte(g.key))
	if err != nil {
		return "", fmt.Errorf("TokenGenerator: %w", err)
	}
	return res, nil
}
