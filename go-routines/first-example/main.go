package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}

	wg.Add(len(words))

	for i, w := range words {
		go printSomething(fmt.Sprintf("%d, %s", i, w), &wg)
	}

	wg.Wait()
	//time.Sleep(1 * time.Second) // Terrible way

	wg.Add(1)
	printSomething("This is the second thing to be printed!", &wg)
}
