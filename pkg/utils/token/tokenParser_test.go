package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJWTTokenParser_Parse(t *testing.T) {
	key := SecureRandom(32)
	g := NewJWTTokenGenerator(key)
	p := NewJWTTokenParser(key)

	tok, err := g.NewAccessToken("112233")
	if err != nil {
		t.Fail()
	}

	parsed, err := p.Parse(tok)
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, JWTTokenData{ID: "112233", Type: "access"}, parsed)
}
