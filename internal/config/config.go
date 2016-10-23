
// Package config loads and validates the Marlin configuration document.
package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Target holds the target-engine binding.
type Target struct {
	Endpoint   string `yaml:"endpoint"`
	MaxKVBytes int64  `yaml:"max_kv_bytes"`
}

// DraftModel describes one draft model in the ensemble.
type DraftModel struct {
	ID              string  `yaml:"id"`
	Endpoint        string  `yaml:"endpoint"`
	MaxProposalLen  int     `yaml:"max_proposal_len"`
	Weight          float64 `yaml:"weight"`
}

// Admission holds the KV-cache admission policy.
type Admission struct {
	HeadroomRatio   float64 `yaml:"headroom_ratio"`
	ProbeIntervalS  int     `yaml:"probe_interval_s"`
}

// Tuning holds the draft-length autotuner settings.
type Tuning struct {
	Enabled     bool `yaml:"enabled"`
	MinDraftLen int  `yaml:"min_draft_len"`
	MaxDraftLen int  `yaml:"max_draft_len"`
	Step        int  `yaml:"step"`
}

// Config is the root document.
type Config struct {
	ListenAddr string        `yaml:"listen_addr"`
	Target     Target        `yaml:"target"`
	Drafts     []DraftModel  `yaml:"drafts"`
	Admission  Admission     `yaml:"admission"`
	Tuning     Tuning        `yaml:"tuning"`
}

