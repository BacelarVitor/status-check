package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()

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

	for i := 0; i < len(links); i++ {
		fmt.Println(<- c)
	}
	
	elapsed := time.Since(start)
	fmt.Println()
	fmt.Println()
	fmt.Println(elapsed)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down, I think"
		return
	}

	fmt.Println(link, "is up!")
	c <- "Yep is Up"
}