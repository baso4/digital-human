package models

// Event struct represents an event
type Event struct {
	Name               string
	Health             int
	Happiness          int
	RelationshipEffect map[string]int // Effect on relationships with other characters
}
