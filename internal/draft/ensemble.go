
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

// NewEnsemble builds the ensemble from configuration.
func NewEnsemble(models []config.DraftModel) *Ensemble {
	w := map[string]float64{}
	for _, m := range models {
		w[m.ID] = m.Weight
	}
	return &Ensemble{models: models, weights: w}
}

// Models returns the draft ids in stable order.
func (e *Ensemble) Models() []string {
	out := make([]string, 0, len(e.models))
	for _, m := range e.models {
		out = append(out, m.ID)
	}
	return out
}

// Proposal is one draft's suggested continuation.
type Proposal struct {
	DraftID   string
	Tokens    []string
	Footprint int64
	Length    int
}

