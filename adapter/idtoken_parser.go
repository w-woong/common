package adapter

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/w-woong/common"
	"github.com/w-woong/common/utils"
)

type jwksIDTokenParser struct {
	jwksStore utils.JwksGetter
}

func NewJwksIDTokenParser(jwksStore utils.JwksGetter) *jwksIDTokenParser {
	return &jwksIDTokenParser{
		jwksStore: jwksStore,
	}
}

func (u *jwksIDTokenParser) ParseWithClaims(idToken string, claims jwt.Claims) (*jwt.Token, error) {
	if idToken == "" {
		return nil, common.ErrIDTokenNotFound
	}

	jwksJson, err := u.jwksStore.Get()
	if err != nil {
		return nil, err
	}

	return utils.ParseJWTWithClaimsJwks(idToken, claims, jwksJson)
}

type rs256SignedIDTokenParser struct {
	jwks []byte
}

func NewRS256SignedIDTokenParser(publicKeyFileNames []string, kids []string) (*rs256SignedIDTokenParser, error) {

	jwks, err := utils.RSAPublicKeyToJwks(publicKeyFileNames, kids)
	if err != nil {
		return nil, err
	}

	return &rs256SignedIDTokenParser{
		jwks: jwks,
	}, nil
}

func (u *rs256SignedIDTokenParser) ParseWithClaims(idToken string, claims jwt.Claims) (*jwt.Token, error) {
	return utils.ParseJWTWithClaimsJwks(idToken, claims, u.jwks)
}
