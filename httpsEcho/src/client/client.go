package main

import (
	"crypto/tls"
	"io"
	"log"
)

func main() {
	conn, err := tls.Dial("tcp", "localhost:8000", nil)
	if err != nil {
		log.Fatalf("Client: dial: %s", err)
	}
	defer conn.Close()
	log.Println("Client: connected to: ", conn.RemoteAddr())

	state := conn.ConnectionState()
	log.Println("Client handshake: ", state.HandshakeComplete)
	log.Println("Client: mutual: ", state.NegotiatedProtocolIsMutual)

	message := "Hello\n"
	n, err := io.WriteString(conn, message)
	if err != nil {
		log.Fatalf("Client: write: %s", err)
	}
	log.Printf("Client: wrote %q (%d bytes)\n", message, n)

	reply := make([]byte, 256)
	n, err = conn.Read(reply)
	log.Printf("Client: read %q (%d bytes)\n", string(reply), n)
	log.Print("Client: exiting")
}
