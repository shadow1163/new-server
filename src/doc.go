package main

import "net/http"

// responseMessage is http response include message
//
//swagger:response responseMessage
type responseMessage struct {
	// Status
	// Message
	//
	// in: body
	Status  int    `json:"Status"`
	Message string `json:"Message"`
}

func (r *responseMessage) Write(w http.ResponseWriter) {
	w.WriteHeader(r.Status)
	w.Write([]byte(r.Message))
}

// swagger:route GET / Index
//
// get index page
//
// This will show the index page
//
//     Responses:
//       200: responseMessage
