package adapter

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/go-wonk/si/v2/sicore"
	"github.com/go-wonk/si/v2/sihttp"
	"github.com/w-woong/common"
	"github.com/w-woong/common/dto"
	"github.com/w-woong/common/port"
)

var _ port.UserSvc = (*userHttp)(nil)

type userHttp struct {
	client  *sihttp.Client
	baseUrl string
	host    string
}

func NewUserHttp(client *http.Client, baseUrl string) *userHttp {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json; charset=utf-8"

	c := sihttp.NewClient(client, sihttp.WithBaseUrl(baseUrl),
		// sihttp.WithRequestOpt(sihttp.WithBearerToken(bearerToken)),
		sihttp.WithWriterOpt(sicore.SetJsonEncoder()),
		sihttp.WithReaderOpt(sicore.SetJsonDecoder()),
		sihttp.WithDefaultHeaders(headers))

	a := &userHttp{
		client:  c,
		baseUrl: baseUrl,
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
	err := a.client.PostDecodeContext(ctx, "/v1/user", nil, &req, &res)
	if err != nil {
		return dto.NilUser, err
	}

	if res.Status != http.StatusOK {
		return resUser, errors.New(res.Message)
	}

	return resUser, nil
}

func (a *userHttp) FindByIDToken(ctx context.Context, idToken string) (dto.User, error) {

	resUser := dto.User{}
	res := common.HttpBody{
		Document: &resUser,
	}
	header := make(http.Header)
	header["Authorization"] = []string{"Bearer " + idToken}
	err := a.client.GetDecodeContext(ctx, "/v1/user/account", header, nil, &res)
	if err != nil {
		if se, ok := err.(*sihttp.Error); ok {
			val := false
			oerr := common.OAuth2Error{
				TryRefresh: &val,
			}
			json.Unmarshal(se.Body, &oerr)

			if *oerr.TryRefresh {
				return dto.NilUser, common.ErrTokenExpired
			}
		}
		return dto.NilUser, err
	}

	return resUser, nil
}
