//                    _
//  _             _ _| |_
// | |           | |_   _|
// | |___  _   _ | | |_|
// | '_  \| | | || | | |
// | | | || |_| || | | |
// |_| |_|\___,_||_| |_|
//
// (c) Huli Inc

package handler

import (
	"net/http"
	"github.com/fahernandez/hermes/request"
)

// Handler contains the app config and the Handle function used by the
// controllers to handle the requests
type Handler struct {
	*Context
	Handle func(c *Context, w http.ResponseWriter, r *http.Request) error
}

// ServeHTTP allows Handler to satisfy the http.Handler interface
// It will call the controller handler function passing along the Context
// It will also handle errors using the handler.Error type
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get error from context here?
	err := h.Handle(h.Context, w, r)
	if err != nil {
		request.HandleError(err, w)
	}
}
