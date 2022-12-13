package adapter

import (
	"context"
	"errors"
	"net/http"

	"github.com/w-woong/common"
	"github.com/w-woong/common/conv"
	"github.com/w-woong/common/dto"
	pb "github.com/w-woong/common/dto/protos/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type userGrpc struct {
	client pb.UserServiceClient
}

func NewUserGrpc(conn *grpc.ClientConn) *userGrpc {
	return &userGrpc{
		client: pb.NewUserServiceClient(conn),
	}
}

func (a *userGrpc) RegisterUser(ctx context.Context, user dto.User) (dto.User, error) {

	userProto, err := conv.ToUserProtoFromDto(user)
	if err != nil {
		return dto.NilUser, err
	}

	reply, err := a.client.RegisterUser(ctx, &pb.RegisterUserRequest{
		LoginSource: userProto.LoginSource,
		Document:    userProto,
	})
	if err != nil {
		if s, ok := status.FromError(err); ok {
			if common.StatusTryRefreshIDToken == int(s.Code()) {
				return dto.NilUser, common.ErrTokenExpired
			}
		}
		return dto.NilUser, err
	}

	if reply.Status != http.StatusOK {
		return dto.NilUser, errors.New(reply.GetMessage())
	}

	return conv.ToUserDtoFromProto(reply.Document)
}

func (a *userGrpc) FindByLoginID(ctx context.Context, loginSource, tokenIdentifier, idToken string) (dto.User, error) {
	reply, err := a.client.FindByLoginID(ctx, &pb.FindByLoginIDRequest{
		Tid:         tokenIdentifier,
		TokenSource: loginSource,
		IdToken:     idToken,
	})
	if err != nil {
		if s, ok := status.FromError(err); ok {
			if common.StatusTryRefreshIDToken == int(s.Code()) {
				return dto.NilUser, common.ErrTokenExpired
			}
		}
		return dto.NilUser, err
	}

	if reply.Status != http.StatusOK {
		return dto.NilUser, errors.New(reply.GetMessage())
	}

	return conv.ToUserDtoFromProto(reply.Document)
}
