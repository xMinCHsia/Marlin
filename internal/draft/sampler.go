
package draft

import "math/rand"

// Sampler draws token proposals from a draft model's logits-like scores.
// In production the daemon receives token ids from the draft endpoint;
