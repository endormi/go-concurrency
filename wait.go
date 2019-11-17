package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var x sync.WaitGroup
	x.Add(1)

	go func() {
		amount("current")

		x.Done()
	}()

	x.Wait()
}

func amount(am string) {
	for i := 1; i < 10; i++ {
		fmt.Println(i, am)

		time.Sleep(time.Millisecond * 100)
	}
}
