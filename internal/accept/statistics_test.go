
package accept

import "testing"

func TestWindowRate(t *testing.T) {
	w := NewWindow(8)
	for i := 0; i < 6; i++ {
