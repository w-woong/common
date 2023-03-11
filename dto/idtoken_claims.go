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
	Azp           string `json:"azp,omitempty"`
	Email         string `json:"email,omitempty"`
	EmailVerified bool   `json:"email_verified,omitempty"`
	FamilyName    string `json:"family_name,omitempty"`
	GivenName     string `json:"given_name,omitempty"`
	Name          string `json:"name,omitempty"`
	TokenSource   string `json:"token_source,omitempty"`
	SourceIssuer  string `json:"source_issuer,omitempty"`

	UserID      *uint64 `json:"user_id,omitempty"`
	PhoneNumber string  `json:"phone,omitempty"`
	Gender      string  `json:"gender,omitempty"`
}
