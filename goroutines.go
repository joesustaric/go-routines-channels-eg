package main

import (
	"fmt"
	"math/rand"
	"time"
)

// some fucntion that does some async thing
func asyncCommand(number int) {
	for i := 1; i <= 10; i++ {
		fmt.Println("asyncCommand #", number, ":loop #", i)
		amount := time.Duration(rand.Intn(1000))
		time.Sleep(time.Millisecond * amount)
	}
}

func main() {
	for i := 1; i <= 10; i++ {
		go asyncCommand(i) //call it
	}
	var input string
	fmt.Scanln(&input)
}
