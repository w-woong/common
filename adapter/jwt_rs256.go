package adapter

import (
	"crypto/rsa"

	"github.com/golang-jwt/jwt/v4"
	"github.com/w-woong/common/port"
	"github.com/w-woong/common/utils"
)

var _ port.JwtRepo = (*RS256SignedJWT)(nil)

type RS256SignedJWT struct {
	kid        string
	privateKey *rsa.PrivateKey
	jwks       []byte
}

func NewRS256SignedJWT(currentKid string, privateKeyFileName string,
	publicKeyFileNames []string, kids []string) (*RS256SignedJWT, error) {

	privateKey, err := utils.LoadRSAPrivateKey(privateKeyFileName)
	if err != nil {
		return nil, err
	}

	jwks, err := utils.RSAPublicKeyToJwks(publicKeyFileNames, kids)
	if err != nil {
		return nil, err
	}

	return &RS256SignedJWT{
		kid:        currentKid,
		privateKey: privateKey,
		jwks:       jwks,
	}, nil
}

func (u *RS256SignedJWT) GenerateToken(claims jwt.Claims) (string, error) {
	return utils.GenerateRS256SignedJWT(u.kid, u.privateKey, claims)
}
func (u *RS256SignedJWT) ParseWithClaims(token string, claims jwt.Claims) (*jwt.Token, error) {
	return utils.ParseJWTWithClaimsJwks(token, claims, u.jwks)
}
