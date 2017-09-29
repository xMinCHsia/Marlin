
package plan

import "testing"

func TestBuildDeterministic(t *testing.T) {
	steps := []Step{{DraftID: "d2", Tokens: []string{"a"}, Len: 1},
		{DraftID: "d1", Tokens: []string{"b"}, Len: 1}}
