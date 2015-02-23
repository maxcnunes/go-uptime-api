package monitor

import (
	"log"
)

// DataMonitor ...
type DataMonitor struct {
	URLS   []string
	Events chan int
}

const (
	// Added ...
	Added = 1 << iota
	// Removed ...
	Removed
)

// AddURL ...
func (d *DataMonitor) AddURL(urlToAdd string) {
	log.Printf("Adding url %s", urlToAdd)
	d.URLS = append(d.URLS, urlToAdd)
	go func() { d.Events <- Added }()
}

// RemoveURL ...
func (d *DataMonitor) RemoveURL(urlToRemove string) {
	i := -1
	for index, url := range d.URLS {
		if url == urlToRemove {
			i = index
			break
		}
	}

	if i > 0 {
		log.Printf("Removing url %s", urlToRemove)
		d.URLS = append(d.URLS[:i], d.URLS[i+1:]...)
		go func() { d.Events <- Removed }()
	}
}
