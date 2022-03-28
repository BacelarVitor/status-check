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

	for _, link := range links {

	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}

	fmt.Println(link, "is up!")
}