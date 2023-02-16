package port

import "github.com/golang-jwt/jwt/v4"

type JwtRepo interface {
	Jwks() []byte
	GenerateToken(claims jwt.Claims) (string, error)
	ParseWithClaims(token string, claims jwt.Claims) (*jwt.Token, error)
}
