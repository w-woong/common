package conv_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/w-woong/common/conv"
	"github.com/w-woong/common/dto"
)

func Test_User_ToPasswordProtoFromDto(t *testing.T) {
	userDto := dto.Password{}
	p, err := conv.ToPasswordProtoFromDto(userDto)
	assert.Nil(t, err)
	assert.NotNil(t, p)
}

func Test_User_ToEmailDtoFromProto(t *testing.T) {
	// p := pb.Email{}
	d, err := conv.ToEmailDtoFromProto(nil)
	assert.Nil(t, err)
	fmt.Println(d)
}
