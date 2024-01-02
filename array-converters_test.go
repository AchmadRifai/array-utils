package arrayutils

import (
	"strconv"
	"testing"
)

type Bahan struct {
	Nilai int
	Team  string
}

func TestFlatMap(t *testing.T) {
	bahan := Map(IntsRangeClosed(1, 10), func(v int, i int) []int {
		var r []int
		r = append(r, v)
		return r
	})
	result := FlatMap(bahan, func(i1 []int, i2 int) []int { return i1 })
	if Contains(result, 0) {
		t.Errorf("Error! 0 is exclude")
	}
}

func TestGrouping(t *testing.T) {
	bahan := Map(IntsRangeClosed(1, 10), func(v int, i int) Bahan {
		var b Bahan
		if v%2 == 0 {
			b.Team = "Genap"
		} else {
			b.Team = "Ganjil"
		}
		b.Nilai = v
		return b
	})
	result := Grouping(bahan, func(b Bahan, i int) string { return b.Team })
	if 2 != len(result) {
		t.Errorf("Error! map size must be 2")
	}
}

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
