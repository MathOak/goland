package main

import (
	"fmt"
	"time"
)

func pong(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func ping(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func gamePrint(c chan string) {
	for {
		fmt.Println(<-c)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go ping(c)
	go gamePrint(c)
	go pong(c)

	var input string

	fmt.Scanln(&input)
}
