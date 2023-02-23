package utils_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/w-woong/common/utils"
)

func Test_getJwks(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}

	h, err := utils.NewJwksHttp("https://appleid.apple.com/auth/keys")
	if !assert.Nil(t, err) {
		t.FailNow()
	}

	fmt.Println(h.Get())
}
