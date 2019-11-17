package main

import (
	"fmt"
	"time"
)

func amount(am string) {
	for i := 1; true; i++ {
		fmt.Println(i, am)

		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	go amount("current")
	go amount("secondary")
	amount("last")

	fmt.Scanln()
}
