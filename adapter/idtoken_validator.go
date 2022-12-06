package adapter

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/MicahParks/keyfunc"
	"github.com/golang-jwt/jwt/v4"
	"github.com/w-woong/common"
	"github.com/w-woong/common/dto"
)

type jwksIDTokenValidator struct {
	jwksUrl string
}

func NewJwksIDTokenValidator(jwksUrl string) *jwksIDTokenValidator {
	return &jwksIDTokenValidator{
		jwksUrl: jwksUrl,
	}
}

func (u *jwksIDTokenValidator) Validate(idToken string) (*jwt.Token, *dto.IDTokenClaims, error) {
	if idToken == "" {
		return nil, nil, common.ErrIDTokenNotFound
	}

	jwksJson, err := GetJwks(u.jwksUrl)
	if err != nil {
		return nil, nil, err
	}

	// Create the JWKS from the resource at the given URL.
	jwks, err := keyfunc.NewJSON(jwksJson)
	if err != nil {
		// log.Fatalf("Failed to create JWKS from resource at the given URL.\nError: %s", err.Error())
		return nil, nil, err
	}

	// Parse the JWT.
	// token, err := jwt.Parse(idToken, jwks.Keyfunc)
	token, err := jwt.ParseWithClaims(idToken, &dto.IDTokenClaims{}, jwks.Keyfunc)
	if token.Valid {
		claims, ok := token.Claims.(*dto.IDTokenClaims)
		if !ok {
			return nil, nil, common.ErrUnexpectedTokenClaims
		}
		return token, claims, err
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		return nil, nil, common.ErrTokenMalformed
	} else if errors.Is(err, jwt.ErrTokenExpired) {
		return nil, nil, common.ErrTokenExpired
	} else if errors.Is(err, jwt.ErrTokenNotValidYet) {
		// return nil, auth.ErrTokenNotValidYet
		claims, ok := token.Claims.(*dto.IDTokenClaims)
		if !ok {
			return nil, nil, common.ErrUnexpectedTokenClaims
		}
		return token, claims, nil
	} else if errors.Is(err, jwt.ErrTokenUsedBeforeIssued) {
		claims, ok := token.Claims.(*dto.IDTokenClaims)
		if !ok {
			return nil, nil, common.ErrUnexpectedTokenClaims
		}
		return token, claims, nil
	}

	return nil, nil, err
}

func GetJwks(url string) (json.RawMessage, error) {
	// TODO: cache jwks

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var jwksJSON json.RawMessage = b

	return jwksJSON, nil
}

func GetJwksUrl(openIDConfUrl string) (string, error) {

	res, err := http.Get(openIDConfUrl)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	resb, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	m := make(map[string]interface{})
	if err = json.Unmarshal(resb, &m); err != nil {
		return "", err
	}
	jwksUrl, ok := m["jwks_uri"]
	if !ok {
		return "", errors.New("not found")
	}
	return jwksUrl.(string), nil
}
