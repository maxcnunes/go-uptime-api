package monitor

// Event ...
type Event struct {
	Event  string  `json:"event"`
	Target *Target `json:"target"`
}

const (
	// Added ..
	Added = "added"
	// Removed ...
	Removed = "removed"
	// Updated ...
	Updated = "updated"
)
