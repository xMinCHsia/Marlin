
package accept

import "sync"

// Window keeps a sliding window of acceptance events per draft model.
type Window struct {
	mu     sync.Mutex
	size   int
	events map[string][]bool // true = accepted
}

// NewWindow builds a window that remembers the last size events per draft.
func NewWindow(size int) *Window {
	if size < 8 {
		size = 8
	}
	return &Window{size: size, events: map[string][]bool{}}
}

