package main

import (
	"io"
	"log"
	"net"
)

func forward(src net.Conn) {
	if dst, err := net.Dial("tcp", "google.com:80"); err == nil {
		defer dst.Close()
		go func() {
			if _, err := io.Copy(dst, src); err != nil {
				log.Fatalln("Cannot forward data to the target:", err)
			}
		} ()

		if _, err := io.Copy(src, dst); err != nil {
			log.Fatalln("Cannot return data from target to the you:", err)
		}
	} else {
		log.Fatalln("Something went wrong while forwarding:", err)
	}

}

func main() {
	if listener, err := net.Listen("tcp", ":80"); err == nil {
		log.Println("TCP Server is up and listening...")
		for {
			if connection, err := listener.Accept(); err == nil {
				go forward(connection)
			} else {
				log.Println("Client Cannot connect to the proxy server:", err)
			}
		}
	} else {
		log.Fatalln("Can not start the proxy server:", err)
	}
}