package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{
		"https://www.google.com",
		"https://www.reddit.com",
		"https://www.github.com",
		"https://app.jvmanagement.in",
		"https://www.stackover.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for {
		go checkLink(<-c, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}