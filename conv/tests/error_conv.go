package conv_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/go-wonk/si/v2/sihttp"
	"github.com/stretchr/testify/assert"
	"github.com/w-woong/common"
	"github.com/w-woong/common/conv"
	"golang.org/x/oauth2"
)

func Test_InvalidRequestError_sihttpError(t *testing.T) {

	err := conv.OAuth2ErrorInvalidRequest(errors.New("test error"))

	statusCode := http.StatusUnauthorized
	expected := &common.OAuth2Error{
		ErrorType:        common.InvalidRequestErrorType,
		ErrorDescription: "test error",
		StatusCode:       &statusCode,
	}
	assert.EqualValues(t, expected, err)

	serr := &sihttp.Error{
		Response: &http.Response{
			Status:     http.StatusText(statusCode),
			StatusCode: statusCode,
		},
		Body: nil,
	}
	err = conv.OAuth2ErrorInvalidRequest(serr)
	expected = &common.OAuth2Error{
		ErrorType:        common.InvalidRequestErrorType,
		ErrorDescription: "status: " + http.StatusText(statusCode),
		StatusCode:       &statusCode,
	}
	assert.EqualValues(t, expected, err)

	serr = &sihttp.Error{
		Response: &http.Response{
			Status:     http.StatusText(statusCode),
			StatusCode: statusCode,
		},
		Body: []byte("error"),
	}
	err = conv.OAuth2ErrorInvalidRequest(serr)
	expected = &common.OAuth2Error{
		ErrorType:        common.InvalidRequestErrorType,
		ErrorDescription: "status: Unauthorized, body: error",
		StatusCode:       &statusCode,
	}
	assert.EqualValues(t, expected, err)

	serr = &sihttp.Error{
		Response: &http.Response{
			Status:     http.StatusText(statusCode),
			StatusCode: statusCode,
		},
		Body: []byte(`{"error":"invalid_grant","error_description":"missing scope","status_code":403}`),
	}
	err = conv.OAuth2ErrorInvalidRequest(serr)
	statusCode = 403
	expected = &common.OAuth2Error{
		ErrorType:        common.InvalidGrantErrorType,
		ErrorDescription: "missing scope",
		StatusCode:       &statusCode,
	}
	assert.EqualValues(t, expected, err)
}

func Test_InvalidRequestError_oauth2RetrieveError(t *testing.T) {

	err := conv.OAuth2ErrorInvalidRequest(errors.New("test error"))

	statusCode := http.StatusUnauthorized
	expected := &common.OAuth2Error{
		ErrorType:        common.InvalidRequestErrorType,
		ErrorDescription: "test error",
		StatusCode:       &statusCode,
	}
	assert.EqualValues(t, expected, err)

	serr := &oauth2.RetrieveError{
		Response: &http.Response{
			Status:     http.StatusText(statusCode),
			StatusCode: statusCode,
		},
		Body: nil,
	}
	err = conv.OAuth2ErrorInvalidRequest(serr)
	expected = &common.OAuth2Error{
		ErrorType:        common.InvalidRequestErrorType,
		ErrorDescription: "oauth2: cannot fetch token: Unauthorized\nResponse: ",
		StatusCode:       &statusCode,
	}
	assert.EqualValues(t, expected, err)

	serr = &oauth2.RetrieveError{
		Response: &http.Response{
			Status:     http.StatusText(statusCode),
			StatusCode: statusCode,
		},
		Body: []byte("error"),
	}
	err = conv.OAuth2ErrorInvalidRequest(serr)
	expected = &common.OAuth2Error{
		ErrorType:        common.InvalidRequestErrorType,
		ErrorDescription: "oauth2: cannot fetch token: Unauthorized\nResponse: error",
		StatusCode:       &statusCode,
	}
	assert.EqualValues(t, expected, err)

	serr = &oauth2.RetrieveError{
		Response: &http.Response{
			Status:     http.StatusText(statusCode),
			StatusCode: statusCode,
		},
		Body: []byte(`{"error":"invalid_grant","error_description":"missing scope","status_code":403}`),
	}
	err = conv.OAuth2ErrorInvalidRequest(serr)
	statusCode = 403
	expected = &common.OAuth2Error{
		ErrorType:        common.InvalidGrantErrorType,
		ErrorDescription: "missing scope",
		StatusCode:       &statusCode,
	}
	assert.EqualValues(t, expected, err)
}
