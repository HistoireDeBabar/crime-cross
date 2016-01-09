package mocks

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type ErrorTransport struct {
}

func (t *ErrorTransport) RoundTrip(request *http.Request) (response *http.Response, err error) {
	return nil, errors.New("something went wrong")
}

type ResponseErrorStatusCode struct {
	StatusCode int
}

func (t *ResponseErrorStatusCode) RoundTrip(request *http.Request) (response *http.Response, err error) {
	reader := strings.NewReader("")
	readerCloser := ioutil.NopCloser(reader)
	response = &http.Response{
		Body:       readerCloser,
		StatusCode: t.StatusCode,
	}
	return response, nil
}

type SuccessForceListNoResponseBody struct {
	StatusCode int
}

func (t *SuccessForceListNoResponseBody) RoundTrip(request *http.Request) (response *http.Response, err error) {
	reader := strings.NewReader("")
	readerCloser := ioutil.NopCloser(reader)
	response = &http.Response{
		Body:       readerCloser,
		StatusCode: t.StatusCode,
	}
	return response, nil
}

type SuccessForceListWithResponseBody struct {
	StatusCode int
}

func (t *SuccessForceListWithResponseBody) RoundTrip(request *http.Request) (response *http.Response, err error) {
	reader := strings.NewReader("[{\"id\": \"south_shields\", \"name\": \"South Shields\"}]")
	readerCloser := ioutil.NopCloser(reader)
	response = &http.Response{
		Body:       readerCloser,
		StatusCode: t.StatusCode,
	}
	return response, nil
}
