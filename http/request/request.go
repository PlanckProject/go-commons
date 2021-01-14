package request

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/PlankProject/go-commons/constants"
	"github.com/PlankProject/go-commons/logger"
	"go.uber.org/multierr"
)

type httpRequest struct {
	request *http.Request
	timeout time.Duration
	payload []byte
	header  map[string]string
	retries uint8
}

type byteReaderCloser struct {
	io.Reader
}

func (byteReaderCloser) Close() error { return nil }

func New() *httpRequest {
	request, err := http.NewRequest(constants.EmptyString, constants.EmptyString, nil)
	if err != nil {
		logger.Error("Failed to create a http request")
		return nil
	}
	return &httpRequest{
		request: request,
		header:  make(map[string]string),
		retries: 4, // 1 + 3 retries. Leaving for now, will need to be properly piped into a worker with retires treated as a seperate action
		timeout: 30 * time.Second,
	}
}

func (h *httpRequest) SetContext(ctx context.Context) *httpRequest {
	h.request = h.request.WithContext(ctx)
	return h
}

func (h *httpRequest) SetMethod(method string) *httpRequest {
	if method != constants.MethodGet &&
		method != constants.MethodPost &&
		method != constants.MethodPut &&
		method != constants.MethodDelete {
		logger.Errorf("Invalid/Unsupported http method: %s", method)
		return nil
	}
	h.request.Method = method
	return h
}

func (h *httpRequest) SetURI(uri string) *httpRequest {
	u, err := url.Parse(uri)
	if err != nil {
		logger.Errorf("Invalid URL %s", uri)
		return nil
	}
	h.request.URL = u
	return h
}

func (h *httpRequest) SetPayloadFromReader(reader io.ReadCloser) *httpRequest {
	h.request.Body = reader
	return h
}

func (h *httpRequest) SetPayload(payload []byte) *httpRequest {
	h.request.Body = byteReaderCloser{bytes.NewBuffer(payload)}
	h.payload = payload
	return h
}

func (h *httpRequest) SetHeader(key, value string) *httpRequest {
	h.request.Header.Set(key, value)
	h.header[key] = value
	return h
}

func (h *httpRequest) SetCookie(requestCookie *http.Cookie) *httpRequest {
	h.request.AddCookie(requestCookie)
	h.header["Cookie"] = requestCookie.String()
	return h
}

func (h *httpRequest) SetTimeout(timeout time.Duration) *httpRequest {
	h.timeout = timeout * time.Second
	return h
}

func (h *httpRequest) SetRetries(retries uint8) *httpRequest {
	h.retries = retries + 1
	return h
}

func (h *httpRequest) Do() (*http.Response, error) {
	if h.request.URL.String() == constants.EmptyString {
		return nil, fmt.Errorf("Request URI must be specified")
	}

	retries := h.retries
	client := &http.Client{Timeout: h.timeout}

	if (h.payload == nil || len(h.payload) == 0) && h.request.Body != nil {
		requestPayload, err := ioutil.ReadAll(h.request.Body)
		requestBodyReader := bytes.NewReader(requestPayload)
		h.request.Body = byteReaderCloser{requestBodyReader}
		if err != nil {
			return nil, err
		}
		h.payload = requestPayload
	}

	for retries != 0 {
		retries--

		response, err := client.Do(h.request)
		if err != nil {
			if urlError, ok := err.(*url.Error); ok {
				if urlError.Timeout() {
					logger.WithFields(getRequestFields(h.request.Method,
						h.request.URL.RequestURI(),
						string(h.payload),
						h.header,
						fmt.Errorf("Call failed at retry number %d", h.retries-retries-1))).
						Errorln("Request timed out")

					continue
				}
			} else {
				err = multierr.Append(err, fmt.Errorf("Call failed at retry number %d", h.retries-retries-1))
				logger.WithFields(getRequestFields(h.request.Method,
					h.request.URL.RequestURI(),
					string(h.payload),
					h.header,
					err)).
					Errorf("API call failed")
			}
			continue
		}

		responsePayload, err := ioutil.ReadAll(response.Body)

		if err != nil {
			return nil, err
		}

		responseBodyReader := bytes.NewReader(responsePayload)
		response.Body = byteReaderCloser{responseBodyReader}

		logFieldMap := getRequestFields(h.request.Method,
			h.request.URL.RequestURI(),
			string(h.payload),
			h.header,
			nil)
		logFieldMap["http.response.payload"] = string(responsePayload)
		logFieldMap["http.response.code"] = response.StatusCode
		logger.WithFields(logFieldMap).
			Infof("API call successful")
		return response, nil
	}

	logger.WithFields(getRequestFields(h.request.Method,
		h.request.URL.RequestURI(),
		string(h.payload),
		h.header,
		fmt.Errorf("Calls failed after %d reties", h.retries))).
		Errorf("API call failed")
	return nil, fmt.Errorf("Request failed")
}

func getRequestFields(method, uri, payload string, headers map[string]string, err error) logger.Fields {
	fields := logger.Fields{}
	if method != "" {
		fields["http.request.method"] = method
	}
	if uri != "" {
		fields["http.request.uri"] = uri
	}
	if payload != "" {
		fields["http.request.payload"] = payload
	}
	if headers != nil {
		fields["http.request.headers"] = headers
	}
	if err != nil {
		fields["error"] = err
	}
	return fields
}
