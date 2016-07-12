
// Package admission implements KV-cache-aware admission control.
package admission

import (
	"math"
	"sync"
	"time"

	"github.com/xMinCHsia/marlin/internal/config"
	"github.com/xMinCHsia/marlin/internal/kvcache"
)

// Controller admits or defers draft proposals based on cache pressure.
type Controller struct {
	mu      sync.Mutex
	cfg     config.Admission
	tracker *kvcache.Tracker
	waits   map[string]time.Time
}

// NewController builds a controller.
func NewController(cfg config.Admission, tracker *kvcache.Tracker) *Controller {
	return &Controller{cfg: cfg, tracker: tracker, waits: map[string]time.Time{}}
}

// Admit decides whether draft id may propose now.
// It returns the seconds to wait when the answer is no.
func (c *Controller) Admit(id string, footprint int64) (ok bool, waitS int) {
	c.mu.Lock()
