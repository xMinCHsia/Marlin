
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
