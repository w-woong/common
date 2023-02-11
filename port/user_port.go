package port

import (
	"context"

	"github.com/w-woong/common/dto"
)

type UserSvc interface {
	RegisterUser(ctx context.Context, user dto.User) (dto.User, error)
	FindByIDToken(ctx context.Context, idToken string) (dto.User, error)
}
