package adapter

import (
	"context"
	"errors"
	"net/http"
	"net/url"

	"github.com/go-wonk/si/sicore"
	"github.com/go-wonk/si/sihttp"
	"github.com/w-woong/common"
	"github.com/w-woong/common/dto"
)

type userHttp struct {
	client             *sihttp.Client
	baseUrl            string
	host               string
	bearerToken        string
	tokenSourceKey     string
	tokenIdentifierKey string
	idTokenKey         string
}

func NewUserHttp(client *http.Client, baseUrl string, bearerToken string,
	tokenSourceKey, tokenIdentifierKey, idTokenKey string) *userHttp {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json; charset=utf-8"

	c := sihttp.NewClient(client, sihttp.WithBaseUrl(baseUrl),
		sihttp.WithRequestOpt(sihttp.WithBearerToken(bearerToken)),
		sihttp.WithWriterOpt(sicore.SetJsonEncoder()),
		sihttp.WithReaderOpt(sicore.SetJsonDecoder()),
		sihttp.WithDefaultHeaders(headers))

	a := &userHttp{
		client:             c,
		baseUrl:            baseUrl,
		bearerToken:        bearerToken,
		tokenSourceKey:     tokenSourceKey,
		tokenIdentifierKey: tokenIdentifierKey,
		idTokenKey:         idTokenKey,
	}
	if u, err := url.Parse(baseUrl); err == nil {
		a.host = u.Host
	}
	return a
}

func (a *userHttp) RegisterUser(ctx context.Context, user dto.User) (dto.User, error) {

	req := common.HttpBody{
		Count:    1,
		Document: &user,
	}

	resUser := dto.User{}
	res := common.HttpBody{
		Document: &resUser,
	}
	err := a.client.RequestPostDecodeContext(ctx, "/v1/user", nil, &req, &res)
	if err != nil {
		if se, ok := err.(*sihttp.SiHttpError); ok {
			if se.Status == common.StatusTryRefreshIDToken {
				return dto.NilUser, common.ErrTokenExpired
			}
		}
		return dto.NilUser, err
	}

	if res.Status != http.StatusOK {
		return resUser, errors.New(res.Message)
	}

	return resUser, nil
}

func (a *userHttp) FindByLoginID(ctx context.Context, loginSource, tokenIdentifier, idToken string) (dto.User, error) {

	resUser := dto.User{}
	res := common.HttpBody{
		Document: &resUser,
	}
	header := make(http.Header)
	header[a.tokenSourceKey] = []string{loginSource}
	header[a.tokenIdentifierKey] = []string{tokenIdentifier}
	header[a.idTokenKey] = []string{idToken}
	err := a.client.RequestGetDecodeContext(ctx, "/v1/user/account", header, nil, &res)
	if err != nil {
		if se, ok := err.(*sihttp.SiHttpError); ok {
			if se.Status == common.StatusTryRefreshIDToken {
				return dto.NilUser, common.ErrTokenExpired
			}
		}
		return dto.NilUser, err
	}

	if res.Status != http.StatusOK {
		return resUser, errors.New(res.Message)
	}

	return resUser, nil
}
