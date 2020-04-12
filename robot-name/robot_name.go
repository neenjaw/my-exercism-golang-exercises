package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const number = "1234567890"

// Robot struct represents a robot with a name
type Robot struct {
	name *string
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// Name returns the robot's name
func (r *Robot) Name() (string, error) {
	if r.name == nil {
		name := makeName()
		r.name = &name
	}

	return *(r.name), nil
}

// Reset resets the robot's name
func (r *Robot) Reset() {
	r.name = nil
}

func makeName() (name string) {
	r1 := string(rand.Intn(26) + 'A')
	r2 := string(rand.Intn(26) + 'A')
	num := rand.Intn(1000)
	return fmt.Sprintf("%s%s%03d", r1, r2, num)
}
