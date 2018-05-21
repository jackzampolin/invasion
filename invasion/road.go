package invasion

import (
	"fmt"
	"math/rand"
	"time"
)

var directions = []string{"north", "east", "south", "west"}

// NewRoads returns 1-4 roads that connect to other cities in the cities array
func NewRoads(cities []string) []*Road {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	roads := make([]*Road, 0)

	// copy directions slice
	// dir := directions

	// shuffle the directions
	for i := range directions {
		j := r.Intn(i + 1)
		directions[i], directions[j] = directions[j], directions[i]
	}

	// Random number of roads
	numRoads := r.Intn(4) + 1

	// Pull off the first numRoads roads
	for _, d := range directions[:numRoads] {
		roads = append(roads, &Road{Direction: d, City: cities[rand.Intn(len(cities))]})
	}

	return roads
}

// Road represents a connection between cities
type Road struct {
	Direction string
	City      string
}

func (r Road) String() string {
	return fmt.Sprintf("%s=%s", r.Direction, r.City)
}
