
// Package tuning self-adjusts draft lengths from acceptance statistics.
package tuning

import (
	"sync"

	"github.com/xMinCHsia/marlin/internal/accept"
	"github.com/xMinCHsia/marlin/internal/config"
)

// Autotuner moves draft lengths up when acceptance is high and down when
// it drifts, one step per probe interval.
type Autotuner struct {
	mu        sync.Mutex
	cfg       config.Tuning
	stats     *accept.Window
	lengths   map[string]int
	baselines map[string]float64
}

// NewAutotuner builds the tuner.
func NewAutotuner(cfg config.Tuning) *Autotuner {
	return &Autotuner{
		cfg:       cfg,
		stats:     accept.NewWindow(64),
		lengths:   map[string]int{},
		baselines: map[string]float64{},
	}
}

// Observe records an acceptance event and returns the new draft length
// for that model (same as current when tuning is disabled).
func (a *Autotuner) Observe(draftID string, accepted bool, currentLen int) int {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.stats.Record(draftID, accepted)
	if !a.cfg.Enabled {
		return currentLen
	}
	rate := a.stats.Rate(draftID)
	base, seen := a.baselines[draftID]
	if !seen {
		a.baselines[draftID] = rate
		a.lengths[draftID] = clamp(currentLen, a.cfg)
		return a.lengths[draftID]
	}
	cur := a.lengths[draftID]
	if cur == 0 {
		cur = clamp(currentLen, a.cfg)
	}
	if a.stats.Drift(draftID, base, 0.15) {
		cur -= a.cfg.Step
	} else if rate > base+0.05 {
		cur += a.cfg.Step
	}
	cur = clamp(cur, a.cfg)
	a.lengths[draftID] = cur
	return cur
}

// Length returns the tuned length for a draft.
func (a *Autotuner) Length(draftID string) int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.lengths[draftID]
