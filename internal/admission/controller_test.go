
package admission

import (
	"testing"

	"github.com/xMinCHsia/marlin/internal/config"
	"github.com/xMinCHsia/marlin/internal/kvcache"
)

func TestAdmitWithinBudget(t *testing.T) {
	tr := kvcache.NewTracker(1000)
	c := NewController(config.Admission{HeadroomRatio: 0.2, ProbeIntervalS: 10}, tr)
