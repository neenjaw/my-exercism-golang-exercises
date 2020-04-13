package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Robot struct represents a robot with a name
type Robot struct {
	name string
}

var usedNames map[string]bool
var maximumNames int = 26 * 26 * 10 * 10 * 10

func init() {
	usedNames = make(map[string]bool)
	rand.Seed(time.Now().UTC().UnixNano())
}

// Name returns the robot's name
func (r *Robot) Name() (string, error) {
	if len(usedNames) == maximumNames {
		return "", errors.New("No more names")
	}

	if r.name == "" {
		r.name = makeName()
		for usedNames[r.name] {
			r.name = makeName()
		}
		usedNames[r.name] = true
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
