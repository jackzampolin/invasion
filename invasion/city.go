package invasion

import (
	"fmt"
	"math/rand"
	"time"
)

// NewCity creates a new city with roads to other cities in the map
func (m *Map) NewCity(name string, cities []string) *City {
	return &City{
		Name:  name,
		Roads: m.NewRoads(cities),
	}
}

func (m *Map) NewCityBytes(name string, cities [][]byte) *City {
	return &City{
		Name:  name,
		Roads: m.NewRoadsBytes(cities),
	}
}

// City represents a city with roads going out of it
type City struct {
	Name  string
	Roads []*Road
}

// RandRoad returns a random road from a city
func (c *City) RandRoad() *Road {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	return c.Roads[r.Intn(len(c.Roads))]
}

// RemoveRoadTo removes a road to a named city from a calling city
func (c *City) RemoveRoadTo(city string) {
	roads := make([]*Road, 0)
	for _, r := range c.Roads {
		if r.City != city {
			roads = append(roads, r)
		}
	}
	c.Roads = roads
}

func (c *City) String() string {
	out := c.Name
	if len(c.Roads) > 0 {
		for _, r := range c.Roads {
			out += fmt.Sprintf(" %s", r.String())
		}
	}
	out += "\n"
	return out
}

// Bytes returns the byte representation of the city
func (c *City) Bytes() []byte {
	return []byte(c.String())
}
