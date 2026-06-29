package main 

import "net/http"

//a http handler that the go standard library expects
func readinessHandler(w http.ResponseWriter, r *http.Request) {
	respondwithJson(w, 200, struct{}{})
	
}