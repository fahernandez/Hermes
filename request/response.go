package request

import (
	"encoding/json"
	"net/http"
)

type responseError struct {
	Code string `json:"code"`
}

// Response contains the data that will be json encoded and send to the client
type Response struct {
	Res  string          `json:"response"`
	Err  []responseError `json:"errors,omitempty"`
	Data interface{}     `json:"data,omitempty"`
}

// GetErrorResponse represents a json encoded http error response
func GetErrorResponse(m string, w http.ResponseWriter) ([]byte, error) {
	re := &responseError{m}
	r, err := json.Marshal(&Response{Res: "ERROR", Err: []responseError{*re}})
	if err != nil {
		return r, err
	}

	return getJSONResponse(r, w)
}

// GetOKResponse represents a json encoded http response
// optionally sends data to the
func GetOKResponse(d interface{}, w http.ResponseWriter) ([]byte, error) {
	r, err := json.Marshal(&Response{Res: "OK", Data: d})
	if err != nil {
		return nil, err
	}

	return getJSONResponse(r, w)
}

// getJSONResponse adds the correct content type header
// and returns the data
func getJSONResponse(r []byte, w http.ResponseWriter) ([]byte, error) {
	w.Header().Set("Content-Type", "application/json")

	return r, nil
}
