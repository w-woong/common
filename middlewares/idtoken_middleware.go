package middlewares

// import (
// 	"context"
// 	"errors"
// 	"net/http"
// 	"strings"

// 	"github.com/golang-jwt/jwt/v4"
// 	"github.com/w-woong/common"
// 	"github.com/w-woong/common/dto"
// 	"github.com/w-woong/common/logger"
// 	"github.com/w-woong/common/port"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/status"
// 	// "github.com/grpc-ecosystem/go-grpc-middleware"
// )

// func GetIDToken(r *http.Request, key string) (string, error) {
// 	idToken := ""
// 	cookie, err := r.Cookie(key)
// 	if err == nil {
// 		idToken = cookie.Value
// 	}

// 	if idToken == "" {
// 		idToken = r.Header.Get(key)
// 		if strings.HasPrefix(idToken, "Bearer") {
// 			authVals := strings.Split(idToken, " ")
// 			if len(authVals) != 2 {
// 				return "", common.ErrIDTokenNotFound
// 			}
// 			idToken = authVals[1]
// 		}
// 	}

// 	if idToken == "" {
// 		return "", common.ErrIDTokenNotFound
// 	}

// 	return idToken, nil
// }

// func GetTokenSource(r *http.Request, key string) (string, error) {
// 	tokenSource := ""
// 	cookie, err := r.Cookie(key)
// 	if err == nil {
// 		tokenSource = cookie.Value
// 	}
// 	if tokenSource == "" {
// 		tokenSource = r.Header.Get(key)
// 	}

// 	if tokenSource == "" {
// 		return "", common.ErrTokenSourceNotFound
// 	}

// 	return tokenSource, nil
// }
// func GetTokenIdentifier(r *http.Request, key string) (string, error) {
// 	val := ""
// 	cookie, err := r.Cookie(key)
// 	if err == nil {
// 		val = cookie.Value
// 	}
// 	if val == "" {
// 		val = r.Header.Get(key)
// 	}

// 	if val == "" {
// 		return "", common.ErrTokenIdentifierNotFound
// 	}

// 	return val, nil
// }

// func AuthIDTokenHandler(next http.HandlerFunc, validator port.IDTokenValidators) http.HandlerFunc {

// 	return func(w http.ResponseWriter, r *http.Request) {

// 		ctx := r.Context()

// 		_, claims, token, err := validate(r, validator)
// 		if err != nil {
// 			if errors.Is(err, common.ErrTokenExpired) {
// 				common.HttpErrorWithBody(w, http.StatusUnauthorized,
// 					common.NewHttpBody(http.StatusText(http.StatusUnauthorized), common.StatusTryRefreshIDToken))
// 				logger.Error(http.StatusText(http.StatusUnauthorized), logger.UrlField(r.URL.String()))
// 				return
// 			}
// 			common.HttpError(w, http.StatusUnauthorized)
// 			logger.Error(http.StatusText(http.StatusUnauthorized), logger.UrlField(r.URL.String()))
// 			return
// 		}
// 		ctx = context.WithValue(ctx, dto.IDTokenClaimsKey{}, *claims)
// 		ctx = context.WithValue(ctx, dto.TokenSourceKey{}, token.TokenSource)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	}
// }

// func validate(r *http.Request, validator port.IDTokenValidators) (*jwt.Token, *dto.IDTokenClaims, dto.Token, error) {
// 	for k, v := range validator {
// 		tokenSource, err := GetTokenSource(r, v.TokenSourceKey())
// 		if err != nil {
// 			return nil, nil, dto.NilToken, err
// 		}
// 		if strings.EqualFold(k, tokenSource) {
// 			tokenIdentifier, err := GetTokenIdentifier(r, v.TokenIdentifierKey())
// 			if err != nil {
// 				return nil, nil, dto.NilToken, err
// 			}
// 			idToken, err := GetIDToken(r, v.IDTokenKey())
// 			if err != nil {
// 				return nil, nil, dto.NilToken, err
// 			}
// 			jwtToken, claims, err := v.Validate(idToken)
// 			return jwtToken, claims, dto.Token{
// 				ID:          tokenIdentifier,
// 				IDToken:     idToken,
// 				TokenSource: tokenSource,
// 			}, err
// 		}
// 	}

// 	return nil, nil, dto.NilToken, common.ErrTokenSourceNotFound
// }

// func validateGrpc(validator port.IDTokenValidators, tokenSource, idToken string) (*jwt.Token, *dto.IDTokenClaims, error) {
// 	if v, ok := validator[tokenSource]; ok {
// 		return v.Validate(idToken)
// 	}

// 	return nil, nil, common.ErrTokenSourceNotFound
// }

// func AuthIDTokenUserAccountHandler(next http.HandlerFunc, validator port.IDTokenValidators,
// 	userSvc port.UserSvc) http.HandlerFunc {

// 	return func(w http.ResponseWriter, r *http.Request) {
// 		ctx := r.Context()

// 		_, claims, token, err := validate(r, validator)
// 		if err != nil {
// 			if errors.Is(err, common.ErrTokenExpired) {
// 				common.HttpErrorWithBody(w, http.StatusUnauthorized,
// 					common.NewHttpBody(http.StatusText(http.StatusUnauthorized), common.StatusTryRefreshIDToken))
// 				logger.Error(http.StatusText(http.StatusUnauthorized), logger.UrlField(r.URL.String()))
// 				return
// 			}
// 			common.HttpError(w, http.StatusUnauthorized)
// 			logger.Error(http.StatusText(http.StatusUnauthorized), logger.UrlField(r.URL.String()))
// 			return
// 		}
// 		userAccount, err := userSvc.FindByLoginID(ctx, token.TokenSource, token.ID, token.IDToken)
// 		if err != nil {
// 			common.HttpError(w, http.StatusUnauthorized)
// 			logger.Error(http.StatusText(http.StatusUnauthorized), logger.UrlField(r.URL.String()))
// 			return
// 		}
// 		ctx = context.WithValue(ctx, dto.IDTokenClaimsKey{}, *claims)
// 		ctx = context.WithValue(ctx, dto.TokenSourceKey{}, token.TokenSource)
// 		ctx = context.WithValue(ctx, dto.UserAccountKey{}, userAccount)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	}
// }

// type IDTokener interface {
// 	GetTid() string
// 	GetIdToken() string
// 	GetTokenSource() string
// }

// func AuthIDTokenInterceptor(validator port.IDTokenValidators) grpc.UnaryServerInterceptor {
// 	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
// 		if info.FullMethod != "/UserService/FindByLoginID" {

// 			return handler(ctx, req)
// 		}

// 		switch v := req.(type) {
// 		case (IDTokener):
// 			idToken := v.GetIdToken()
// 			tokenSource := v.GetTokenSource()

// 			_, claims, err := validateGrpc(validator, tokenSource, idToken)
// 			if err != nil {
// 				if errors.Is(err, common.ErrTokenExpired) {
// 					return nil, status.Error(codes.Code(common.StatusTryRefreshIDToken), err.Error())
// 				}
// 				return nil, err
// 			}
// 			ctx = context.WithValue(ctx, dto.IDTokenClaimsKey{}, *claims)
// 			ctx = context.WithValue(ctx, dto.TokenSourceKey{}, tokenSource)
// 			return handler(ctx, req)
// 		default:
// 			return nil, common.ErrIDTokenNotFound
// 		}
// 	}
// }

// func AuthIDTokenUserAccountInterceptor(validator port.IDTokenValidators, userSvc port.UserSvc) grpc.UnaryServerInterceptor {
// 	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
// 		if info.FullMethod != "/UserService/FindByLoginID" {

// 			return handler(ctx, req)
// 		}

// 		switch v := req.(type) {
// 		case IDTokener:
// 			idToken := v.GetIdToken()
// 			tokenSource := v.GetTokenSource()
// 			tokenIdentifier := v.GetTid()

// 			_, claims, err := validateGrpc(validator, tokenSource, idToken)
// 			if err != nil {
// 				if errors.Is(err, common.ErrTokenExpired) {
// 					return nil, status.Error(codes.Code(common.StatusTryRefreshIDToken), err.Error())
// 				}
// 				return nil, err
// 			}
// 			userAccount, err := userSvc.FindByLoginID(ctx, tokenSource, tokenIdentifier, idToken)
// 			if err != nil {
// 				return nil, err
// 			}
// 			ctx = context.WithValue(ctx, dto.IDTokenClaimsKey{}, *claims)
// 			ctx = context.WithValue(ctx, dto.TokenSourceKey{}, tokenSource)
// 			ctx = context.WithValue(ctx, dto.UserAccountKey{}, userAccount)
// 			return handler(ctx, req)
// 		default:
// 			return nil, common.ErrIDTokenNotFound
// 		}
// 	}
// }
