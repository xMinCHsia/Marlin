
// Package draft manages the draft-model ensemble and proposal batches.
package draft

import (
	"fmt"
	"sort"

	"github.com/xMinCHsia/marlin/internal/config"
)

// Ensemble holds the draft models and their current weights.
type Ensemble struct {
	models  []config.DraftModel
	weights map[string]float64
}

