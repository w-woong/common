package utils_test

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
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
	jwtToken, err := utils.ParseJWTWithClaimsJwks(token, &OpenIDClaims{}, jwks)
	assert.Nil(t, err)
	outClaims, ok := jwtToken.Claims.(*OpenIDClaims)
	assert.Equal(t, true, ok)
	fmt.Println(outClaims)
}

func Test_GenerateES256SignedJWT(t *testing.T) {
	signingKey := `-----BEGIN PRIVATE KEY-----
MIGTAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBHkwdwIBAQQg+94fs23vSrhBIXNz
OdeRb7+FJkIsVrnTSf7eIYKdf4mgCgYIKoZIzj0DAQehRANCAATyBS3eRgOJ53OQ
LFhGSJw4aiqju7muVwoIWFxCcFJasRwyGcbs0C7vt3xKV/DRJvID4UljaI53wETq
RxlkNCeV
-----END PRIVATE KEY-----`
	clientID := "my_client_id"
	teamID := "my_team_id"
	audience := []string{"https://appleid.apple.com"}
	kid := "my_kid"
	alg := "ES256"

	block, _ := pem.Decode([]byte(signingKey))
	// assert.Nil(t, err)
	if !assert.NotNil(t, block) {
		t.FailNow()
	}

	privKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	assert.Nil(t, err)

	now := time.Now()
	exp := now.Add(time.Minute * 5)
	claims := &jwt.RegisteredClaims{
		Issuer:    teamID,
		IssuedAt:  jwt.NewNumericDate(time.Unix(now.Unix(), 0)),
		ExpiresAt: jwt.NewNumericDate(time.Unix(exp.Unix(), 0)),
		Audience:  audience,
		Subject:   clientID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token.Header["kid"] = kid
	token.Header["alg"] = alg

	signed, err := token.SignedString(privKey)
	assert.Nil(t, err)
	fmt.Println(signed)
}

func Test_GenerateES256SignedJWT_file(t *testing.T) {

	clientID := "my_client_id"
	teamID := "my_team_id"
	audience := []string{"https://appleid.apple.com"}
	kid := "my_kid"
	// alg := "ES256"

	privKey, err := utils.LoadPKCS8PrivateKey("./private_key_es256.p8")
	if !assert.Nil(t, err) {
		t.FailNow()
	}

	now := time.Now()
	exp := now.Add(time.Minute * 5)
	claims := &jwt.RegisteredClaims{
		Issuer:    teamID,
		IssuedAt:  jwt.NewNumericDate(time.Unix(now.Unix(), 0)),
		ExpiresAt: jwt.NewNumericDate(time.Unix(exp.Unix(), 0)),
		Audience:  audience,
		Subject:   clientID,
	}

	signed, err := utils.GenerateES256SignedJWT(kid, privKey, claims)
	if !assert.Nil(t, err) {
		t.FailNow()
	}
	fmt.Println(signed)
}
