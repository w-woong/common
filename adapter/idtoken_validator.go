package adapter

import (
	"errors"

	"github.com/MicahParks/keyfunc"
	"github.com/golang-jwt/jwt/v4"
	"github.com/w-woong/common/dto"
	"github.com/w-woong/common/utils"
)

type jwksIDTokenValidator struct {
	jwksStore          utils.JwksGetter
	tokenSourceKey     string
	tokenIdentifierKey string
	idTokenKey         string
}

func NewJwksIDTokenValidator(jwksStore utils.JwksGetter, tokenSourceKey, tokenIdentifierKey, idTokenKey string) *jwksIDTokenValidator {
	return &jwksIDTokenValidator{
		jwksStore:          jwksStore,
		tokenSourceKey:     tokenSourceKey,
		tokenIdentifierKey: tokenIdentifierKey,
		idTokenKey:         idTokenKey,
	}
}

func (u *jwksIDTokenValidator) TokenSourceKey() string {
	return u.tokenSourceKey
}

func (u *jwksIDTokenValidator) TokenIdentifierKey() string {
	return u.tokenIdentifierKey
}

func (u *jwksIDTokenValidator) IDTokenKey() string {
	return u.idTokenKey
}

func (u *jwksIDTokenValidator) Validate(idToken string) (*jwt.Token, *dto.IDTokenClaims, error) {
	if idToken == "" {
		return nil, nil, dto.ErrIDTokenNotFound
	}

	jwksJson, err := u.jwksStore.Get()
	if err != nil {
		return nil, nil, err
	}

	// Create the JWKS from the resource at the given URL.
	jwks, err := keyfunc.NewJSON(jwksJson)
	if err != nil {
		return nil, nil, err
	}

	// Parse the JWT.
	// token, err := jwt.Parse(idToken, jwks.Keyfunc)
	token, err := jwt.ParseWithClaims(idToken, &dto.IDTokenClaims{}, jwks.Keyfunc)
	if token.Valid {
		claims, ok := token.Claims.(*dto.IDTokenClaims)
		if !ok {
			return nil, nil, dto.ErrUnexpectedTokenClaims
		}
		return token, claims, err
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		return nil, nil, dto.ErrTokenMalformed
	} else if errors.Is(err, jwt.ErrTokenExpired) {
		claims, ok := token.Claims.(*dto.IDTokenClaims)
		if !ok {
			return nil, nil, dto.ErrTokenExpired
		}
		return token, claims, dto.ErrTokenExpired
	} else if errors.Is(err, jwt.ErrTokenNotValidYet) {
		// return nil, auth.ErrTokenNotValidYet
		claims, ok := token.Claims.(*dto.IDTokenClaims)
		if !ok {
			return nil, nil, dto.ErrUnexpectedTokenClaims
		}
		return token, claims, nil
	} else if errors.Is(err, jwt.ErrTokenUsedBeforeIssued) {
		claims, ok := token.Claims.(*dto.IDTokenClaims)
		if !ok {
			return nil, nil, dto.ErrUnexpectedTokenClaims
		}
		return token, claims, nil
	}

	claims, ok := token.Claims.(*dto.IDTokenClaims)
	if !ok {
		return nil, nil, err
	}
	return token, claims, err
}
