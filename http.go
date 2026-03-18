package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// Handler function
func handler(conn net.Conn, path string) {

	if path == "/test" {
		response := "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\n\r\nHello from /test endpoint"
		conn.Write([]byte(response))
	} else {
		response := "HTTP/1.1 404 Not Found\r\n\r\nRoute Not Found"
		conn.Write([]byte(response))
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Println("New TCP connection established")

	reader := bufio.NewReader(conn)

	// HTTP Request Parsed
	requestLine, _ := reader.ReadString('\n')
	fmt.Println("Raw Request:", requestLine)

	parts := strings.Split(requestLine, " ")

	if len(parts) < 2 {
		return
	}

	method := parts[0]
	path := parts[1]

	fmt.Println("HTTP Method:", method)
	fmt.Println("Request Path:", path)

	// ServeMux Route Matching
	handler(conn, path)

	fmt.Println("Response written")
}

func main() {

	// TCP Listener (port binding)
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	fmt.Println("Server listening on port 8080")

	for {

		// Accept Loop
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		// New Goroutine Created
		go handleConnection(conn)
	}
}
