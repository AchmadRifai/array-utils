package arrayutils

import (
	"strconv"
	"testing"
)

func TestDistinct(t *testing.T) {
	result := Distinct(Map(IntsRange(0, 6), func(v int, _ int) int { return 0 }))
	if 1 != len(result) {
		t.Errorf("Error! array size must be 1")
	}
}

func TestNoneOf(t *testing.T) {
	if !NoneOf(IntsRange(0, 6), func(v int, _ int) bool { return v > 6 }) {
		t.Errorf("Error! array all must be lest then 6")
	}
}

func TestAnyOf(t *testing.T) {
	if !AnyOf(IntsRange(0, 6), func(v int, _ int) bool { return v == 3 }) {
		t.Errorf("Error! array all must be contains 3")
	}
}

func TestAllOf(t *testing.T) {
	if !AllOf(IntsRange(0, 6), func(v int, _ int) bool { return v < 6 }) {
		t.Errorf("Error! array all must be lest then 6")
	}
}

func TestFilter(t *testing.T) {
	result := Filter(IntsRange(0, 6), func(v int, _ int) bool { return v%2 != 0 })
	if Contains(result, 2) {
		t.Errorf("Error! array not cantains 2")
	}
}

func TestMap(t *testing.T) {
	result := Map(Int64sRange(0, 6), func(v int64, _ int) string { return strconv.FormatInt(v, 10) })
	if Contains(result, "6") {
		t.Errorf("Error! array not cantains \"6\"")
	}
}

func TestContains(t *testing.T) {
	array := IntsRange(0, 6)
	if Contains(array, 6) {
		t.Errorf("Error! array not cantains 6")
	}
}

func TestAddAll(t *testing.T) {
	array1, array2 := IntsRange(0, 6), IntsRange(7, 10)
	result := AddAll(array1, array2)
	if len(result) != len(array1)+len(array2) {
		t.Errorf("Error! result must be %d", len(array1)+len(array2))
	}
}
