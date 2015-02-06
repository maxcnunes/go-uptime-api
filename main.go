package main

import (
	"fmt"
	"github.com/maxcnunes/monitor/http"
)

var urls = []string{
	"https://google.com/",
	"http://twitter.com/",
}

func main() {
	results := http.AsyncHTTPGets(urls)
	for _, result := range results {
		if result.Response != nil {
			fmt.Printf("%s status: %s\n", result.URL, result.Response.Status)
		}
	}
}
