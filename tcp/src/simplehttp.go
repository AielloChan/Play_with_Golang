package main

import (
	"os"
	"fmt"
	"net"
	"bytes"
	"io"
)

func main()  {
    if len(os.Args) != 2 {
        fmt.Println(os.Stderr, "USAGE: %s host:port", os.Args[0])
        os.Exit(1)
    }

    service := os.Args[1] 
    conn, err := net.Dial("tcp", service)
    checkErr(err)

    _, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
    checkErr(err)

    var result []byte
    result, err = readFully(conn)
    checkErr(err)

    fmt.Println(string(result))

    os.Exit(0)
}

func checkErr(err error)  {
    if err != nil {
        fmt.Println("Fatal error: ", err.Error);
        os.Exit(1)
    }
}

func readFully(conn net.Conn) ([]byte, error) {
    defer conn.Close()

    result := bytes.NewBuffer(nil)
    var buf [512]byte
    for{
        n, err := conn.Read(buf[:])
        result.Write(buf[:n])
        if err != nil {
            if err == io.EOF{
                break
            }
            return nil, err
        }
    }

    return result.Bytes(), nil
}