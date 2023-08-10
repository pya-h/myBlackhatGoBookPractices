package main

import (
	"fmt"
	"time"
)

func CountSecs(c chan int8) {
	mins := int8(0)

	init := int8(0)
	for {
		if init == 0 {

			fmt.Printf("sec: 00")
		}
		for i := 1; i < 60; i++ {
			time.Sleep(time.Second)
			fmt.Printf("\b\b%02d", i)
		}
		mins = (mins + 1) % 60
		c <- mins
		init = <-c
	}
}

func CountMins(c chan int8) {
	hours := 0
	for {
		mins := <-c
		if mins == 0 {
			hours++
		}
		fmt.Println("\nmin: ", mins, " hour: ", hours)
		c <- 0
	}
}

func main() {
	c := make(chan int8)
	go CountSecs(c)
	go CountMins(c)

	var x string
	fmt.Scanln(&x)
}
