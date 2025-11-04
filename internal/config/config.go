
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

// Load reads and parses the YAML document.
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", path, err)
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parse %s: %w", path, err)
	}
	cfg.ApplyDefaults()
	return &cfg, nil
}

// ApplyDefaults fills optional fields that are absent from a minimal config
// with the documented defaults, so operators can omit everything that has a
// sane fallback and still boot successfully.
func (c *Config) ApplyDefaults() {
	if c.ListenAddr == "" {
		c.ListenAddr = ":8590"
	}
	if c.Target.MaxKVBytes <= 0 {
		c.Target.MaxKVBytes = 8 << 30
	}
	if c.Admission.HeadroomRatio <= 0 {
		c.Admission.HeadroomRatio = 0.2
	}
	if c.Admission.ProbeIntervalS <= 0 {
		c.Admission.ProbeIntervalS = 10
	}
	if c.Tuning.MinDraftLen <= 0 {
		c.Tuning.MinDraftLen = 2
	}
	if c.Tuning.MaxDraftLen <= 0 {
		c.Tuning.MaxDraftLen = 8
	}
	if c.Tuning.Step <= 0 {
		c.Tuning.Step = 1
	}
	for i := range c.Drafts {
		if c.Drafts[i].MaxProposalLen <= 0 {
			c.Drafts[i].MaxProposalLen = 8
		}
		if c.Drafts[i].Weight <= 0 {
			c.Drafts[i].Weight = 1.0
		}
	}
}

// Validate reports configuration problems at boot.
func (c *Config) Validate() error {
	if c.ListenAddr == "" {
		return fmt.Errorf("listen_addr is required")
	}
	if c.Target.Endpoint == "" {
		return fmt.Errorf("target.endpoint is required")
	}
	if c.Target.MaxKVBytes <= 0 {
		return fmt.Errorf("target.max_kv_bytes must be positive")
	}
	if len(c.Drafts) == 0 {
		return fmt.Errorf("at least one draft model is required")
	}
	seen := map[string]bool{}
	for _, dr := range c.Drafts {
		if dr.ID == "" || dr.Endpoint == "" {
			return fmt.Errorf("each draft needs id and endpoint")
		}
		if seen[dr.ID] {
			return fmt.Errorf("duplicate draft id %q", dr.ID)
		}
		seen[dr.ID] = true
		if dr.MaxProposalLen < 1 || dr.MaxProposalLen > 64 {
			return fmt.Errorf("draft %q: max_proposal_len must be 1..64", dr.ID)
		}
		if dr.Weight <= 0 {
			return fmt.Errorf("draft %q: weight must be positive", dr.ID)
		}
	}
	if c.Admission.HeadroomRatio < 0 || c.Admission.HeadroomRatio > 0.9 {
		return fmt.Errorf("admission.headroom_ratio must be 0..0.9")
	}
	if c.Tuning.Enabled && (c.Tuning.MinDraftLen < 1 || c.Tuning.MaxDraftLen < c.Tuning.MinDraftLen) {
		return fmt.Errorf("tuning bounds are invalid")
	}
	return nil
}
