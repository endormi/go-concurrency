package main

import (
	"fmt"
	"runtime"
)

func availableThreads() int {
	threads := runtime.GOMAXPROCS(0)

	return threads
}

func main() {
	runtime.GOMAXPROCS(5)

	fmt.Printf("%d thread(s) available.", availableThreads())
}
