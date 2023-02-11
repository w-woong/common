package adapter

import (
	"context"

	"github.com/w-woong/common/dto"
	"github.com/w-woong/common/port"
)

var _ port.UserSvc = (*UserSvcNop)(nil)

type UserSvcNop struct {
}

func NewUserSvcNop() *UserSvcNop {
	return &UserSvcNop{}
}
func (UserSvcNop) RegisterUser(ctx context.Context, user dto.User) (dto.User, error) {
	return dto.NilUser, nil
}

func (UserSvcNop) FindByIDToken(ctx context.Context, idToken string) (dto.User, error) {
	return dto.NilUser, nil
}
