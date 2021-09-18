package woklang

/*
func Reduce(items []WokData, reducer func(WokData, WokData, int, []WokData) WokData, initial WokData) WokData {
	accumulator := initial

	for index, item := range items {
		accumulator = reducer(accumulator, item, index, items)
	}
	return accumulator
}

func RReduce(items []interface{}, reducer func(interface{}, interface{}, int, []interface{}) interface{}, initial interface{}) interface{} {
	accumulator := initial

	for index, item := range items {
		accumulator = reducer(accumulator, item, index, items)
	}
	return accumulator
}
*/
func MathReduce(items []WokData, reducer func(int64, WokData, int) int64, initial int64) int64 {
	accumulator := initial

	for index, item := range items {
		accumulator = reducer(accumulator, item, index)
	}
	return accumulator
}

/*
func Any(vs []WokData, f func(WokData) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func All(vs []WokData, f func(WokData) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(vs []WokData, f func(WokData) bool) []WokData {
	vsf := make([]WokData, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func Map(vs []WokData, f func(WokData) WokData) []WokData {
	vsm := make([]WokData, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
*/
