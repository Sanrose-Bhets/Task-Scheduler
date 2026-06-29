package main

import (
	"net/http"
	"log"
	"encoding/json"
)

func respondwithJson(w http.ResponseWriter, code int, payload interface{}){
	data,err:= json.Marshal(payload)
	if err != nil{
		log.Println("Failed to Marshal the JSON payload:", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func respondwithError(w http.ResponseWriter, code int, msg string){
	if code>499{
		log.Println("THE FUCK ARE YOU DOING MAN, THIS A 5XX ERROR MAN",msg)
	}
	type errResponse struct{
		Error string `json:"error"`
	}
	respondwithJson(w,code,errResponse{Error:msg})

}