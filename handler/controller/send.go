//                    _
//  _             _ _| |_
// | |           | |_   _|
// | |___  _   _ | | |_|
// | '_  \| | | || | | |
// | | | || |_| || | | |
// |_| |_|\___,_||_| |_|
//
// (c) Huli Inc

package controller

import (
	"net/http"
	"fmt"
)


/**
 * Mesasge send function controller
 */
func Send(w http.ResponseWriter, r *http.Request) error {
	fmt.Printf("Funciono esta putada.\n")
	return nil
}
