
package accept

import "sync"

// Window keeps a sliding window of acceptance events per draft model.
type Window struct {
	mu     sync.Mutex
	size   int
