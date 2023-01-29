package adapter_test

import (
	"context"
	"testing"

	"github.com/go-wonk/si/sihttp"
	"github.com/stretchr/testify/assert"
	"github.com/w-woong/common/adapter"
	"github.com/w-woong/common/dto"
)

func Test_userHttp_RegisterUser(t *testing.T) {

	if !onlinetest {
		t.Skip("skipping online tests")
	}

	ctx := context.Background()
	userHttp := adapter.NewUserHttp(sihttp.DefaultInsecureClient(), "https://localhost:49007", "ab2316584873095f017f6dfa7a9415794f563fcc473eb3fe65b9167e37fd5a4b",
		"token_source", "tid", "id_token")

	_, err := userHttp.RegisterUser(ctx, dto.User{
		LoginID:     "asdf",
		LoginType:   "token",
		LoginSource: "google",
	})

	assert.Nil(t, err)
}
