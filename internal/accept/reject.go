
// Package accept implements rejection sampling bookkeeping.
package accept

// Verdict for one proposed token.
const (
	Accepted = "accepted"
	Rejected = "rejected"
)

// Rejector reports which proposed tokens the target accepted.
