
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
