package arrayutils

func NoneOf[V interface{}](array []V, predicate func(v V, index int) bool) bool {
	return !AnyOf(array, predicate)
}

func AnyOf[V interface{}](array []V, predicate func(v V, index int) bool) bool {
	for i, v := range array {
		if predicate(v, i) {
			return true
		}
	}
	return false
}

func AllOf[V interface{}](array []V, predicate func(v V, index int) bool) bool {
	for i, v := range array {
		if !predicate(v, i) {
			return false
		}
	}
	return true
}

func Filter[V interface{}](array []V, predicate func(v V, index int) bool) []V {
	var newArray []V
	for i, v := range array {
		if predicate(v, i) {
			newArray = append(newArray, v)
		}
	}
	return newArray
}

func Grouping[V interface{}, S comparable](array []V, classifier func(v V, index int) S) map[S][]V {
	result := make(map[S][]V)
	for _, v := range Distinct(Map(array, classifier)) {
		result[v] = Filter(array,
			func(d V, i int) bool {
				return classifier(d, i) == v
			})
	}
	return result
}

func FlatMap[V, S interface{}](source []V, convert func(V, int) []S) []S {
	var newArray []S
	for i, v := range source {
		newArray = AddAll(newArray, convert(v, i))
	}
	return newArray
}

func Map[V, S interface{}](array []V, convert func(v V, index int) S) []S {
	var newArray []S
	for i, v := range array {
		newArray = append(newArray, convert(v, i))
	}
	return newArray
}

func Distinct[V comparable](array []V) []V {
	var newArray []V
	for _, v := range array {
		if !Contains(newArray, v) {
			newArray = append(newArray, v)
		}
	}
	return newArray
}

func Contains[V comparable](array []V, v V) bool {
	for _, i := range array {
		if i == v {
			return true
		}
	}
	return false
}

func AddAll[V interface{}](array1, array2 []V) []V {
	var newArray []V
	newArray = append(newArray, array1...)
	newArray = append(newArray, array2...)
	return newArray
}
