package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"bufio"
	"io"
)
func nonBufferedCommunication(connection net.Conn) {
	for {
		fmt.Print("You:> ")
		//var input string
		//fmt.Scanln(&input)
		input := make([]byte, 1024)
		os.Stdin.Read(input)

		if len(input) == 0 {
			break
		}
		connection.Write(input)
		bytes := make([]byte, 1024)
		if _, err := connection.Read(bytes); err == nil {
			log.Println("Server:>", string(bytes))
		} else if err == io.EOF {
			log.Fatalln("Server down!", err)
		} else {
			log.Println("Error while reading data from server!", err)
		}
	}
}

func bufferedCommunication(connection net.Conn) {
	for {
		fmt.Print("You:> ")
		input := make([]byte, 1024)
		os.Stdin.Read(input)
		message := string(input)
		writer := bufio.NewWriter(connection)
		if _, err := writer.WriteString(message); err == io.EOF {
			log.Fatalln("Server down!", err)
		} else if err != nil {
			log.Println("Error while sending the message:", err)
		}
		writer.Flush()
		reader := bufio.NewReader(connection)
		if res, err := reader.ReadString('\n'); err == nil {
			log.Println("Server> ", res)
		} else if err == io.EOF {
			log.Fatalln("Server down!", err)
		} else {
			log.Println("Something went wrong while receiving message from the server:", err)
		}
	}
}
func main() {
	if connection, err := net.Dial("tcp", "0.0.0.0:8080"); err == nil {
		defer connection.Close()
		//nonBufferedCommunication(connection)
		bufferedCommunication(connection)
	} else {
		log.Fatalln("Connection got fucked:", err)
	}
}