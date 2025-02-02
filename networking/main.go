package main

import (
	"bufio"
	"context"
	"log"
	"net"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Error occurred while listening: ", err)
	}
	defer lis.Close()

	log.Println("Server is listening on port 8080")

	var wg sync.WaitGroup

	go func() {
		<-ctx.Done()
		log.Println("Shutting down server...")
		lis.Close()
		wg.Wait()
		log.Println("Server shut down gracefully")
	}()

	for {
		conn, err := lis.Accept()
		if err != nil {
			select {
			case <-ctx.Done():
				return
			default:
				log.Println("Error occurred while accepting connection: ", err)
				continue
			}
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			handleConnection(conn)
		}()
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	log.Println("Connection accepted from", conn.RemoteAddr())

	reader := bufio.NewReader(conn)
	for {
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		message, err := reader.ReadString('\n')
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				log.Println("Read timeout:", err)
				return
			}
			log.Println("Error reading from connection:", err)
			return
		}

		log.Printf("Received message: %s", message)
		conn.Write([]byte("Echo: " + message))
	}
}
