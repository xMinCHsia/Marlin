
package draft

import (
	"math"
	"math/rand"
)

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

// Sample returns a token id chosen from weighted scores, or -1 when the
// candidate list is empty. Non-finite scores (NaN, infinities) are treated
// as zero so a broken draft cannot poison the draw.
func (s *Sampler) Sample(scores []float64) int {
	if len(scores) == 0 {
		return -1
	}
	total := 0.0
	for _, v := range scores {
		total += clampScore(v)
	}
	if total <= 0 {
		return 0
	}
	r := s.rng.Float64() * total
	acc := 0.0
	for i, v := range scores {
		acc += clampScore(v)
		if r < acc {
			return i
		}
	}
	return len(scores) - 1
}

// clampScore floors negative and non-finite logits at zero so a rejected
// token never steals probability mass from the live candidates.
func clampScore(v float64) float64 {
	if math.IsNaN(v) || math.IsInf(v, 0) || v < 0 {
		return 0
	}
	return v
}
