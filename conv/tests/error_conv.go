package conv_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/go-wonk/si/v2/sihttp"
	"github.com/stretchr/testify/assert"
	"github.com/w-woong/common/conv"
	"github.com/w-woong/common/dto"
	"golang.org/x/oauth2"
)

func Test_InvalidRequestError_sihttpError(t *testing.T) {

	err := conv.OAuth2ErrorInvalidRequest(errors.New("test error"))

	statusCode := http.StatusUnauthorized
	expected := &dto.OAuth2Error{
		ErrorType:        dto.InvalidRequestErrorType,
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
	expected = &dto.OAuth2Error{
		ErrorType:        dto.InvalidRequestErrorType,
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
	expected = &dto.OAuth2Error{
		ErrorType:        dto.InvalidRequestErrorType,
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
	expected = &dto.OAuth2Error{
		ErrorType:        dto.InvalidGrantErrorType,
		ErrorDescription: "missing scope",
		StatusCode:       &statusCode,
	}
	assert.EqualValues(t, expected, err)
}

func Test_InvalidRequestError_oauth2RetrieveError(t *testing.T) {

	err := conv.OAuth2ErrorInvalidRequest(errors.New("test error"))

	statusCode := http.StatusUnauthorized
	expected := &dto.OAuth2Error{
		ErrorType:        dto.InvalidRequestErrorType,
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
	expected = &dto.OAuth2Error{
		ErrorType:        dto.InvalidRequestErrorType,
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
	expected = &dto.OAuth2Error{
		ErrorType:        dto.InvalidRequestErrorType,
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
	expected = &dto.OAuth2Error{
		ErrorType:        dto.InvalidGrantErrorType,
		ErrorDescription: "missing scope",
		StatusCode:       &statusCode,
	}
	assert.EqualValues(t, expected, err)
}
