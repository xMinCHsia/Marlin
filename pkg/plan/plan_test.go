
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

func TestBuildSortsSteps(t *testing.T) {
	steps := []Step{{DraftID: "z", Tokens: []string{"a"}, Len: 1},
		{DraftID: "a", Tokens: []string{"b"}, Len: 1}}
	p := Build("r1", "t", 2, steps)
	if p.Steps[0].DraftID != "a" {
		t.Fatal("steps must be sorted by draft id")
	}
}

func TestToJSON(t *testing.T) {
	p := Build("r1", "t", 1, []Step{{DraftID: "d", Tokens: []string{"x"}, Len: 1}})
	raw, err := p.ToJSON()
	if err != nil || len(raw) == 0 {
		t.Fatalf("toJSON: %v", err)
	}
}
