package port

import (
	"context"

	"github.com/golang-jwt/jwt/v4"
	"github.com/w-woong/common/dto"
)

type IDTokenRefresher interface {
	Refresh(ctx context.Context, tokenSource, tokenIdentifier string, idToken string) (dto.Token, error)
}

type IDTokenValidators map[string]IDTokenValidator

type IDTokenValidator interface {
	Validate(idToken string) (*jwt.Token, *dto.IDTokenClaims, error)
}
