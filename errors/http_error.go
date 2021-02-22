package errors

import "fmt"

var showStatusCodeInError bool

// HTTPError interface gives error and status code
type HTTPError interface {
	error
	StatusCode() uint

	WithStatusCode(uint) HTTPError
}

// New returns a new HTTPError
func NewHTTPError(errorString string) HTTPError {
	return &errorImpl{errorString: errorString, statusCode: 500}
}

// Errorf returns a HTTPError with status code 500, format, and arguments
func HTTPErrorf(format string, args ...interface{}) HTTPError {
	return &errorImpl{errorString: fmt.Sprintf(format, args...), statusCode: 500}
}

// ErrorfWithStatusCode returns a HTTPError with status code, format, and arguments
func HTTPErrorfWithStatusCode(statusCode uint, format string, args ...interface{}) HTTPError {
	return &errorImpl{errorString: fmt.Sprintf(format, args...), statusCode: statusCode}
}

// ShowStatusCodeInError enables showing status code package wide.
// Status code is shown in error message only if it's set
func ShowStatusCodeInHTTPError(value bool) {
	showStatusCodeInError = value
}

type errorImpl struct {
	statusCode  uint
	errorString string
}

func (e *errorImpl) Error() string {
	if showStatusCodeInError && e.statusCode != 0 {
		return fmt.Sprintf("%d: %s", e.statusCode, e.errorString)
	}
	return e.errorString
}

func (e *errorImpl) StatusCode() uint {
	return e.statusCode
}

func (e *errorImpl) WithStatusCode(statusCode uint) HTTPError {
	if statusCode < 100 || statusCode >= 600 {
		panic("Invalid status code")
	}
	e.statusCode = statusCode
	return e
}
