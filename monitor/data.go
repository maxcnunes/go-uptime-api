package monitor

import (
	"log"
)

// DataMonitor ...
type DataMonitor struct {
	Targets []Target
	Events  chan Event
}

// AddTarget ...
func (d *DataMonitor) AddTarget(url string) {
	i := d.findIndexByURL(url)
	if i < 0 {
		log.Printf("Adding target %s", url)
		target := Target{URL: url, Status: StatusNone}
		d.Targets = append(d.Targets, target)
		go func() { d.Events <- Event{Event: Added, Target: target} }()
	}
}

// RemoveTarget ...
func (d *DataMonitor) RemoveTarget(url string) {
	i := d.findIndexByURL(url)
	if i > 0 {
		log.Printf("Removing url %s", url)
		target := d.Targets[i]
		d.Targets = append(d.Targets[:i], d.Targets[i+1:]...)
		go func() { d.Events <- Event{Event: Added, Target: target} }()
	}
}

// UpdateStatusByURL ...
func (d *DataMonitor) UpdateStatusByURL(url string, status string) {
	i := d.findIndexByURL(url)
	if i > 0 {
		log.Printf("Updating url %s to status %s", url, status)
		target := d.Targets[i]
		target.Status = status
		go func() { d.Events <- Event{Event: Updated, Target: target} }()
	}
}

// GetAllURLS ...
func (d *DataMonitor) GetAllURLS() []string {
	urls := []string{}
	for _, target := range d.Targets {
		urls = append(urls, target.URL)
	}
	return urls
}

func (d *DataMonitor) findIndexByURL(url string) int {
	for index, target := range d.Targets {
		if target.URL == url {
			return index
		}
	}
	return -1
}
