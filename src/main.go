// Package main Todos API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http
//     Host: localhost
//     BasePath: /apis
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: shadow1163<674602286@qq.com> https://github.com/shadow1163
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"net/http"

	"github.com/shadow1163/logger"
)

var log = logger.Log

func main() {

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":80", router))
}
