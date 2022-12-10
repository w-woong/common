package adapter

import (
	"context"

	"github.com/w-woong/common/dto"
)

type UserSvcNop struct {
}

func NewUserSvcNop() *UserSvcNop {
	return &UserSvcNop{}
}
func (UserSvcNop) RegisterUser(ctx context.Context, loginSource string, user dto.User) (dto.User, error) {
	return dto.NilUser, nil
}

func (UserSvcNop) FindByLoginID(ctx context.Context, tokenIdentifier, idToken string) (dto.User, error) {
	return dto.NilUser, nil
}
