
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

// Compose builds a proposal batch: each draft contributes up to its
// per-token budget, capped by the total token budget for the request.
type cand struct {
	model config.DraftModel
	len   int
}

func (e *Ensemble) Compose(budget int) ([]Proposal, int) {
	totalWeight := total(e.weights)
	var cands []cand
	for _, m := range e.models {
		w := e.weights[m.ID]
		if w <= 0 {
			continue
		}
		n := int(float64(budget) * w / totalWeight)
		if n < 1 {
			n = 1
		}
		if n > m.MaxProposalLen {
			n = m.MaxProposalLen
		}
		cands = append(cands, cand{model: m, len: n})
	}
	// allocate remaining budget to the highest-weight drafts
	remaining := budget - sumLen(cands)
	for remaining > 0 && len(cands) > 0 {
		sort.Slice(cands, func(i, j int) bool {
			return e.weights[cands[i].model.ID] > e.weights[cands[j].model.ID]
		})
		for i := range cands {
			if remaining <= 0 {
				break
			}
			if cands[i].len < cands[i].model.MaxProposalLen {
				cands[i].len++
				remaining--
			}
