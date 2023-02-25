package adapter

import (
	"context"
	"net/http"
	"net/url"

	"github.com/go-wonk/si/v2/sicore"
	"github.com/go-wonk/si/v2/sihttp"
	"github.com/w-woong/common/dto"
)

type IDTokenHttp struct {
	client *sihttp.Client

	baseUrl     string
	host        string
	bearerToken string

	tokenIdentifierKey string
	idTokenKey         string
}

func NewIDTokenHttp(client *http.Client, baseUrl string, bearerToken string,
	tokenIdentifierKey, idTokenKey string) *IDTokenHttp {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json; charset=utf-8"

	c := sihttp.NewClient(client, sihttp.WithBaseUrl(baseUrl),
		sihttp.WithRequestOpt(sihttp.WithBearerToken(bearerToken)),
		sihttp.WithWriterOpt(sicore.SetJsonEncoder()),
		sihttp.WithReaderOpt(sicore.SetJsonDecoder()),
		sihttp.WithDefaultHeaders(headers))

	a := &IDTokenHttp{
		client:             c,
		baseUrl:            baseUrl,
		bearerToken:        bearerToken,
		tokenIdentifierKey: tokenIdentifierKey,
		idTokenKey:         idTokenKey,
	}
	if u, err := url.Parse(baseUrl); err == nil {
		a.host = u.Host
	}
	return a
}

func (a *IDTokenHttp) Refresh(ctx context.Context, tokenSource, tokenIdentifier string, idToken string) (dto.Token, error) {

	header := http.Header{}
	header[a.tokenIdentifierKey] = []string{tokenIdentifier}
	header[a.idTokenKey] = []string{idToken}

	res := dto.Token{}
	err := a.client.GetDecodeContext(ctx, "/v1/oauth2/validate/"+tokenSource, header, nil, &res)
	if err != nil {
		return dto.NilToken, err
	}

	return res, nil
}
