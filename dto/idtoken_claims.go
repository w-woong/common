package dto

import "github.com/golang-jwt/jwt/v4"

var (
	NilIDTokenClaims = IDTokenClaims{}
)

type TokenIdentifierKey struct{}
type IDTokenClaimsKey struct{}
type TokenSourceKey struct{}
type UserAccountKey struct{}

type IDTokenClaims struct {
	jwt.RegisteredClaims
	Azp           string `json:"azp"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	FamilyName    string `json:"family_name"`
	GivenName     string `json:"given_name"`
	Name          string `json:"name"`
}
