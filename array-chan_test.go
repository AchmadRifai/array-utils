package arrayutils

import (
	"testing"
)

func TestReduceChan(t *testing.T) {
	result := ReduceChan(ArrayToChan(IntsRangeClosed(0, 2)), 0, func(acc, v int) int { return acc + v })
	if result != 3 {
		t.Error("Error! result must be 3")
	}
}

func TestFlatMapChan(t *testing.T) {
	result := ChanToArray(FlatMapChan(ArrayToChan(IntsRangeClosed(1, 3)), func(v int) <-chan int { return IntsChanClosed(0, v) }))
	if 9 != len(result) {
		t.Error("Error! result size must be 9")
	}
}

func TestCombineChan(t *testing.T) {
	a1, a2 := IntsRangeClosed(1, 2), IntsRangeClosed(1, 3)
	result := ChanToArray(CombineChan(ArrayToChan(a1), ArrayToChan(a2)))
	if 5 != len(result) {
		t.Error("Error! result size must be 5")
	}
}

func TestFilterChan(t *testing.T) {
	result := ChanToArray(FilterChan(ArrayToChan(IntsRangeClosed(1, 3)), func(v int) bool { return v%2 == 1 }))
	if 2 != len(result) {
		t.Error("Error! result size must be 2")
	}
}

func TestMapChan(t *testing.T) {
	chanIn := make(chan int)
	resultChan := MapChan(chanIn, func(v int) bool { return v%2 == 0 })
	chanIn <- 3
	close(chanIn)
	if <-resultChan {
		t.Error("Error! result must be false")
	}
}

func TestChanToArray(t *testing.T) {
	source := IntsRange(1, 2)
	result := ChanToArray(ArrayToChan(source))
	if NoneOf(result, func(v int, _ int) bool { return Contains(source, v) }) {
		t.Error("Error! source and result must be same")
	}
}

func TestArrayToChan(t *testing.T) {
	result := ArrayToChan(IntsRangeClosed(1, 2))
	val1, val2 := <-result, <-result
	if val1 == val2 {
		t.Error("Error! val1 must be 1 and val2 must be 2")
	}
}
