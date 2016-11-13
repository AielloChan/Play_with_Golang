package main

import (
	"net/http"
	"log"
)

func main() {
    http.HandleFunc("/",func (w http.ResponseWriter, r *http.Request)  {
        // fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
        w.Write([]byte("nihao"))
    })

    log.Fatal(http.ListenAndServe(":8080", nil))
}