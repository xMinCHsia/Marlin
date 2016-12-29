
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
