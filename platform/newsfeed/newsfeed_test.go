package newsfeed

import "testing"

func TestHalo(t *testing.T) {
	got := 2
	if got != 1 {
		t.Errorf("Halo(-1) = %d; want 1", got)
	}
}
