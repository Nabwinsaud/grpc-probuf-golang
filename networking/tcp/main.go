package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

type Server struct {
	Addr    string
	message string
}

func NewServer(addr string) *Server {
	return &Server{Addr: addr}
}

func (s *Server) Start() {
	ln, err := net.Listen("tcp", s.Addr)
	fmt.Println("Server is listening on port", s.Addr)

	if err != nil {
		log.Fatal("Error occurred while listening: ", err)
	}

	for {
		con, err := ln.Accept()
		if err != nil {
			log.Println("Error occurred while accepting connection: ", err)
			continue
		}
		go handleConnection(con)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	log.Println("Connection accepted from", conn.RemoteAddr())

	reader := bufio.NewReader(conn)
	for {

		message, err := reader.ReadString('\n')

		if err == io.EOF {
			log.Println("Connection closed by client", conn.RemoteAddr().String())
			return
		}

		if err != nil {
			log.Println("Error occurred while reading message: ", err)
			return
		}
		log.Println("Received message: ", message)
		conn.Write([]byte("You said: " + message))
	}
}

func main() {
	srv := NewServer(":4000")
	srv.Start()
}
