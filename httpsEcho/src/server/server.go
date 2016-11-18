package main

import (
	"crypto/tls"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	cert, err := tls.LoadX509KeyPair("../server.crt", "../server.key")
	if err != nil {
		log.Fatalf("Server: loadkeys %s", err)
	}

	config := tls.Config{Certificates: []tls.Certificate{cert}}
	config.Time = time.Now

	service := "localhost:8000"

	listener, err := tls.Listen("tcp", service, &config)
	if err != nil {
		log.Fatalf("Server: loadkeys %s", err)
	}

	log.Println("Server: listening")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Server: accept: %s", err)
		}

		log.Printf("Server: accept from %s", conn.RemoteAddr())
		go handlerClient(conn)
	}
}

func handlerClient(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 512)
	for {
		log.Print("Server: conn: waiting")
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Printf("Server: conn: read: %s", err)
			}
			break
		}
		log.Printf("Server: conn: echo %q\n", string(buf[:n]))

		if err != nil {
			log.Printf("Server: write: %s", err)
			break
		}
	}
	log.Println("Server: conn: closed")
}
