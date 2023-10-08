package arrayutils

import (
	"github.com/golang-collections/collections/stack"
)

func Sort[V interface{}](array []V, compare func(v1, v2 V) int) {
	quickSort(array, 0, len(array)-1, compare)
}

func quickSort[V interface{}](arr []V, low, high int, compare func(v1, v2 V) int) {
	tumpukan := stack.New()
	tumpukan.Push(low)
	tumpukan.Push(high)
	for tumpukan.Len() > 0 {
		h, l, p := tumpukan.Pop().(int), tumpukan.Pop().(int), 0
		arr, p = partition(arr, l, h, compare)
		if p-1 > l {
			tumpukan.Push(l)
			tumpukan.Push(p - 1)
		}
		if p+1 < h {
			tumpukan.Push(p + 1)
			tumpukan.Push(h)
		}
	}
}

func partition[V interface{}](arr []V, low, high int, compare func(v1, v2 V) int) ([]V, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if compare(arr[j], pivot) == 1 {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}
