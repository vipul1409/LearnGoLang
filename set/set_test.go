package set

import (
	"testing"
)

func testAddItems(t *testing.T) {
	set := New()
	set.Add(1, 2, 3, 4, 5)
	arr := []int{1, 2, 3, 4, 5}
	for _, v := range arr {
		if !set.Exists(v) {
			t.Errorf("Items missing in set")
		}
	}
}
