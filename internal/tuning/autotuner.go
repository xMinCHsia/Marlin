
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
