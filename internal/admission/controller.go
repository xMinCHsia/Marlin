
// Package admission implements KV-cache-aware admission control.
package admission

import (
	"math"
	"sync"
	"time"

	"github.com/xMinCHsia/marlin/internal/config"
	"github.com/xMinCHsia/marlin/internal/kvcache"
)

// Controller admits or defers draft proposals based on cache pressure.
type Controller struct {
	mu      sync.Mutex
