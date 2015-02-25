# goroutines + go channels in action

Mostly pulled from  
http://www.golang-book.com/10/index.htm  

## goroutine
A goroutine is a function that is capable of running concurrently with other functions.
To create one you use the keyword `go` followed by the function.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// some function that does some async thing
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

```

## go Channel
Provide a way for two goroutines to communicate with one another and sychronise their execution.
``` go
package main

import (
	"fmt"
	"strconv"
	"time"
)

func pinger(c chan<- string) {
	for i := 0; ; i++ { //will loop and send message onto the channel
		c <- "ping " + strconv.Itoa(i)
	}
}

func printer(c chan string) {
	for { //will loop forever reading off the channel
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
```
