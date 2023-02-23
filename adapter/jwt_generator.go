package adapter

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/w-woong/common/utils"
)

type es256SignedJwtGenerator struct {
	kid        string
	privateKey any
}

func NewES256SignedJwtGenerator(kid string, privateKeyFileName string) (*es256SignedJwtGenerator, error) {
	privateKey, err := utils.LoadPKCS8PrivateKey(privateKeyFileName)
	if err != nil {
		return nil, err
	}

	return &es256SignedJwtGenerator{
		kid:        kid,
		privateKey: privateKey,
	}, nil

}

func (a *es256SignedJwtGenerator) GenerateToken(claims jwt.Claims) (string, error) {
	return utils.GenerateES256SignedJWT(a.kid, a.privateKey, claims)
}
