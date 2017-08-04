
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
