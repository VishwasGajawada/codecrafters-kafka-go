package main

import (
	"fmt"
	"net"
	"os"
)

func handleConnection(connection net.Conn) {
	defer connection.Close()

	buffer := make([]byte, 1024)
	_, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading from connection: ", err.Error())
		os.Exit(1)
	}

	correlationID_bytes := buffer[8:12]

	response := make([]byte, 8)              // 4 bytes message, 4 bytes correlation ID
	copy(response[0:4], []byte{0, 0, 0, 0})  // Placeholder for message
	copy(response[4:8], correlationID_bytes) // Copy correlation ID

	// Simulate processing the request
	connection.Write(response)

}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:9092")
	if err != nil {
		fmt.Println("Failed to bind to port 9092")
		os.Exit(1)
	}
	defer l.Close()
	for {

		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		// send a message to the client
		go handleConnection(conn)
	}

}
