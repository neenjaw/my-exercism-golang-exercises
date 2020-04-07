package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const maximumNames = 26 * 26 * 10 * 10 * 10
const alpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Robot struct represents a robot with a name
type Robot struct {
	name *string
}

var nameRegistry registry

// Build the list of all robot names
func init() {
	nameRegistry = registry{
		availableNames: make([]string, 0, maximumNames)}

	var sb strings.Builder
	for _, l1 := range alpha {
		for _, l2 := range alpha {
			for i := 0; i <= 999; i++ {
				sb.WriteRune(l1)
				sb.WriteRune(l2)
				sb.WriteString(fmt.Sprintf("%03d", i))
				nameRegistry.availableNames = append(nameRegistry.availableNames, sb.String())
				sb.Reset()
			}
		}
	}
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// Name returns the robot's name
func (r *Robot) Name() (string, error) {
	if r.name == nil {
		name, err := nameRegistry.GetName()
		if err != nil {
			return "", errors.New("no name available")
		}
		r.name = &name
	}

	return *(r.name), nil
}

// Reset resets the robot's name
func (r *Robot) Reset() {
	r.name = nil
}
