package arrayutils

func Int64sRangeClosed(start, end int64) []int64 {
	return Iterating(start, func(v int64) bool { return v <= end }, func(v int64) int64 { return v + 1 })
}

func Int64sRange(start, end int64) []int64 {
	return Iterating(start, func(v int64) bool { return v < end }, func(v int64) int64 { return v + 1 })
}

func Int32sRangeClosed(start, end int32) []int32 {
	return Iterating(start, func(v int32) bool { return v <= end }, func(v int32) int32 { return v + 1 })
}

func Int32sRange(start, end int32) []int32 {
	return Iterating(start, func(v int32) bool { return v < end }, func(v int32) int32 { return v + 1 })
}

func Int16sRangeClosed(start, end int16) []int16 {
	return Iterating(start, func(v int16) bool { return v <= end }, func(v int16) int16 { return v + 1 })
}

func Int16sRange(start, end int16) []int16 {
	return Iterating(start, func(v int16) bool { return v < end }, func(v int16) int16 { return v + 1 })
}

func IntsRangeClosed(start, end int) []int {
	return Iterating(start, func(v int) bool { return v <= end }, func(v int) int { return v + 1 })
}

func IntsRange(start, end int) []int {
	return Iterating(start, func(v int) bool { return v < end }, func(v int) int { return v + 1 })
}

func Iterating[V interface{}](initial V, hasNext func(v V) bool, next func(v V) V) []V {
	var newArray []V
	for position := initial; hasNext(position); position = next(position) {
		newArray = append(newArray, position)
	}
	return newArray
}
