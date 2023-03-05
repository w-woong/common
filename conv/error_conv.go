package conv

import (
	"encoding/json"
	"net/http"

	"github.com/go-wonk/si/v2/sihttp"
	"github.com/w-woong/common/dto"
	"golang.org/x/oauth2"
)

func OAuth2ErrorInvalidRequest(err error) *dto.OAuth2Error {
	return OAuth2Error(err, dto.InvalidRequestErrorType, http.StatusUnauthorized)
}

func OAuth2Error(err error, defaultError dto.OAuth2ErrorType, defaultStatusCode int) *dto.OAuth2Error {
	if err == nil {
		return nil
	}

	switch t := err.(type) {
	case *dto.OAuth2Error:
		return t
	case *sihttp.Error:
		statusCode := t.GetStatusCode(defaultStatusCode)
		oerr := dto.OAuth2Error{
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
		oerr := dto.OAuth2Error{
			ErrorType:        defaultError,
			ErrorDescription: err.Error(),
			StatusCode:       &statusCode,
		}
		json.Unmarshal(t.Body, &oerr) // ignore unmarshal error
		return &oerr
	}

	return dto.NewOAuth2Error(defaultError, err.Error(), defaultStatusCode)
}

func OAuth2ErrorTryReauthenticate(err error) *dto.OAuth2Error {
	oerr := OAuth2Error(err, dto.InvalidRequestErrorType, http.StatusUnauthorized)
	oerr.SetTryReauthenticate(true)
	return oerr
}

func OAuth2ErrorTryRefresh(err error) *dto.OAuth2Error {
	oerr := OAuth2Error(err, dto.InvalidRequestErrorType, http.StatusUnauthorized)
	oerr.SetTryRefresh(true)
	return oerr
}
