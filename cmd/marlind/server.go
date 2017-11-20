// Package main: HTTP surface for the orchestrator.
package main

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/xMinCHsia/marlin/internal/admission"
	"github.com/xMinCHsia/marlin/internal/kvcache"
	"github.com/xMinCHsia/marlin/internal/tuning"
)

// server exposes the operator and plan endpoints.
type server struct {
	cfg         configLike
	controller  *admission.Controller
	autotuner   *tuning.Autotuner
	tracker     *kvcache.Tracker
	mu          sync.Mutex
	planHashes  []string
	started     time.Time
}

