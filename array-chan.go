package arrayutils

import (
	"sync"
	"sync/atomic"
)

func ReduceChan[T interface{}](chanT <-chan T, init T, combine func(acc, v T) T) T {
	result := init
	for v := range chanT {
		result = combine(result, v)
	}
	return result
}

func FlatMapChan[T, S interface{}](chanT <-chan T, convert func(t T) <-chan S) <-chan S {
	chanOut := make(chan S)
	var counter atomic.Int32
	go func() {
		var wg sync.WaitGroup
		for t := range chanT {
			wg.Add(1)
			counter.Add(1)
			go func(t T) {
				defer func() {
					counter.Store(counter.Load() - 1)
					wg.Done()
				}()
				for s := range convert(t) {
					chanOut <- s
				}
			}(t)
			for 20 <= counter.Load() {
			}
		}
		wg.Wait()
		close(chanOut)
	}()
	return chanOut
}

func CombineChan[T interface{}](chan1, chan2 <-chan T) <-chan T {
	chanOut := make(chan T)
	go func() {
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			defer wg.Done()
			for t := range chan1 {
				chanOut <- t
			}
		}()
		go func() {
			defer wg.Done()
			for t := range chan2 {
				chanOut <- t
			}
		}()
		wg.Wait()
		close(chanOut)
	}()
	return chanOut
}

func FilterChan[V interface{}](chanV <-chan V, predicate func(v V) bool) <-chan V {
	chanOut := make(chan V)
	go func() {
		defer close(chanOut)
		for v := range chanV {
			if predicate(v) {
				chanOut <- v
			}
		}
	}()
	return chanOut
}

func MapChan[V, S interface{}](chanV <-chan V, convert func(v V) S) <-chan S {
	chanOut := make(chan S)
	go func() {
		defer close(chanOut)
		for v := range chanV {
			chanOut <- convert(v)
		}
	}()
	return chanOut
}

func ChanToArray[T interface{}](chanT <-chan T) []T {
	var array []T
	for v := range chanT {
		array = append(array, v)
	}
	return array
}

func ArrayToChan[T interface{}](array []T) <-chan T {
	chanOut := make(chan T)
	go func(array []T) {
		defer close(chanOut)
		for _, v := range array {
			chanOut <- v
		}
	}(array)
	return chanOut
}
