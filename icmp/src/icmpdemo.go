package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("USAGE: ", os.Args[0], "host")
		os.Exit(1)
	}

	service := os.Args[1]
	fmt.Println(1)

	conn, err := net.Dial("ip4:icmp", service)
	checkErr(err)
	fmt.Println(2)

	var msg [512]byte
	msg[0] = 8
	msg[1] = 0
	msg[2] = 0
	msg[3] = 0
	msg[4] = 0
	msg[5] = 13
	msg[6] = 0
	msg[7] = 37
	len := 8

	check := checkSum(msg[:len])
	msg[2] = byte(check >> 8)
	msg[3] = byte(check & 255)

	fmt.Println(3)
	_, err = conn.Write(msg[0:len])
	checkErr(err)

	fmt.Println(4)
	_, err = conn.Read(msg[0:])
	checkErr(err)

	fmt.Println("Got response")
	if msg[5] == 13 {
		fmt.Println("Identifier matches")
	}
	if msg[7] == 37 {
		fmt.Println("Sequence matches")
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(os.Stderr, "Fatal error: %s", err.Error)
		os.Exit(1)
	}
}

func checkSum(msg []byte) uint16 {
	sum := 0

	for n := 1; n < len(msg)-1; n += 2 {
		sum += int(msg[n])*256 + int(msg[n+1])
	}

	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)
	var answer uint16 = uint16(^sum)
	return answer
}
