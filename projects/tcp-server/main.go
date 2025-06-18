package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

/**
 * A simple TCP server that listens on port 1729 and accepts incoming connections.
 * It prints the address of the connected client and closes the connection immediately.
 * To run this server, use the command: go run main.go
 * The server will listen on port 1729 and can be tested using curl http://localhost:1729
 */
func simpleTCPServer() {
	// Create a TCP listener on port 1729
	listener, err := net.Listen("tcp", ":1729")
	if err != nil {
		log.Fatal("Error starting TCP server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("TCP server is listening on port 1729...")

	// Accept incoming connections
	conn, err := listener.Accept()
	if err != nil {
		log.Println("Error accepting connection:", err)
		return
	}

	fmt.Println("Client connected:", conn.RemoteAddr().String())
	conn.Close() // Close the connection after handling it
}

/**
 * A more advanced TCP server that reads data from the client,
 * simulates processing the data, and sends a response back to the client.
 * To run this server, use the command: go run main.go
 * The server will listen on port 1729 and can be tested using curl http://localhost:1729
 * The server will read data from the client, simulate processing it,
 * and then send a response back to the client.
 * The response will be "Hello from TCP server!".
 * The server will close the connection after sending the response.
 */
func tcpServerV2() {
	// Create a TCP listener on port 1729
	listener, err := net.Listen("tcp", ":1729")
	if err != nil {
		log.Fatal("Error starting TCP server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("TCP server is listening on port 1729...")

	// Accept incoming connections
	conn, err := listener.Accept()
	if err != nil {
		log.Println("Error accepting connection:", err)
		return
	}

	do_tcpV2(conn)
}

// A more advanced TCP server that handles multiple connections sequentially.
// It reads data from the client, simulates processing it, and sends a response back to the client.
// To run this server, use the command: go run main.go
// The server will listen on port 1729 and can be tested using curl http://localhost:1729
func tcpServerV3() {
	// Create a TCP listener on port 1729
	listener, err := net.Listen("tcp", ":1729")
	if err != nil {
		log.Fatal("Error starting TCP server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("TCP server is listening on port 1729...")

	for {
		fmt.Println("Waiting for client connection...")
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			return
		}
		fmt.Println("Client connected:", conn.RemoteAddr().String())
		do_tcpV2(conn)
	}
}

/** A more advanced TCP server that handles multiple connections concurrently.
 * It reads data from the client, simulates processing it, and sends a response back to the client.
 */
func tcpServerV4() {
	// Create a TCP listener on port 1729
	listener, err := net.Listen("tcp", ":1729")
	if err != nil {
		log.Fatal("Error starting TCP server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("TCP server is listening on port 1729...")

	for {
		fmt.Println("Waiting for client connection...")
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			return
		}
		fmt.Println("Client connected:", conn.RemoteAddr().String())
		go do_tcpV2(conn)
	}
}

func tcpServerV5() {
	/**
	 * Some optimizations to consider:
	 * 1. Use a thread pool to limit the number of threads/goroutines handling connections.
	 * 2. Use a connection timeout to close idle connections.
	 * 3. TCP backlog queue configuration to handle burst traffic.
	 * 4. Use a buffered channel to queue incoming connections.
	 * 5. Implement graceful shutdown to close connections properly.
	 * 6. Use a more efficient data structure for handling connections.
	 */

}

func do_tcpV2(connection net.Conn) {
	buffer := make([]byte, 1024)
	//Read call is a syscall that blocks until data is available
	_, err := connection.Read(buffer)
	if err != nil {
		log.Fatal("Error reading from connection:", err)
		return
	}

	//process the data received from the client
	fmt.Println("Processing data from client...")
	time.Sleep(8 * time.Second) // Simulate processing delay
	//fmt.Println("Data received from client:", string(buffer))

	connection.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello from TCP server!\r\n"))
	connection.Close()
}

func main() {
	//simpleTCPServer()
	//tcpServerV2()
	//tcpServerV3()
	tcpServerV4()
	fmt.Println("Server stopped.")
}
