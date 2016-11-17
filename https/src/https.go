package main

import (
	"fmt"
	"net/http"
)

const (
	ServerDomain     = "localhost"
	ServerPort       = 8080
	ResponseTemplate = "Hello"
)

func main() {
	http.HandleFunc(fmt.Sprintf("%s:%d/", ServerDomain, ServerPort), rootHandler)
	if err := http.ListenAndServeTLS(fmt.Sprintf(":%d", ServerPort), "server.crt", "server.key", nil); err != nil {
		fmt.Println(err.Error())
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Content-Length", fmt.Sprint(len(ResponseTemplate)))
	w.Write([]byte(ResponseTemplate))
}
