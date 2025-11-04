
package accept

import "testing"

func TestWindowRate(t *testing.T) {
	w := NewWindow(8)
	for i := 0; i < 6; i++ {
		w.Record("d1", true)
	}
	for i := 0; i < 2; i++ {
		w.Record("d1", false)
	}
	if got := w.Rate("d1"); got != 0.75 {
		t.Fatalf("expected 0.75, got %v", got)
	}
}

func TestWindowDropsOldest(t *testing.T) {
	w := NewWindow(4)
	for i := 0; i < 6; i++ {
		w.Record("d1", true)
	}
	if got := w.Rate("d1"); got != 1.0 {
		t.Fatalf("expected 1.0 over the sliding window, got %v", got)
	}
}

func TestDriftDetected(t *testing.T) {
	w := NewWindow(8)
	for i := 0; i < 8; i++ {
		w.Record("d1", true)
	}
	base := w.Rate("d1")
	for i := 0; i < 8; i++ {
		w.Record("d1", false)
	}
	if !w.Drift("d1", base, 0.15) {
		t.Fatal("expected drift past threshold")
	}
}

func TestReset(t *testing.T) {
	w := NewWindow(8)
	w.Record("d1", true)
	w.Reset()
	if got := w.Rate("d1"); got != 0 {
		t.Fatalf("expected 0 after reset, got %v", got)
	}
}
