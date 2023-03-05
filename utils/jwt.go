package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/MicahParks/keyfunc"
	"github.com/golang-jwt/jwt/v4"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/w-woong/common/dto"
)

func LoadRSAPrivateKey(fileName string) (*rsa.PrivateKey, error) {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(b)
	if err != nil {
		return nil, err
	}
	return key, nil

}

func LoadRSAPublicKey(fileName string) (*rsa.PublicKey, error) {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(b)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func LoadPKCS8PrivateKey(fileName string) (any, error) {
	signingKey, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(signingKey)
	if block == nil {
		return nil, fmt.Errorf("failure decoding %s", fileName)
	}
	privKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return privKey, nil

}

func RSAPublicKeyToJwks(fileNames []string, kids []string) ([]byte, error) {
	keys := make([]jwk.RSAPublicKey, 0)
	for i, fileName := range fileNames {
		rsaKey, err := LoadRSAPublicKey(fileName)
		if err != nil {
			return nil, err
		}
		key, err := jwk.FromRaw(rsaKey)
		if err != nil {
			return nil, err
		}
		jwkPub, ok := key.(jwk.RSAPublicKey)
		if !ok {
			return nil, fmt.Errorf("expected jwk.RSAPrivateKey, got %T", jwkPub)
		}

		jwkPub.Set(jwk.KeyIDKey, kids[i])

		keys = append(keys, jwkPub)
	}

	keySets := make(map[string]interface{})
	keySets["keys"] = keys
	return json.Marshal(keySets)
}

func GenerateRS256SignedJWT(kid string, key *rsa.PrivateKey, claims jwt.Claims) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	jwtToken.Header["kid"] = kid
	return jwtToken.SignedString(key)
}

func GenerateES256SignedJWT(kid string, key any, claims jwt.Claims) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	jwtToken.Header["kid"] = kid
	return jwtToken.SignedString(key)
}

func ParseJWTWithClaimsJwks(token string, claims jwt.Claims, jwksBytes []byte) (*jwt.Token, error) {
	if token == "" {
		return nil, errors.New("token is empty")
	}
	jwks, err := keyfunc.NewJSON(jwksBytes)
	if err != nil {
		return nil, err
	}
	jwtToken, err := jwt.ParseWithClaims(token, claims, jwks.Keyfunc)
	if jwtToken.Valid {
		return jwtToken, err
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		return nil, dto.ErrTokenMalformed
	} else if errors.Is(err, jwt.ErrTokenExpired) {
		return jwtToken, dto.ErrTokenExpired
	} else if errors.Is(err, jwt.ErrTokenNotValidYet) {
		return jwtToken, nil
	} else if errors.Is(err, jwt.ErrTokenUsedBeforeIssued) {
		return jwtToken, nil
	}
	return jwtToken, err
}
