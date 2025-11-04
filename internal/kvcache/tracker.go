
// Package kvcache tracks KV-cache memory reservations for drafts.
package kvcache

import (
	"sync"
)

// Tracker prices draft cache footprints against the target budget.
type Tracker struct {
	mu         sync.RWMutex
	maxBytes   int64
	reserved   int64
	headroom   int64
	overBudget bool
}

// NewTracker builds a tracker for the given target cache budget.
func NewTracker(maxBytes int64) *Tracker {
	return &Tracker{maxBytes: maxBytes}
}

// Reserve books footprint bytes for one draft proposal.
// It returns false when the reservation would push the target over budget
// (after headroom is subtracted).
func (t *Tracker) Reserve(footprint int64, headroomRatio float64) bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	headroom := int64(float64(t.maxBytes) * headroomRatio)
	usable := t.maxBytes - headroom
	if t.reserved+footprint > usable {
		t.overBudget = true
		return false
	}
	t.reserved += footprint
	t.headroom = headroom
	return true
}

// Release returns footprint bytes to the pool.
func (t *Tracker) Release(footprint int64) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.reserved -= footprint
	if t.reserved < 0 {
		t.reserved = 0
	}
	t.overBudget = false
}

// Used reports the current reserved footprint.
func (t *Tracker) Used() int64 {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.reserved
}

// Headroom reports the configured safety margin in bytes.
func (t *Tracker) Headroom() int64 {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.headroom
}

// OverBudget reports whether the last reservation was rejected.
func (t *Tracker) OverBudget() bool {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.overBudget
}
