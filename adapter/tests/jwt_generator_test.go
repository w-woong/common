package adapter_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"github.com/w-woong/common/adapter"
	"github.com/w-woong/common/utils"
)

func Test_es256SignedJwtGenerator_GenerateToken(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}

	privateKey, err := utils.LoadPKCS8PrivateKey("./private_key_es256.p8")
	if !assert.Nil(t, err) {
		t.FailNow()
	}
	gen, err := adapter.NewES256SignedJwtGenerator("my_kid", privateKey)
	if !assert.Nil(t, err) {
		t.FailNow()
	}

	now := time.Now()
	exp := now.Add(time.Second * 360)
	claims := jwt.RegisteredClaims{
		Issuer:    "my_issuer",
		IssuedAt:  jwt.NewNumericDate(time.Unix(now.Unix(), 0)),
		ExpiresAt: jwt.NewNumericDate(time.Unix(exp.Unix(), 0)),
		Audience:  []string{"my_audience"},
		Subject:   "my_subject",
	}
	token, err := gen.GenerateToken(&claims)
	if !assert.Nil(t, err) {
		t.FailNow()
	}

	fmt.Println(token)
}
