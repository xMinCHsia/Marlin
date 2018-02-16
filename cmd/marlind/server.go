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

type configLike interface {
	Validate() error
}

// newServer builds the HTTP surface.
func newServer(cfg configLike, c *admission.Controller, a *tuning.Autotuner, t *kvcache.Tracker) *server {
	return &server{cfg: cfg, controller: c, autotuner: a, tracker: t,
		started: time.Now(), planHashes: []string{}}
}

// Handler returns the route table.
func (s *server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", s.health)
	mux.HandleFunc("GET /v1/status", s.status)
	mux.HandleFunc("GET /v1/plans", s.plans)
	mux.HandleFunc("POST /v1/tune", s.tune)
	return mux
}

func (s *server) health(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, 200, map[string]any{
		"status": "ok", "uptime_s": int(time.Since(s.started).Seconds()),
	})
}

func (s *server) status(w http.ResponseWriter, _ *http.Request) {
