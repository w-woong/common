package utils_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/w-woong/common/utils"
)

func Test_Sha512WithSalt(t *testing.T) {
	h, err := utils.Sha512WithSaltHex([]byte("hello_sir?"), nil, []byte("sh32ye4Nd3o932Djqqdtnm4v"))
	assert.Nil(t, err)

	assert.Equal(t, "3e2d850e36c1cb0785df0bdd0bc4b704c22ee0b421f0f16d227e290f5211b5ccf5a80a00445de6343fb27556342f658bdbc3cf64d8815778c2be8ee1ebee6185", h)
}

func Test_HmacSha512HexEncoded(t *testing.T) {
	h, err := utils.HmacSha512HexEncoded("sh32ye4Nd3o932Djqqdtnm4v", []byte("hello_sir?"))
	assert.Nil(t, err)

	fmt.Println(h)
	assert.Equal(t, "978ee85599b8c00246287add2b8641b39fe42fbe7523feabe684f0d8d993edbfe6ae066b38a45a4b806caf6ef64344c01bd8672c5ed359dfccdcd401739a0048", h)

}
