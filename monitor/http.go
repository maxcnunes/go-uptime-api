package monitor

import (
	"fmt"
	"net/http"
	"time"
)

// HTTPResponse defines the results from a http request
type HTTPResponse struct {
	URL      string
	Response *http.Response
	Error    error
}

// AsyncHTTPGets requests urls concurrently
func AsyncHTTPGets(urls []string) []*HTTPResponse {
	ch := make(chan *HTTPResponse)
	responses := []*HTTPResponse{}

	// executes concurrently the request to all urls
	for _, url := range urls {
		go func(url string) {
			fmt.Printf("Fetching %s \n", url)
			resp, err := http.Get(url)
			ch <- &HTTPResponse{url, resp, err}
		}(url)
	}

	// watis finish all requests
	for {
		select {
		case r := <-ch:
			fmt.Printf("%s was fetched\n", r.URL)
			responses = append(responses, r)
			if len(responses) == len(urls) {
				return responses
			}
		case <-time.After(50 * time.Millisecond):
			fmt.Println(".")
		}
	}

	return responses
}
