package token

import (
	"fmt"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type JWTTokenParser struct {
	key string
}

func NewJWTTokenParser(key string) *JWTTokenParser {
	return &JWTTokenParser{key: key}
}

type JWTTokenData struct {
	ID   id.SnowFlakeID
	Type string
}

// Parse トークンからユーザー情報を抜き出す
func (g *JWTTokenParser) Parse(token string) (JWTTokenData, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(g.key), nil
	})
	if err != nil {
		return JWTTokenData{}, fmt.Errorf("failed to parse token: %w", err)
	}

	if !t.Valid {
		return JWTTokenData{}, fmt.Errorf("invalid token: %w", err)
	}

	subject, err := t.Claims.GetSubject()
	if err != nil {
		return JWTTokenData{}, fmt.Errorf("failed to get ClaimSubject: %w", err)
	}

	_, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return JWTTokenData{}, fmt.Errorf("failed to parse token: %w", err)
	}

	return JWTTokenData{ID: id.SnowFlakeID(subject), Type: t.Claims.(jwt.MapClaims)["type"].(string)}, nil
}
