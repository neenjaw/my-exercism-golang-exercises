package robotname

import (
	"errors"
	"math/rand"
	"sync"
)

type registry struct {
	availableNames []string
	takenNames     map[string]bool
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
	r.takenNames[name] = true
	return name, nil
}

// ReleaseName releases the name
func (r *registry) ReleaseName(name string) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	if _, chk := r.takenNames[name]; !chk {
		return errors.New("name is not currently in use")
	}

	delete(r.takenNames, name)
	r.availableNames = append(r.availableNames, name)

	return nil
}
