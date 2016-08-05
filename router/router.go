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
)

// Init sets all the routes of the app
func Init() *mux.Router {

	r := mux.NewRouter().StrictSlash(true)

	return r
}
