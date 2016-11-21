
package admission

import (
	"testing"

	"github.com/xMinCHsia/marlin/internal/config"
	"github.com/xMinCHsia/marlin/internal/kvcache"
)

func TestAdmitWithinBudget(t *testing.T) {
	tr := kvcache.NewTracker(1000)
	c := NewController(config.Admission{HeadroomRatio: 0.2, ProbeIntervalS: 10}, tr)
	ok, _ := c.Admit("d1", 100)
	if !ok {
		t.Fatal("expected admission within budget")
	}
	c.Release(100)
}

func TestAdmitRejectsOverBudget(t *testing.T) {
	tr := kvcache.NewTracker(1000)
	c := NewController(config.Admission{HeadroomRatio: 0.2, ProbeIntervalS: 10}, tr)
	if ok, _ := c.Admit("d1", 900); ok {
		t.Fatal("900 of 800 usable should be rejected")
	}
	if ok, _ := c.Admit("d1", 900); ok {
		t.Fatal("expected rejection")
	}
}
