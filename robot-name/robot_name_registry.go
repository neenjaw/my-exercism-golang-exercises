package robotname

import (
	"errors"
	"math/rand"
	"sync"
)

type registry struct {
	availableNames []string
	mux            sync.Mutex
}

// GetName returns a name at random from the name registry if one is available
func (r *registry) GetName() (string, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	names := r.availableNames
	if len(names) == 0 {
		return "", errors.New("no names remaining")
	}
	randIndex := rand.Intn(len(names))
	name := names[randIndex]
	names[len(names)-1], names[randIndex] = names[randIndex], names[len(names)-1]
	r.availableNames = names[:len(names)-1]
	return name, nil
}
