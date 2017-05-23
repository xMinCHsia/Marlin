
// Package tuning self-adjusts draft lengths from acceptance statistics.
package tuning

import (
	"sync"

	"github.com/xMinCHsia/marlin/internal/accept"
	"github.com/xMinCHsia/marlin/internal/config"
)

// Autotuner moves draft lengths up when acceptance is high and down when
