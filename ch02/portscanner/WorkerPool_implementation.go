package main

import (
	"fmt"
	//"net"
	"sync"
)

func WorkerPool(id int8, ports chan int16, wg *sync.WaitGroup) {
	i := 0
	for port := range ports {
		fmt.Println("worker id = ", id, "\tport = ", port, "\tcounter = ", i)
		wg.Done()
		i++
	}
	fmt.Println("Worker ", id, " has handled", i, "ports")
}
func main() {

	var wg sync.WaitGroup
	ports := make(chan int16, 100)
	fmt.Println(ports, cap(ports))
	// start goroutines waiting for channel
	capacity := int8(cap(ports))
	for i := int8(0); i < capacity; i++ {
		go WorkerPool(i, ports, &wg)  // create 100 go routines, which their execution is blocked on for loop, 
		// until something is sent to the channel
	}

	// sending data to then channel 
	for i := int16(1); i < 1024; i++ {
		wg.Add(1)
		ports <- i
	}

	wg.Wait()
	close(ports)

	var x string
	fmt.Scanln(&x)
}