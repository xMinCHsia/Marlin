
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

// Record appends one event and drops the oldest beyond the window size.
func (w *Window) Record(draftID string, accepted bool) {
	w.mu.Lock()
	defer w.mu.Unlock()
	ev := w.events[draftID]
	ev = append(ev, accepted)
	if len(ev) > w.size {
		ev = ev[len(ev)-w.size:]
	}
	w.events[draftID] = ev
}

// Rate returns the acceptance rate over the current window.
func (w *Window) Rate(draftID string) float64 {
	w.mu.Lock()
	defer w.mu.Unlock()
	ev := w.events[draftID]
	if len(ev) == 0 {
		return 0
	}
	ok := 0
	for _, a := range ev {
		if a {
			ok++
		}
	}
	return float64(ok) / float64(len(ev))
}

// Drift reports a rate drop larger than threshold vs baseline.
// baseline is the rate observed before the current window.
func (w *Window) Drift(draftID string, baseline, threshold float64) bool {
	rate := w.Rate(draftID)
