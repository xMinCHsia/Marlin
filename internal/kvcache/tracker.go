
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
