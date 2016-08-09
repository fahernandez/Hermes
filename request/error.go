package request

import (
	"fmt"
	"log"
	"net/http"
)

// Error interface contains the error and has a method to retrieve the http
// status code
type Error interface {
	error
	StatusCode() int
	Message() []byte
}

// HTTPError represents a HTTP 4xx/5xx error
type HTTPError struct {
	Code int
	Err  error
	Mes  []byte
}

// Allows HTTPError to satisfy the error interface.
func (ce HTTPError) Error() string {
	return ce.Err.Error()
}

// StatusCode returns our HTTP status code.
func (ce HTTPError) StatusCode() int {
	return ce.Code
}

// Message returns the user friendly error message
func (ce HTTPError) Message() []byte {
	return ce.Mes
}

// HandleError will recieve an error message and write the message
// to the logs and client.
func HandleError(err error, w http.ResponseWriter) {
	switch e := err.(type) {
	case HTTPError:
		log.Printf("HTTP %d - %s", e.StatusCode(), e)

		// set the default content type header if not set
		if len(w.Header().Get("Content-Type")) == 0 {
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		}
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(e.StatusCode())
		fmt.Fprintln(w, string(e.Message()))
	default:
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
	}
}

// ClientError returns an error due to a client problem (codes 4xx)
func ClientError(e error, mes string, code int, w http.ResponseWriter) Error {
	res, err := GetErrorResponse(mes, w)
	if err != nil {
		return ServerError(err)
	}

	return HTTPError{
		Code: code,
		Err:  e,
		Mes:  res,
	}
}

// ServerError returns an error due to a server issue (codes 5xx)
func ServerError(err error) Error {
	return HTTPError{
		Code: http.StatusInternalServerError,
		Err:  err,
		Mes:  []byte(http.StatusText(http.StatusInternalServerError)),
	}
}
