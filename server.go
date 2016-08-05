//                    _
//  _             _ _| |_
// | |           | |_   _|
// | |___  _   _ | | |_|
// | '_  \| | | || | | |
// | | | || |_| || | | |
// |_| |_|\___,_||_| |_|
//
// (c) Huli Inc
package main

import (
	"github.com/codegangsta/negroni"
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/fahernandez/hermes/router"
)

func main() {
	r := router.Init()

	n := negroni.New(
		negroni.NewLogger(),
		negroni.NewRecovery(),
		gzip.Gzip(gzip.DefaultCompression),
	)

	n.UseHandler(r)
	n.Run(":9001")
}