
// Package accept implements rejection sampling bookkeeping.
package accept

// Verdict for one proposed token.
const (
	Accepted = "accepted"
	Rejected = "rejected"
)

// Rejector reports which proposed tokens the target accepted.
type Rejector struct {
	accepted map[string]int
	rejected map[string]int
}

// NewRejector builds the per-draft counters.
func NewRejector() *Rejector {
	return &Rejector{
		accepted: map[string]int{},
		rejected: map[string]int{},
	}
}

// Record registers one token verdict for a draft.
func (r *Rejector) Record(draftID, verdict string) {
	if verdict == Accepted {
		r.accepted[draftID]++
	} else {
		r.rejected[draftID]++
	}
}

// Accepted returns the accepted count for a draft.
func (r *Rejector) Accepted(draftID string) int {
	return r.accepted[draftID]
}

// Rejected returns the rejected count for a draft.
func (r *Rejector) Rejected(draftID string) int {
	return r.rejected[draftID]
}

// Rate returns the acceptance rate (0..1) for a draft.
