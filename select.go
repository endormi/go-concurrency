package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	cs := make(chan string)

	go func() {
		for {
			c <- "Every 100ms"
			time.Sleep(time.Millisecond * 100)
		}
	}()

	go func() {
		for {
			cs <- "Every second"
			time.Sleep(time.Second * 1)
		}
	}()

	for {
		select {
		case r := <-c:
			fmt.Println(r)
		case x := <-cs:
			fmt.Println(x)
		}
	}
}
