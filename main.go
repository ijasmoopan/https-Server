package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", HelloHandler)
	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatalln("Server couldn't start-", err)
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello world"))
}