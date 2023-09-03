package main

import (
	"fmt"
	"net"
	"sort"
	"flag"
)

func WorkerPool(address *string, ports chan uint16, results chan uint16, progress_trigger chan bool) {
	for port := range ports {
		if connection, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *address, port)); err == nil {
			results <- port
			connection.Close()
		} else {
			results <- 0
		}
		progress_trigger <- true
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

func UpdateProgress(number_of_ports uint16 , trigger chan bool) {
	checked_ports := uint16(0)
	var triggered bool
	for {
		triggered = <- trigger
		if triggered {
			checked_ports++
		}
		fmt.Printf("\b\b\b\b\b\b\b%5.2f %%", 100 * float32(checked_ports) / float32(number_of_ports))
	}
}

func main() {
	address := flag.String("a", "localhost", "Host address")
	first_port := flag.Int("f", 1, "First port")
	last_port := flag.Int("l", 1024, "Last port")
	
	flag.Parse()
	fmt.Println("Search started on:", *address, ", Ports:", *first_port, "->", *last_port)
	// The higher the count of workers(100 here), the faster your program should execute. But if you add too many
		//workers, your results could become unreliable
	ports := make(chan uint16, 100)
	results := make(chan uint16)
	// start goroutines waiting for channel
	capacity := int8(cap(ports))
	progress_trigger := make(chan bool)
	go UpdateProgress(uint16(*last_port - *first_port), progress_trigger)

	fmt.Print("Progress:  0.00 %")
	for i := int8(0); i < capacity; i++ {
		go WorkerPool(address, ports, results, progress_trigger)  // create 100 go routines, which their execution is blocked on for loop, 
		// until something is sent to the channel
	}
	//load := make(chan bool)
	//go LoadingBar(load)

	go func() {
		// this must be go routine so that results channel can be handled in the next lines
		for i := *first_port; i <= *last_port; i++ {
			// sending data to then channel 
			ports <- uint16(i)
		}
	} ()

	var openPorts []uint16
	for i := *first_port; i <= *last_port; i++ {
		port := <-results
		//load <- true
		if port > 0 {
			openPorts = append(openPorts, port)
		}

	}
	// close(load)
	close(results)
	close(ports)
	// close(progress_trigger)
	// sort.Ints(openPorts as  int)
	sort.Slice(openPorts, func(i, j int) bool { return openPorts[i] < openPorts[j]})
	fmt.Println()
	for index, port := range openPorts {
		fmt.Printf("%d\t", port)
		if (index+1) % 5 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
}