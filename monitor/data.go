package monitor

// DataMonitor ...
type DataMonitor struct {
	URLS []string
}

// AddURL ...
func (d *DataMonitor) AddURL(urlToAdd string) {
	d.URLS = append(d.URLS, urlToAdd)
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
		d.URLS = append(d.URLS[:i], d.URLS[i+1:]...)
	}
}
