package dto

import "github.com/golang-jwt/jwt/v4"

var (
	NilIDTokenClaims = IDTokenClaims{}
)

const (
	JwtCtxKey           = "jwt_ctx_key"
	IDTokenCtxKey       = "id_token_ctx_key"
	IDTokenClaimsCtxKey = "id_token_claims_ctx_key"
	UserAccountCtxKey   = "user_account_ctx_key"
)

type JwtTokenContextKey struct{}
type IDTokenContextKey struct{}
type IDTokenClaimsContextKey struct{}
type UserAccountContextKey struct{}

type IDTokenClaims struct {
	jwt.RegisteredClaims
	Azp           string `json:"azp"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	FamilyName    string `json:"family_name"`
	GivenName     string `json:"given_name"`
	Name          string `json:"name"`
	TokenSource   string `json:"token_source"`
}
