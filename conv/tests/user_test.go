package conv_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/w-woong/common/conv"
	"github.com/w-woong/common/dto"
	pb "github.com/w-woong/common/dto/protos/user/v1"
)

func Test_User_ToPasswordProtoFromDto(t *testing.T) {
	userDto := dto.CredentialPassword{}
	p, err := conv.ToCredentialPasswordProtoFromDto(userDto)
	assert.Nil(t, err)
	assert.NotNil(t, p)
}

func Test_User_ToEmailDtoFromProto(t *testing.T) {
	// p := pb.Email{}
	d, err := conv.ToEmailDtoFromProto(nil)
	assert.Nil(t, err)
	fmt.Println(d)
}

func Test_User_ToUserDtoFromProto(t *testing.T) {
	user := pb.User{}
	fmt.Println(user.GetCreatedAt().AsTime())
}
