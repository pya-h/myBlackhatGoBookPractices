package main

import (
	"log"
	"net"
	"io"
	"fmt"
	"bufio"
)

func echo(client net.Conn) {
	defer client.Close()

	bytes := make([]byte, 1024)

	for {
		if size, err := client.Read(bytes[0:]); err == nil {
			log.Println("Received", size, "bytes of data from client:", client, "\n\t", string(bytes[:size]))
			log.Println("Writing data...")
			if _, err = client.Write(bytes[:size]); err != nil {
				log.Println("Error while echoing data to client:", client, err)
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

func bufferedEcho(client net.Conn) {
	defer client.Close()

	reader := bufio.NewReader(client)
	for {
		if str, err := reader.ReadString('\n'); err == nil {
			log.Println("Received", len(str), "bytes of data from client:", client, "\n\t", str)
			log.Println("Writing Data...")
			writer := bufio.NewWriter(client)
			if _, err := writer.WriteString(str); err != nil {
				log.Println("Unable to write data because", err)
			}
			writer.Flush()
		} else {
			log.Println("Unexpected error on client:", client, err)
			break
		}
	}
}

func bufferedEchoByCopy(client net.Conn) {

	if written, err := io.Copy(client, client); err == nil {
		log.Println("Successfully echoed", written, "bytes to client:", client)
	} else {
		log.Println("Something went wrong while echoing!", err)
	}
}
func main() {
	const PORT int16 = 8080
	if server, err := net.Listen("tcp", fmt.Sprintf(":%d", PORT)); err == nil {
		fmt.Printf("Listening on 0.0.0.0:%d...\n", PORT)

		for {
			if connection, err := server.Accept(); err == nil {
				fmt.Println("New Client:", connection)
				// go echo(connection)
				// go bufferedEcho(connection)
				go bufferedEchoByCopy(connection)
			} else {
				log.Println("Error while accepting a client:", err)
			}
		}
	} else {
		log.Fatalln("Cannot start the server at", PORT, err)
	}
}