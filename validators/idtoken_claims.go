package validators

import "github.com/golang-jwt/jwt/v4"

var (
	NilIDTokenClaims = IDTokenClaims{}
)

type IDTokenClaimsKey struct{}

type IDTokenClaims struct {
	jwt.RegisteredClaims
	Azp           string `json:"azp"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	FamilyName    string `json:"family_name"`
	GivenName     string `json:"given_name"`
	Name          string `json:"name"`
}
