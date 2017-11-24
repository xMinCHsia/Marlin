
package plan

import "testing"

func TestBuildDeterministic(t *testing.T) {
	steps := []Step{{DraftID: "d2", Tokens: []string{"a"}, Len: 1},
		{DraftID: "d1", Tokens: []string{"b"}, Len: 1}}
	p1 := Build("r1", "t", 2, steps)
	p2 := Build("r1", "t", 2, steps)
	if p1.SummaryHash != p2.SummaryHash {
		t.Fatal("same input must produce the same hash")
	}
}
