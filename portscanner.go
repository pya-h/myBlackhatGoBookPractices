package main

import ("fmt"
	"net"
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
	// NonConcurrentScan()

	// Concurrent scan, naive way
	// for i := 1; i < 
}
