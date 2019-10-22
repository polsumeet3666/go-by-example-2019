package main

import "fmt"

// only sending channel
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// only receives channel
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {

	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)

	fmt.Println(<-pongs)

}
