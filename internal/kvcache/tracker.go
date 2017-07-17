
// Package kvcache tracks KV-cache memory reservations for drafts.
package kvcache

import (
	"sync"
)

// Tracker prices draft cache footprints against the target budget.
type Tracker struct {
	mu         sync.RWMutex
