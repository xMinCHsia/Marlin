
// Package plan emits deterministic speculative-decoding plans.
package plan

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"sort"
)

// Step is one unit of a plan.
type Step struct {
	DraftID  string   `json:"draft_id"`
	Tokens   []string `json:"tokens"`
	Accepted int      `json:"accepted"`
	Len      int      `json:"len"`
}

// Plan is the full execution plan for one request.
type Plan struct {
	RequestID   string `json:"request_id"`
	Target      string `json:"target"`
	Budget      int    `json:"budget"`
	Steps       []Step `json:"steps"`
	SummaryHash string `json:"summary_hash"`
}

// Build assembles a plan from steps and stamps a deterministic hash.
func Build(requestID, target string, budget int, steps []Step) *Plan {
	sort.Slice(steps, func(i, j int) bool {
		return steps[i].DraftID < steps[j].DraftID
	})
	p := &Plan{RequestID: requestID, Target: target, Budget: budget, Steps: steps}
	p.SummaryHash = p.hash()
	return p
}

func (p *Plan) hash() string {
	h := sha256.New()
	raw, _ := json.Marshal(p.Steps)
	_, _ = h.Write(raw)
