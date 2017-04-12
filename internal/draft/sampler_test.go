
package draft

import "testing"

func TestSampleRespectsWeights(t *testing.T) {
	s := NewSampler(42)
	scores := []float64{0.1, 9.0}
