package port

import (
	"context"

	"github.com/w-woong/common/dto"
)

type UserSvc interface {
	RegisterUser(ctx context.Context, user dto.User) (dto.User, error)
	FindByLoginID(ctx context.Context, loginSource, tokenIdentifier, idToken string) (dto.User, error)
}
