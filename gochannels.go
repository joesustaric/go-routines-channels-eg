package main

import (
	"fmt"
	"strconv"
	"time"
)

func pinger(c chan<- string) { // only write to the channel
	for i := 0; ; i++ { //will loop and send message onto the channel
		c <- "ping " + strconv.Itoa(i)
	}
}

func ponger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func printer(c <-chan string) { //only read from the channel
	for { //will loop forever reading off the channel
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
