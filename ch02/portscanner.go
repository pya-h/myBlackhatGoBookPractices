package main

import (
	"fmt"
	"net"
	"sync"
)

func NonConcurrentScan() {
	for port := 1; port < 1024; port++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", port)
		if connection, err := net.Dial("tcp", address); err == nil {
			fmt.Println(address, ": Yo bitch", connection)
			connection.Close()
		} else {
			fmt.Println(address, ": Connection failed; ", err)
		}
	}
}


func main() {
	//NonConcurrentScan()

	// Concurrent scan, naive way
	var waiter sync.WaitGroup
	var address string
	fmt.Print("Target Address: ")
	fmt.Scanln(&address)
	openPorts := make([]int16, 0, 0)
	for i := int16(1); i < 1024; i++ {
		waiter.Add(1) // increment the counter of wait group to inform that there is new routine to wait for
		go func(port int16) {
			defer waiter.Done() // decrease the wait counter after the routine is finished
			if connection, err := net.Dial("tcp", fmt.Sprintf("%s:%d", address, port)); err == nil {
				fmt.Printf("Port #%d is open\n", port)
				openPorts = append(openPorts, port)
				connection.Close()
			} else {
				fmt.Printf("Port #%d is closed or filtered: %s\n", port, err)
			}
		}(i)
	}
	waiter.Wait() // wait until the wait griup counter becomes zero
	fmt.Println("---------------------------------------------------\nOpen Ports are:\t", openPorts)
}
