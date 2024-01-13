package main

import "github.com/baso4/digital-human/models"

func main() {
	// Create two characters
	person1 := models.NewPerson("Alice", 25)
	person2 := models.NewPerson("Bob", 28)

	// Set initial relationship
	person1.Relationship[person2.Name] = 20
	person2.Relationship[person1.Name] = 20

	// Create some events
	event1 := models.Event{
		Name:      "Birthday Party",
		Health:    10,
		Happiness: 20,
		RelationshipEffect: map[string]int{
			person1.Name: 15,
			person2.Name: 15,
		},
	}
	event2 := models.Event{
		Name:      "Work Deadline",
		Health:    -10,
		Happiness: -20,
		RelationshipEffect: map[string]int{
			person1.Name: -10,
			person2.Name: -10,
		},
	}

	// Add events to characters
	person1.AddEvent(event1)
	person2.AddEvent(event2)

	// Simulate life for a period of time
	for i := 0; i < 5; i++ {
		person1.Live()
		person2.Live()
	}
}
