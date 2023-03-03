package adapter

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
	"github.com/w-woong/common"
	"github.com/w-woong/common/utils"
)

type jwksIDTokenParser struct {
	jwksStore utils.JwksGetter

	openidConfUrl  string
	defaultJwksUri string
}

func NewJwksIDTokenParser(jwksStore utils.JwksGetter) *jwksIDTokenParser {
	return &jwksIDTokenParser{
		jwksStore: jwksStore,
	}
}

func NewJwksIDTokenParserWithUrl(openidConfUrl string, defaultJwksUri string) *jwksIDTokenParser {
	return &jwksIDTokenParser{
		openidConfUrl:  openidConfUrl,
		defaultJwksUri: defaultJwksUri,
	}
}

func (u *jwksIDTokenParser) getJwksStore() (utils.JwksGetter, error) {
	if u.jwksStore != nil {
		return u.jwksStore, nil
	}

	var err error
	jwksUrl := ""
	if u.openidConfUrl != "" {
		jwksUrl, err = utils.GetJwksUrl(u.openidConfUrl)
		if err != nil {
			return nil, err
		}
	}
	if jwksUrl == "" {
		jwksUrl = u.defaultJwksUri
	}

	if jwksUrl == "" {
		return nil, errors.New("invalid jwks url")
	}

	u.jwksStore, err = utils.NewJwksCache(jwksUrl)
	if err != nil {
		return nil, err
	}
	return u.jwksStore, nil
}

func (u *jwksIDTokenParser) ParseWithClaims(idToken string, claims jwt.Claims) (*jwt.Token, error) {
	if idToken == "" {
		return nil, common.ErrIDTokenNotFound
	}

	store, err := u.getJwksStore()
	if err != nil {
		return nil, err
	}

	jwksJson, err := store.Get()
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
