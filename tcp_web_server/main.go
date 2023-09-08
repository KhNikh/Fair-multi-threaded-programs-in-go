package main

import (
	"log"
	"net"
	"time"
)

func handelIncomingRequest(conn net.Conn) {
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(15 * time.Second)
	_, err = conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello, World!\r\n"))

	conn.Close()
}
func main() {
	listener, err := net.Listen("tcp", ":1624")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println("Hello world", conn)
		go handelIncomingRequest(conn)
	}

}
