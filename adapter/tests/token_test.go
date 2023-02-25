package adapter_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-wonk/si/v2/sihttp"
	"github.com/stretchr/testify/assert"
	"github.com/w-woong/common/adapter"
)

func TestRefresh(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}

	client := sihttp.DefaultInsecureStandardClient()
	svc := adapter.NewIDTokenHttp(client, "https://localhost:5558",
		"ab2316584873095f017f6dfa7a9415794f563fcc473eb3fe65b9167e37fd5a4b", "tid", "id_token")

	id := "id"
	idToken := `asdf`
	token, err := svc.Refresh(context.Background(), "google", id, idToken)
	assert.NotNil(t, err)

	fmt.Println(token)
}
