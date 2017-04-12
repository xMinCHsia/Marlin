
package draft

import "testing"

func TestSampleRespectsWeights(t *testing.T) {
	s := NewSampler(42)
	scores := []float64{0.1, 9.0}
	idx2 := 0
	for i := 0; i < 200; i++ {
		if s.Sample(scores) == 1 {
