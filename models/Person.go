package models

import (
	"fmt"
	"math/rand"
	"time"
)

// Person struct represents a character
type Person struct {
	Name         string
	Age          int
	Health       int
	Happiness    int
	Relationship map[string]int // Relationship with other characters, using other characters' names as keys
	Events       []Event
}

// NewPerson creates a new character
func NewPerson(name string, age int) *Person {
	return &Person{
		Name:         name,
		Age:          age,
		Health:       100,
		Happiness:    50,
		Relationship: make(map[string]int),
		Events:       []Event{},
	}
}

// AddEvent adds an event to the character's events list
func (p *Person) AddEvent(event Event) {
	p.Events = append(p.Events, event)
}

// Live simulates the character's life
func (p *Person) Live() {
	fmt.Printf("%s starts a new day...\n", p.Name)

	// Simulate some random events
	p.randomEvent()

	// Display the current status of the character
	p.displayStatus()
}

// randomEvent simulates some random events
func (p *Person) randomEvent() {
	rand.Seed(time.Now().UnixNano())
	eventType := rand.Intn(4) // Randomly choose event type

	switch eventType {
	case 0:
		p.work()
	case 1:
		p.socialize()
	case 2:
		p.rest()
	case 3:
		p.handleEvent()
	}
}

// work simulates work
func (p *Person) work() {
	fmt.Printf("%s is working...\n", p.Name)
	p.Happiness -= 10
	p.Health -= 5
}

// socialize simulates socializing
func (p *Person) socialize() {
	fmt.Printf("%s is socializing...\n", p.Name)
	p.Happiness += 15
	p.Health -= 5

	// Randomly choose a friend
	friend := p.chooseRandomFriend()
	if friend != "" {
		fmt.Printf("%s has a pleasant conversation with %s.\n", p.Name, friend)
		p.Relationship[friend] += 10
	}
}

// rest simulates resting
func (p *Person) rest() {
	fmt.Printf("%s is resting...\n", p.Name)
	p.Happiness += 5
	p.Health += 10
}

// handleEvent simulates handling an event from the events list
func (p *Person) handleEvent() {
	if len(p.Events) == 0 {
		fmt.Printf("%s has no pending events.\n", p.Name)
		return
	}

	// Randomly choose an event from the list
	eventIndex := rand.Intn(len(p.Events))
	event := p.Events[eventIndex]

	// Apply the event's effects
	fmt.Printf("%s is handling the event: %s...\n", p.Name, event.Name)
	p.Happiness += event.Happiness
	p.Health += event.Health

	// Apply relationship effects
	for friend, effect := range event.RelationshipEffect {
		p.Relationship[friend] += effect
	}

	// Remove the handled event from the list
	p.Events = append(p.Events[:eventIndex], p.Events[eventIndex+1:]...)
}

// chooseRandomFriend randomly chooses a friend
func (p *Person) chooseRandomFriend() string {
	if len(p.Relationship) == 0 {
		return ""
	}

	friendIndex := rand.Intn(len(p.Relationship))
	i := 0
	for friend := range p.Relationship {
		if i == friendIndex {
			return friend
		}
		i++
	}
	return ""
}

// displayStatus displays the current status of the character
func (p *Person) displayStatus() {
	fmt.Printf("%s's status:\n", p.Name)
	fmt.Printf("Age: %d\n", p.Age)
	fmt.Printf("Health: %d\n", p.Health)
	fmt.Printf("Happiness: %d\n", p.Happiness)
	fmt.Printf("Relationships with others: %v\n", p.Relationship)
	fmt.Printf("Pending Events: %v\n", p.Events)
	fmt.Println()
}
