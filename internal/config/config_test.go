
package config

import (
	"os"
	"path/filepath"
	"testing"
)

func write(t *testing.T, body string) string {
	p := filepath.Join(t.TempDir(), "m.yaml")
	if err := os.WriteFile(p, []byte(body), 0o600); err != nil {
		t.Fatal(err)
	}
	return p
}

func TestLoadValid(t *testing.T) {
	p := write(t, `
listen_addr: ":8590"
target:
  endpoint: "http://127.0.0.1:8000"
  max_kv_bytes: 1000000
drafts:
  - id: "d1"
    endpoint: "http://127.0.0.1:8101"
    max_proposal_len: 6
    weight: 1.0
admission:
  headroom_ratio: 0.2
  probe_interval_s: 10
`)
	cfg, err := Load(p)
	if err != nil {
		t.Fatal(err)
	}
	if err := cfg.Validate(); err != nil {
		t.Fatalf("validate: %v", err)
	}
}

func TestValidateRequiresDraft(t *testing.T) {
	p := write(t, `
listen_addr: ":8590"
target:
  endpoint: "http://x"
  max_kv_bytes: 1000
`)
	cfg, _ := Load(p)
	if err := cfg.Validate(); err == nil {
		t.Fatal("expected error for missing drafts")
	}
}

func TestValidateDuplicateDraftID(t *testing.T) {
	p := write(t, `
listen_addr: ":8590"
target:
  endpoint: "http://x"
  max_kv_bytes: 1000
drafts:
  - id: "d1"
    endpoint: "http://a"
    max_proposal_len: 4
    weight: 1
  - id: "d1"
    endpoint: "http://b"
    max_proposal_len: 4
    weight: 1
`)
	cfg, _ := Load(p)
	if err := cfg.Validate(); err == nil {
		t.Fatal("expected error for duplicate draft id")
	}
}

func TestApplyDefaults(t *testing.T) {
	p := write(t, `
target:
  endpoint: "http://127.0.0.1:8000"
drafts:
  - id: "d1"
    endpoint: "http://127.0.0.1:8101"
`)
	cfg, err := Load(p)
	if err != nil {
		t.Fatal(err)
	}
	if cfg.ListenAddr != ":8590" {
		t.Fatalf("expected default listen addr, got %q", cfg.ListenAddr)
	}
	if cfg.Target.MaxKVBytes != 8<<30 {
