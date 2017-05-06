
package draft

import "testing"

func TestSampleRespectsWeights(t *testing.T) {
	s := NewSampler(42)
	scores := []float64{0.1, 9.0}
	idx2 := 0
	for i := 0; i < 200; i++ {
		if s.Sample(scores) == 1 {
			idx2++
		}
	}
	if idx2 < 150 {
		t.Fatalf("high-weight token should dominate, got %d/200", idx2)
	}
}

func TestSampleZeroScores(t *testing.T) {
	s := NewSampler(1)
	if got := s.Sample([]float64{0, 0}); got != 0 {
		t.Fatalf("expected index 0 on zero scores, got %d", got)
	}
}

func TestSampleSingle(t *testing.T) {
