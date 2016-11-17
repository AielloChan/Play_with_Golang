package main

import "net/http"

func main() {
	h := http.FileServer(http.Dir("."))
	http.ListenAndServeTLS(":8081", "server.crt", "server.key", h)
}
