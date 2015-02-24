package monitor

// Target ...
type Target struct {
	URL    string `json:"url"`
	Status string `json:"status"`
}

// Status
const (
	StatusUp   = "up"
	StatusDown = "down"
	StatusNone = "none"
)
