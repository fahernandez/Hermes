//                    _
//  _             _ _| |_
// | |           | |_   _|
// | |___  _   _ | | |_|
// | '_  \| | | || | | |
// | | | || |_| || | | |
// |_| |_|\___,_||_| |_|
//
// (c) Huli Inc
package router

import (
	"github.com/gorilla/mux"
	"github.com/fahernandez/hermes/handler/controller"
)

// Init sets all the routes of the app
func Init() *mux.Router {

	routes := mux.NewRouter().StrictSlash(true)

	// Message send controller
	routes.Path("/send").
		HandlerFunc(controller.Send).
		Methods("POST").
		Name("Send")

	return routes
}
