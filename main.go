package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

var logger = log.Default()

func main() {
	logger.SetFlags(log.Ldate | log.Lmicroseconds | log.Ltime)
	listener, err := net.Listen("tcp", ":8080")
	log.Printf("Listening to address %s\n", listener.Addr())

	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
	}
	for {
		logger.Printf("Accepting new connection...")
		// Here we have succesfully listen to the connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
		}
		// Here we have all the info in the connection request
		go handleconnection(conn)
	}
}

func handleconnection(conn net.Conn) {
	logger.Printf("Handling connection...")
	defer conn.Close()
	conn.Read(make([]byte, 1024))
	// Here we read the the data and we store it in an arry of 1024
	response := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/plain\r\n" +
		"Content-Length: 12\r\n" +
		"\r\n" +
		"Hola mundo!"
	conn.Write([]byte(response))
}
