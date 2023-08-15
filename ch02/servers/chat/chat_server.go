package main

import (
	"log"
	"net"
	"io"
	"fmt"
)

var clients []net.Conn
func broadcast(id int, client net.Conn) {
	defer client.Close()

	bytes := make([]byte, 1024)

	for {
		if size, err := client.Read(bytes); err == nil {
			received := string(bytes[:size])
			log.Println("Received", size, "bytes of data from client:", client, "\n\t", received)
			log.Println("Writing data...")
			for _, c := range clients {
				message := fmt.Sprintf("Client#%d: %s", id, received)
				if _, err = c.Write([]byte(message)); err != nil {
					log.Println("Error while echoing data to client:", c, err)
				}
			}
		} else if err == io.EOF {
			log.Println("Client", client, "disconnected.")
			break
		} else {
			log.Println("Unexpected error on client:", client, err)
			break
		}
		
	}
}
func main() {
	const PORT int16 = 8080
	if server, err := net.Listen("tcp", fmt.Sprintf(":%d", PORT)); err == nil {
		fmt.Printf("Listening on 0.0.0.0:%d...\n", PORT)

		for {
			if connection, err := server.Accept(); err == nil {
				clients = append(clients, connection)
				id := len(clients)
				fmt.Println("New Client:", connection)
				go broadcast(id, connection)
			} else {
				log.Fatalln("Error while accepting a client:", err)
			}
		}
	} else {
		log.Fatalln("Cannot start the server at", PORT, err)
	}
}