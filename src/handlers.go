package main

import "net/http"

//ApisIndex apis index pages
func ApisIndex(w http.ResponseWriter, r *http.Request) {
	respM := responseMessage{
		Message: "OK",
		Status:  http.StatusOK,
	}
	respM.Write(w)
}
