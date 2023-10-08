package arrayutils

import "testing"

func TestInt64sRangeClosed(t *testing.T) {
	result := Int64sRangeClosed(0, 6)
	if 7 != len(result) {
		t.Errorf("Error! Must be 7")
	}
}

func TestInt64sRange(t *testing.T) {
	result := Int64sRange(0, 6)
	if 6 != len(result) {
		t.Errorf("Error! Must be 6")
	}
}

func TestInt32sRangeClosed(t *testing.T) {
	result := Int32sRangeClosed(0, 6)
	if 7 != len(result) {
		t.Errorf("Error! Must be 7")
	}
}

func TestInt32sRange(t *testing.T) {
	result := Int32sRange(0, 6)
	if 6 != len(result) {
		t.Errorf("Error! Must be 6")
	}
}

func TestInt16sRangeClosed(t *testing.T) {
	result := Int16sRangeClosed(0, 6)
	if 7 != len(result) {
		t.Errorf("Error! Must be 7")
	}
}

func TestInt16sRange(t *testing.T) {
	result := Int16sRange(0, 6)
	if 6 != len(result) {
		t.Errorf("Error! Must be 6")
	}
}

func TestIntsRangeClosed(t *testing.T) {
	result := IntsRangeClosed(0, 6)
	if 7 != len(result) {
		t.Errorf("Error! Must be 7")
	}
}

func TestIntsRange(t *testing.T) {
	result := IntsRange(0, 6)
	if 6 != len(result) {
		t.Errorf("Error! Must be 6")
	}
}

func TestIterating(t *testing.T) {
	start, end := 0, 6
	t.Logf("Start %d until %d", start, end)
	result := Iterating(start, func(v int) bool { return v < end }, func(v int) int { return v + 1 })
	if 6 != len(result) {
		t.Errorf("Error! Must be 6")
	}
}
