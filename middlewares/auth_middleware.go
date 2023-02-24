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

// GetIDTokenJwtAndClaims validates and refresh id_token. It first retrieves id_token from request, r.
// Then, it parses the token to get the claims. It ignores ErrTokenExpired.
func GetIDToken(next http.HandlerFunc, cookie port.TokenCookie) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		idToken := utils.AuthBearer(r)
		if idToken == "" {
			idToken = cookie.GetIDToken(r)
		}

		ctx = context.WithValue(ctx, dto.IDTokenContextKey{}, idToken)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// GetIDTokenJwtAndClaims validates and refresh id_token. It first retrieves id_token from request, r.
// Then, it parses the token to get the claims. It ignores ErrTokenExpired.
func GetIDTokenJwtAndClaims(next http.HandlerFunc, cookie port.TokenCookie, parser port.IDTokenParser) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		idToken := utils.AuthBearer(r)
		if idToken == "" {
			idToken = cookie.GetIDToken(r)
		}

		jwtToken, err := parser.ParseWithClaims(idToken, &dto.IDTokenClaims{})
		if err != nil {
			if !errors.Is(err, common.ErrTokenExpired) {
				oerr := common.OAuth2ErrorInvalidRequest(err.Error(), http.StatusUnauthorized)
				http.Error(w, oerr.Error(), oerr.GetStatusCode())
				logger.Error(oerr.Error(), logger.UrlField(r.URL.String()))
				return
			}
		}

		claims, ok := jwtToken.Claims.(*dto.IDTokenClaims)
		if !ok {
			oerr := common.OAuth2ErrorInvalidClaims()
			http.Error(w, oerr.Error(), oerr.GetStatusCode())
			logger.Error(oerr.Error(), logger.UrlField(r.URL.String()))
			return
		}

		ctx = context.WithValue(ctx, dto.IDTokenClaimsContextKey{}, *claims)
		ctx = context.WithValue(ctx, dto.JwtTokenContextKey{}, *jwtToken)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

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
				oerr := common.OAuth2ErrorTryRefresh(err.Error())
				http.Error(w, oerr.Error(), oerr.GetStatusCode())
				logger.Error(oerr.Error(), logger.UrlField(r.URL.String()))

				// common.HttpErrorWithBody(w, http.StatusUnauthorized,
				// 	common.NewHttpBody(http.StatusText(http.StatusUnauthorized), common.StatusTryRefreshIDToken))
				// logger.Error(http.StatusText(http.StatusUnauthorized), logger.UrlField(r.URL.String()))
				return
			}

			oerr := common.OAuth2ErrorInvalidRequest(err.Error(), http.StatusUnauthorized)
			http.Error(w, oerr.Error(), oerr.GetStatusCode())
			logger.Error(oerr.Error(), logger.UrlField(r.URL.String()))
			return
		}

		claims, ok := jwtToken.Claims.(*dto.IDTokenClaims)
		if !ok {
			oerr := common.OAuth2ErrorInvalidClaims()
			http.Error(w, oerr.Error(), oerr.GetStatusCode())
			logger.Error(oerr.Error(), logger.UrlField(r.URL.String()))
			return
		}
		ctx = context.WithValue(ctx, dto.IDTokenClaimsContextKey{}, *claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func AuthIDTokenIgnoreErr(next http.HandlerFunc, cookie port.TokenCookie, parser port.IDTokenParser) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		idToken := utils.AuthBearer(r)
		if idToken == "" {
			idToken = cookie.GetIDToken(r)
		}

		jwtToken, _ := parser.ParseWithClaims(idToken, &dto.IDTokenClaims{})
		if jwtToken == nil {
			ctx = context.WithValue(ctx, dto.IDTokenClaimsContextKey{}, dto.IDTokenClaims{})
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		claims, ok := jwtToken.Claims.(*dto.IDTokenClaims)
		if !ok || claims == nil {
			ctx = context.WithValue(ctx, dto.IDTokenClaimsContextKey{}, dto.IDTokenClaims{})
			next.ServeHTTP(w, r.WithContext(ctx))
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
				oerr := common.OAuth2ErrorTryRefresh(err.Error())
				http.Error(w, oerr.Error(), oerr.GetStatusCode())
				logger.Error(oerr.Error(), logger.UrlField(r.URL.String()))

				return
			}
			oerr := common.OAuth2ErrorInvalidRequest(err.Error(), http.StatusUnauthorized)
			http.Error(w, oerr.Error(), oerr.GetStatusCode())
			logger.Error(oerr.Error(), logger.UrlField(r.URL.String()))
			return
		}

		claims, ok := jwtToken.Claims.(*dto.IDTokenClaims)
		if !ok {
			oerr := common.OAuth2ErrorInvalidClaims()
			http.Error(w, oerr.Error(), oerr.GetStatusCode())
			logger.Error(oerr.Error(), logger.UrlField(r.URL.String()))
			return
		}

		userAccount, err := userSvc.FindByIDToken(ctx, idToken)
		if err != nil {
			oerr := common.OAuth2ErrorInvalidRequest(err.Error(), http.StatusUnauthorized)
			http.Error(w, oerr.Error(), oerr.GetStatusCode())
			logger.Error(oerr.Error(), logger.UrlField(r.URL.String()))
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
