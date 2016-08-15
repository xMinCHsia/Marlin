
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
	defer c.mu.Unlock()
	if c.tracker.Reserve(footprint, c.cfg.HeadroomRatio) {
		delete(c.waits, id)
		return true, 0
	}
	last, seen := c.waits[id]
	if !seen {
		c.waits[id] = time.Now()
		return false, c.cfg.ProbeIntervalS
	}
	elapsed := time.Since(last).Seconds()
	backoff := math.Pow(2, elapsed/float64(c.cfg.ProbeIntervalS))
	wait := int(math.Ceil(backoff * float64(c.cfg.ProbeIntervalS)))
	c.waits[id] = time.Now()
	return false, wait
}

// Release returns the footprint after a proposal completes.
func (c *Controller) Release(footprint int64) {
	c.tracker.Release(footprint)
}
