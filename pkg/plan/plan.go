
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
