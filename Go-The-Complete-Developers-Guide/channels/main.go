package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		 "http://facebook.com",
		 "http://stackoverflow.com",
		 "http://golang.org",
		 "http://amazon.com",
		 "http://localhost",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c) //receiving a message from a channel is a blocking call
	// }
	count := 0
	for l := range c{
		go func (l string)  {
			time.Sleep(time.Second * 5)
			checkLink(l, c)
		}(l)

		count++
		if count > 10 {
			return
		}

	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	
	if err != nil {
		fmt.Println(link, "seems down now!")
		c <- link
		return
	}
	fmt.Println(link, "is up now!")
	c <- link
}