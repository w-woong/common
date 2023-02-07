package utils_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"github.com/w-woong/common/utils"
)

type OpenIDClaims struct {
	Foo string `json:"foo"`
	jwt.RegisteredClaims
}

func Test_GenerateRS256SignedJWT(t *testing.T) {
	b, _ := json.Marshal([]string{"a", "b"})
	fmt.Println(string(b))
	privKey, _ := utils.LoadRSAPrivateKey("jwk_private_key.pem")
	// pubKey, _ := utils.LoadRSAPublicKey("jwk_public_key.pem")

	// Create the claims
	claims := OpenIDClaims{
		"bar",
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "test",
			Subject:   "somebody",
			ID:        "1",
			Audience:  []string{"somebody_else"},
		},
	}

	token, err := utils.GenerateRS256SignedJWT("myid", privKey, &claims)
	assert.Nil(t, err)

	jwks, err := utils.RSAPublicKeyToJwks([]string{"jwk_public_key.pem"}, []string{"myid"})
	assert.Nil(t, err)
	jwtToken, err := utils.ParseRS256SignedJWT(token, &OpenIDClaims{}, jwks)
	assert.Nil(t, err)
	outClaims, ok := jwtToken.Claims.(*OpenIDClaims)
	assert.Equal(t, true, ok)
	fmt.Println(outClaims)
}
