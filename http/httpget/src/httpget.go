package main

import (
	"net/http"
	"fmt"
	"os"
	"io"
)

func main()  {
    if len(os.Args) != 2 {
        fmt.Println("USAGE: %s url", os.Args[0]);
        os.Exit(1);
    }

    resp, err := http.Get(os.Args[1]);
    if err != nil {
        fmt.Println("Error");
        os.Exit(1)
    }

    defer resp.Body.Close()

    io.Copy(os.Stdout, resp.Body)
}