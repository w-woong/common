package adapter

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/w-woong/common"
	"github.com/w-woong/common/dto"
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

	return utils.ParseJWTWithClaimsJwks(idToken, &dto.IDTokenClaims{}, jwksJson)
}
