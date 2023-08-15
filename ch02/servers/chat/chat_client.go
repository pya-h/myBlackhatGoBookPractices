package main

import (
	"fmt"
	"log"
	"net"
	"io"
	"os"
)

func transmit(connection *net.Conn, chat_channel chan string) {
	for {
		(*connection).Write([]byte(<-chat_channel))
	}
}
func chat(chat_channel chan string) {
	bytes := make([]byte, 1024)
	for {
		fmt.Print("> ")
		os.Stdin.Read(bytes)
		chat_channel <- string(bytes)
	}
}

func main() {

	if connection, err := net.Dial("tcp", "0.0.0.0:8080"); err == nil {
		defer connection.Close()
		chat_channel := make(chan string)
		go chat(chat_channel)
		go transmit(&connection, chat_channel)
		for {
			bytes := make([]byte, 1024)
			if _, err := connection.Read(bytes); err == nil {
				log.Println("Server:>", string(bytes))
			} else if err == io.EOF {
				log.Fatalln("Server down!", err)
			} else {
				log.Println("Error while reading data from server!", err)
			}
		}
	} else {
		log.Fatalln("Connection got fucked:", err)
	}
}