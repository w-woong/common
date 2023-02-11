package middlewares

import (
	"context"
	"errors"
	"net/http"

	"github.com/w-woong/common"
	"github.com/w-woong/common/dto"
	"github.com/w-woong/common/logger"
	"github.com/w-woong/common/port"
	"github.com/w-woong/common/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func AuthIDToken(next http.HandlerFunc, cookie port.TokenCookie, parser port.IDTokenParser) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		idToken := utils.AuthBearer(r)
		if idToken == "" {
			idToken = cookie.GetIDToken(r)
		}
		jwtToken, err := parser.ParseWithClaims(idToken, &dto.IDTokenClaims{})
		if err != nil {
			if errors.Is(err, common.ErrTokenExpired) {
				common.HttpErrorWithBody(w, http.StatusUnauthorized,
					common.NewHttpBody(http.StatusText(http.StatusUnauthorized), common.StatusTryRefreshIDToken))
				logger.Error(http.StatusText(http.StatusUnauthorized), logger.UrlField(r.URL.String()))
				return
			}
			common.HttpError(w, http.StatusUnauthorized)
			logger.Error(http.StatusText(http.StatusUnauthorized), logger.UrlField(r.URL.String()))
			return
		}

		claims, ok := jwtToken.Claims.(*dto.IDTokenClaims)
		if !ok {
			common.HttpError(w, http.StatusUnauthorized)
			logger.Error(http.StatusText(http.StatusUnauthorized), logger.UrlField(r.URL.String()))
			return
		}
		ctx = context.WithValue(ctx, dto.IDTokenClaimsContextKey{}, *claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func AuthIDTokenUserAccountSvc(next http.HandlerFunc, cookie port.TokenCookie, parser port.IDTokenParser,
	userSvc port.UserSvc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		idToken := utils.AuthBearer(r)
		if idToken == "" {
			idToken = cookie.GetIDToken(r)
		}
		jwtToken, err := parser.ParseWithClaims(idToken, &dto.IDTokenClaims{})
		if err != nil {
			if errors.Is(err, common.ErrTokenExpired) {
				common.HttpErrorWithBody(w, http.StatusUnauthorized,
					common.NewHttpBody(http.StatusText(http.StatusUnauthorized), common.StatusTryRefreshIDToken))
				logger.Error(http.StatusText(http.StatusUnauthorized), logger.UrlField(r.URL.String()))
				return
			}
			common.HttpError(w, http.StatusUnauthorized)
			logger.Error(http.StatusText(http.StatusUnauthorized), logger.UrlField(r.URL.String()))
			return
		}

		claims, ok := jwtToken.Claims.(*dto.IDTokenClaims)
		if !ok {
			common.HttpError(w, http.StatusUnauthorized)
			logger.Error(http.StatusText(http.StatusUnauthorized), logger.UrlField(r.URL.String()))
			return
		}

		userAccount, err := userSvc.FindByIDToken(ctx, idToken)
		if err != nil {
			common.HttpError(w, http.StatusUnauthorized)
			logger.Error(http.StatusText(http.StatusUnauthorized), logger.UrlField(r.URL.String()))
			return
		}
		ctx = context.WithValue(ctx, dto.IDTokenClaimsContextKey{}, *claims)
		ctx = context.WithValue(ctx, dto.UserAccountContextKey{}, userAccount)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

type IDTokenerGrpc interface {
	GetIdToken() string
}

func AuthIDTokenGrpc(parser port.IDTokenParser) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if info.FullMethod != "/UserService/FindByLoginID" {

			return handler(ctx, req)
		}

		switch v := req.(type) {
		case (IDTokenerGrpc):
			idToken := v.GetIdToken()

			jwtToken, err := parser.ParseWithClaims(idToken, &dto.IDTokenClaims{})
			if err != nil {
				if errors.Is(err, common.ErrTokenExpired) {
					return nil, status.Error(codes.Code(common.StatusTryRefreshIDToken), err.Error())
				}
				return nil, err
			}
			claims, ok := jwtToken.Claims.(*dto.IDTokenClaims)
			if !ok {
				return nil, status.Error(codes.Code(common.StatusInvalidTokenClaims), common.ErrTokenInvalidClaims.Error())
			}
			ctx = context.WithValue(ctx, dto.IDTokenClaimsContextKey{}, *claims)
			return handler(ctx, req)
		default:
			return nil, common.ErrIDTokenNotFound
		}
	}
}

func AuthIDTokenUserAccountGrpc(parser port.IDTokenParser, userSvc port.UserSvc) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if info.FullMethod != "/UserService/FindByLoginID" {

			return handler(ctx, req)
		}

		switch v := req.(type) {
		case IDTokenerGrpc:
			idToken := v.GetIdToken()

			jwtToken, err := parser.ParseWithClaims(idToken, &dto.IDTokenClaims{})
			if err != nil {
				if errors.Is(err, common.ErrTokenExpired) {
					return nil, status.Error(codes.Code(common.StatusTryRefreshIDToken), err.Error())
				}
				return nil, err
			}
			claims, ok := jwtToken.Claims.(*dto.IDTokenClaims)
			if !ok {
				return nil, status.Error(codes.Code(common.StatusInvalidTokenClaims), common.ErrTokenInvalidClaims.Error())
			}

			userAccount, err := userSvc.FindByIDToken(ctx, idToken)
			if err != nil {
				return nil, err
			}
			ctx = context.WithValue(ctx, dto.IDTokenClaimsContextKey{}, *claims)
			ctx = context.WithValue(ctx, dto.UserAccountContextKey{}, userAccount)
			return handler(ctx, req)
		default:
			return nil, common.ErrIDTokenNotFound
		}
	}
}
