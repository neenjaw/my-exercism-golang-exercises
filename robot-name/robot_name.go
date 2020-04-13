package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

// Robot struct represents a robot with a name
type Robot struct {
	name string
}

var usedNames map[string]bool

func init() {
	usedNames = make(map[string]bool, 26*26*10*10*10)
	rand.Seed(time.Now().UTC().UnixNano())
}

// Name returns the robot's name
func (r *Robot) Name() (string, error) {
	for r.name == "" {
		r.name = makeName()
		if _, used := usedNames[r.name]; used {
			r.name = ""
		} else {
			usedNames[r.name] = true
		}
	}

	return r.name, nil
}

// Reset resets the robot's name
func (r *Robot) Reset() {
	r.name = ""
}

func makeName() (name string) {
	r1 := string(rand.Intn(26) + 'A')
	r2 := string(rand.Intn(26) + 'A')
	num := rand.Intn(1000)
	return fmt.Sprintf("%s%s%03d", r1, r2, num)
}
