package adapter_test

import (
	"fmt"
	"testing"

	"github.com/w-woong/common/adapter"
	"github.com/w-woong/common/utils"
)

func Test_jwksIDTokenValidator_Validate(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}

	jwksUrl, _ := utils.GetJwksUrl("https://cognito-idp.ap-northeast-2.amazonaws.com/ap-northeast-2_????/.well-known/openid-configuration")
	jwksStore, _ := utils.NewJwksCache(jwksUrl)
	validator := adapter.NewJwksIDTokenValidator(jwksStore,
		"", "", "")

	idToken := ""
	fmt.Println(validator.Validate(idToken))

}
