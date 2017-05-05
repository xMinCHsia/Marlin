
package draft

import "math/rand"

// Sampler draws token proposals from a draft model's logits-like scores.
// In production the daemon receives token ids from the draft endpoint;
// this type keeps the ensemble logic testable without a live model.
type Sampler struct {
	rng *rand.Rand
}

// NewSampler builds a sampler with a local source.
func NewSampler(seed int64) *Sampler {
	return &Sampler{rng: rand.New(rand.NewSource(seed))}
}

// Sample returns a token id chosen from weighted scores.
