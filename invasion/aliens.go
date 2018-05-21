package invasion

import "fmt"

const alienNameLength = 8

// Alien represents an alien
type Alien struct {
	Name  string
	City  string
	Turns int
}

// NewAliens creates n new Alien objects with random names
func (m *Map) newAliens(n int) []*Alien {
	out := make([]*Alien, 0)
	names := m.randStrings(n, alienNameLength)
	for _, name := range names {
		out = append(out, &Alien{Name: name, Turns: 0})
	}
	return out
}

func (a Alien) String() string {
	return fmt.Sprintf("%s city=%s turns=%d", a.Name, a.City, a.Turns)
}
