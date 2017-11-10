
// Command marlind runs the speculative decoding orchestrator.
package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/xMinCHsia/marlin/internal/admission"
	"github.com/xMinCHsia/marlin/internal/config"
	"github.com/xMinCHsia/marlin/internal/kvcache"
	"github.com/xMinCHsia/marlin/internal/tuning"
)

func main() {
	cfgPath := flag.String("config", "marlin.yaml", "path to config")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	if err != nil {
		log.Fatalf("config: %v", err)
	}
	if err := cfg.Validate(); err != nil {
		log.Fatalf("config invalid: %v", err)
	}

	tracker := kvcache.NewTracker(cfg.Target.MaxKVBytes)
	controller := admission.NewController(cfg.Admission, tracker)
	autotuner := tuning.NewAutotuner(cfg.Tuning)
