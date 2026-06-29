package main 

import "net/http"

//a http handler that the go standard library expects
func errorHandler(w http.ResponseWriter, r *http.Request) {
	respondwithError(w, 400, "Something Went Wrong")
	
}