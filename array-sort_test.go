package arrayutils

import (
	"testing"
)

func TestSort(t *testing.T) {
	bahan := Map(IntsRange(0, 6), func(v int, _ int) int { return 6 - v })
	Sort(bahan, func(v1, v2 int) int {
		if v1 < v2 {
			return 1
		} else if v1 > v2 {
			return -1
		} else {
			return 0
		}
	})
	if bahan[0] != 1 {
		t.Errorf("First item must be 1")
	}
}
