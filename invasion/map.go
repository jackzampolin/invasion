package invasion

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	numTurns       = 10000
	CityNameLength = 16
	cityNameLength = 16
)

// NewMap creates a new map initialized with a source of randomness
func NewMap() *Map {
	return &Map{rand: rand.New(rand.NewSource(time.Now().UnixNano()))}
}

// NewMapFromFile returns a new instance of map based on the file
func NewMapFromFile(file string) (*Map, error) {
	// Instantiate the map
	m := NewMap()

	// Open the file and return any errors
	f, err := os.Open(file)
	if err != nil {
		return m, err
	}
	defer f.Close()

	// Scan the file by lines
	scan := bufio.NewScanner(f)
	scan.Split(bufio.ScanLines)
	for scan.Scan() {
		city := &City{Roads: make([]*Road, 0)}
		// Split the lines on spaces
		l := strings.Split(strings.TrimSpace(scan.Text()), " ")
		// If the line is empty or too long ignore it
		if len(l) > 0 && len(l) < 6 && !strings.Contains(l[0], "=") {
			city.Name = l[0]
			l = l[1:]
			for _, rd := range l {
				// Split the roads on "="
				kv := strings.Split(rd, "=")
				city.Roads = append(city.Roads, &Road{Direction: kv[0], City: kv[1]})
			}
			m.Cities = append(m.Cities, city)
			continue
		}
	}
	return m, nil
}

// Map represents the game map. It contains an array of cities and some aliens
type Map struct {
	Cities []*City
	Aliens []*Alien

	rand *rand.Rand
}

// NewCities generates a random map to play this game on
func (m *Map) NewCities(n int) {
	cityNames := m.randStrings(n, cityNameLength)
	m.Cities = make([]*City, 0)
	for _, c := range cityNames {
		m.Cities = append(m.Cities, m.NewCity(c, cityNames))
	}
}

// Marshall returns the map in string format
func (m *Map) String() string {
	outBytes := make([]byte, 0, cityNameLength*5*len(m.Cities))
	for _, c := range m.Cities {
		outBytes = append(outBytes, []byte(c.String())...)
	}
	return string(outBytes)
}

// Play runs the game
func (m *Map) Play() {
	for i := 1; i < numTurns+1; i++ {
		m.MoveAliens()
		m.DestroyCitiesAndAliens(i)
		// If there are no remaining aliens then end the game
		if len(m.Aliens) == 0 {
			break
		}
	}
}

// MoveAliens moves all of the aliens
func (m *Map) MoveAliens() {
	for _, a := range m.Aliens {
		m.moveAlien(a)
	}
}

// DestroyCitiesAndAliens reconciles the map at the end of the turn
// by removing the aliens, cities, and roads that have been destroyed
func (m *Map) DestroyCitiesAndAliens(turn int) {
	// First create a mapping of city -> aliens
	out := make(map[string][]string, 0)
	for _, a := range m.Aliens {
		if _, pres := out[a.City]; pres {
			out[a.City] = append(out[a.City], a.Name)
			continue
		}
		out[a.City] = []string{a.Name}
	}

	// Then loop over and destroy all the cities > 2 aliens
	for k, v := range out {
		if len(v) > 1 {
			msg := fmt.Sprintf("%s has been destroyed by ", k)
			for _, a := range v {
				msg += fmt.Sprintf("alien %s and ", a)
			}
			msg = strings.TrimSuffix(msg, " and ")
			msg += fmt.Sprintf("! (turn %d)", turn)
			fmt.Println(msg)
			m.removeCity(k)
			m.removeAliens(v)
		}
	}
}

//
// // randStrings returns a list of n random strings of length l
// func (m *Map) randStrings(n, l int) []string {
// 	cityNames := make([]string, 0, n)
// 	for i := 0; i < n; i++ {
// 		b := make([]byte, l)
// 		for i := range b {
// 			b[i] = letterBytes[m.rand.Intn(len(letterBytes))]
// 		}
// 		cityNames = append(cityNames, string(b))
// 	}
// 	return cityNames
// }

func (m *Map) removeAliens(aliens []string) {
	out := make([]*Alien, 0)
	for _, a := range m.Aliens {
		delete := false
		for _, r := range aliens {
			if a.Name == r {
				delete = true
			}
		}
		if !delete {
			out = append(out, a)
		}
	}
	m.Aliens = out
}

func (m *Map) removeCity(city string) {
	cities := make([]*City, 0)
	for _, c := range m.Cities {
		if c.Name == city {
			continue
		}
		c.RemoveRoadTo(city)
		cities = append(cities, c)
	}
	m.Cities = cities
}

// Return a new city for the alien, if there are no roads return an empty string
func (m *Map) moveAlien(a *Alien) {
	a.Turns++
	for _, c := range m.Cities {
		if c.Name == a.City && len(c.Roads) > 0 {
			a.City = c.RandRoad().City
			return
		}
	}
}

// randCity returns a random city from the list of cities
func (m *Map) randCity() *City {
	return m.Cities[rand.Intn(len(m.Cities))]
}

// NewAliens adds n aliens to the map
func (m *Map) NewAliens(n int) {
	for _, a := range m.newAliens(n) {
		a.City = m.randCity().Name
		m.Aliens = append(m.Aliens, a)
	}
}
