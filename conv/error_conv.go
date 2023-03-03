package conv

import (
	"encoding/json"
	"net/http"

	"github.com/go-wonk/si/v2/sihttp"
	"github.com/w-woong/common"
	"golang.org/x/oauth2"
)

func OAuth2ErrorInvalidRequest(err error) *common.OAuth2Error {
	return OAuth2Error(err, common.InvalidRequestErrorType, http.StatusUnauthorized)
}

func OAuth2Error(err error, defaultError common.OAuth2ErrorType, defaultStatusCode int) *common.OAuth2Error {
	if err == nil {
		return nil
	}

	switch t := err.(type) {
	case *common.OAuth2Error:
		return t
	case *sihttp.Error:
		statusCode := t.GetStatusCode(defaultStatusCode)
		oerr := common.OAuth2Error{
			ErrorType:        defaultError,
			ErrorDescription: t.Error(),
			StatusCode:       &statusCode,
		}
		if t.Body == nil {
			return &oerr
		}
		json.Unmarshal(t.Body, &oerr)
		return &oerr
	case *oauth2.RetrieveError:
		statusCode := t.Response.StatusCode
		oerr := common.OAuth2Error{
			ErrorType:        defaultError,
			ErrorDescription: err.Error(),
			StatusCode:       &statusCode,
		}
		json.Unmarshal(t.Body, &oerr) // ignore unmarshal error
		return &oerr
	}

	return common.NewOAuth2Error(defaultError, err.Error(), defaultStatusCode)
}

func OAuth2ErrorTryReauthenticate(err error) *common.OAuth2Error {
	oerr := OAuth2Error(err, common.InvalidRequestErrorType, http.StatusUnauthorized)
	oerr.SetTryReauthenticate(true)
	return oerr
}

func OAuth2ErrorTryRefresh(err error) *common.OAuth2Error {
	oerr := OAuth2Error(err, common.InvalidRequestErrorType, http.StatusUnauthorized)
	oerr.SetTryRefresh(true)
	return oerr
}
