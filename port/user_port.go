package port

import (
	"context"

	"github.com/w-woong/common/dto"
)

type UserSvc interface {
	RegisterUser(ctx context.Context, loginSource string, user dto.User) (dto.User, error)
	FindByLoginID(ctx context.Context, tokenIdentifier, idToken string) (dto.User, error)
}
