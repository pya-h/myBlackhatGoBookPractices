package main

import (
	"fmt"
	"net"
	"sort"
)

func WorkerPool(address *string, ports chan int16, results chan int16) {
	for port := range ports {
		if connection, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *address, port)); err == nil {
			results <- port
			connection.Close()
		} else {
			results <- 0
		}
	}

}
// simple loading bar
func LoadingBar(load chan bool) {
	i := int8(0)
	for p := range load {
		p = p
		if i >= 3 {
			fmt.Print("\b\b\b   \b\b\b")
			i = 0
		}
		fmt.Print(".")
		i++
	}
}
func main() {
	var address string
	const LAST_PORT int16 = 1024
	fmt.Print("Target Address: ")
	fmt.Scanln(&address)
	// The higher the count of workers(100 here), the faster your program should execute. But if you add too many
		//workers, your results could become unreliable
	ports := make(chan int16, 100)
	results := make(chan int16)
	// start goroutines waiting for channel
	capacity := int8(cap(ports))
	for i := int8(0); i < capacity; i++ {
		go WorkerPool(&address, ports, results)  // create 100 go routines, which their execution is blocked on for loop, 
		// until something is sent to the channel
	}
	fmt.Print("Scanning ")
	load := make(chan bool)
	go LoadingBar(load)

	go func() {
		// this must be go routine so that results channel can be handled in the next lines
		for i := int16(1); i <= LAST_PORT; i++ {
			// sending data to then channel 
			ports <- i
		}
	} ()

	var openPorts []int16
	for i := int16(1); i <= LAST_PORT; i++ {
		port := <-results
		load <- true
		if port > 0 {
			openPorts = append(openPorts, port)
		}

	}
	close(load)
	close(results)
	close(ports)
	fmt.Println("\nDone")
	// sort.Ints(openPorts as  int)
	sort.Slice(openPorts, func(i, j int) bool { return openPorts[i] < openPorts[j]})

	for _, port := range openPorts {
		fmt.Printf("%d is open\n", port)
	}
}