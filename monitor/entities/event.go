package entities

// Event has the event type triggered target data manipulation actions
// and also has the related target data
type Event struct {
	Event  string  `json:"event"`
	Target *Target `json:"target"`
}

const (
	// Added is the event type triggered when a new has been created
	Added = "added"
	// Removed is the event type triggered when an existing target has been removed
	Removed = "removed"
	// Updated is the event type triggered when an existing target has been updated
	Updated = "updated"
)
