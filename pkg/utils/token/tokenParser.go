package token

import (
	"errors"

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
		return JWTTokenData{}, errors.New("failed to parse token")
	}

	if !t.Valid {
		return JWTTokenData{}, errors.New("token is invalid")
	}

	subject, err := t.Claims.GetSubject()
	if err != nil {
		return JWTTokenData{}, err
	}

	_, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return JWTTokenData{}, errors.New("failed to parse token")
	}

	return JWTTokenData{ID: id.SnowFlakeID(subject), Type: t.Claims.(jwt.MapClaims)["type"].(string)}, nil
}
