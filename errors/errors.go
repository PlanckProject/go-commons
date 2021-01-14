package errors

import "fmt"

var showStatusCodeInError bool

// HttpError interface gives error and status code
type HttpError interface {
	error
	StatusCode() uint

	WithStatusCode(uint) HttpError
}

// New returns a new HttpError
func New(errorString string) HttpError {
	return &errorImpl{errorString: errorString, statusCode: 500}
}

// Errorf returns a HttpError with status code 500, format, and arguments
func Errorf(format string, args ...interface{}) HttpError {
	return &errorImpl{errorString: fmt.Sprintf(format, args...), statusCode: 500}
}

// ErrorfWithStatusCode returns a HttpError with status code, format, and arguments
func ErrorfWithStatusCode(statusCode uint, format string, args ...interface{}) HttpError {
	return &errorImpl{errorString: fmt.Sprintf(format, args...), statusCode: statusCode}
}

// ShowStatusCodeInError enables showing status code package wide.
// Status code is shown in error message only if it's set
func ShowStatusCodeInError(value bool) {
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

func (e *errorImpl) WithStatusCode(statusCode uint) HttpError {
	if statusCode < 100 || statusCode >= 600 {
		panic("Invalid status code")
	}
	e.statusCode = statusCode
	return e
}
