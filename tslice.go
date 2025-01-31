package tslice

import "slices"

func At[V any](dataArray []V, index int) V {
	if index >= len(dataArray) || index < -len(dataArray) {
		panic("should be index < len(dataArray)")
	}

	if index >= 0 {
		return dataArray[index]
	}

	return dataArray[len(dataArray)+index]
}

func Concat[V any](dataArray []V, dataArray2 []V) []V {
	return append(dataArray, dataArray2...)
}

func CopyWithin[V any](dataArray []V, startIndex int, args ...int) []V {
	if len(args) > 2 || len(args) == 0 {
		panic("should be 0 < args < 2")
	}

	if startIndex < 0 || startIndex >= len(dataArray) {
		panic("should be 0 <= startIndex < len(dataArray)")
	}

	copied := make([]V, 0)
	if len(dataArray) == 0 {
		return copied
	}

	from, to := 0, len(dataArray)
	if len(args) >= 1 {
		from = args[0]
		if len(args) == 2 {
			to = args[1]
		}
	}

	if from >= to {
		panic("should be from < to")
	}

	if from <= -1 || to <= 0 {
		panic("should be from >= 0, to >= 1")
	}

	if from >= len(dataArray) || to > len(dataArray) {
		panic("should be from < len(dataArray) && to <= len(dataArray)")
	}

	copied = append(copied, dataArray...)

	targetIndex := startIndex
	for m := from; m < to; m++ {
		mask := dataArray[m]

		copied[targetIndex] = mask

		targetIndex++
		if targetIndex >= len(dataArray) {
			break
		}
	}

	return copied
}

func Every[V any](dataArray []V, handler func(data V) bool) bool {
	if len(dataArray) == 0 {
		return true
	}

	for _, data := range dataArray {
		if !handler(data) {
			return false
		}
	}

	return true
}

func Fill[V any](dataArray []V, mask V, args ...int) []V {
	if len(args) > 2 {
		panic("should be args <= 2")
	}

	filled := make([]V, 0)
	from, to := 0, len(dataArray)
	if len(args) >= 1 {
		from = args[0]
		if len(args) == 2 {
			to = args[1]
		}
	}

	if from >= to {
		panic("should be from < to")
	}

	if from <= -1 || to <= 0 {
		panic("should be from >= 0, to >= 1")
	}

	if from >= len(dataArray) || to > len(dataArray) {
		panic("should be from < len(dataArray) && to <= len(dataArray)")
	}

	for i := 0; i < len(dataArray); i++ {
		if i >= from && i < to {
			filled = append(filled, mask)
			continue
		}

		filled = append(filled, dataArray[i])
	}

	return filled
}

func Filter[V any](dataArray []V, handler func(data V) bool) []V {
	filtered := make([]V, 0)
	if len(dataArray) == 0 {
		return filtered
	}

	for _, data := range dataArray {
		if handler(data) {
			filtered = append(filtered, data)
		}
	}

	return filtered
}

// NOTE: return (0, false) if not found
func Find[V any](dataArray []V, handler func(data V) bool) (found V, ok bool) {
	if len(dataArray) == 0 {
		return found, ok
	}

	for _, data := range dataArray {
		if handler(data) {
			return data, true
		}
	}

	return found, ok
}

// NOTE: returned -1 if not found
func FindIndex[V any](dataArray []V, handler func(data V) bool) int {
	if len(dataArray) == 0 {
		return -1
	}

	for i, data := range dataArray {
		if handler(data) {
			return i
		}
	}

	return -1
}

// NOTE: return (0, false) if not found
func FindLast[V any](dataArray []V, handler func(data V) bool) (found V, ok bool) {
	if len(dataArray) == 0 {
		return found, ok
	}

	for i := len(dataArray) - 1; i >= 0; i-- {
		if handler(dataArray[i]) {
			return dataArray[i], true
		}
	}

	return found, ok
}

// NOTE: returned -1 if not found
func FindLastIndex[V any](dataArray []V, handler func(data V) bool) int {
	if len(dataArray) == 0 {
		return -1
	}

	for i := len(dataArray) - 1; i >= 0; i-- {
		if handler(dataArray[i]) {
			return i
		}
	}

	return -1
}

func Foreach[V any](dataArray []V, handler func(data V)) {
	for _, data := range dataArray {
		handler(data)
	}
}

// NOTE: Internal call slices.Contains
func Includes[E comparable](dataArray []E, data E) bool {
	return slices.Contains(dataArray, data)
}

// return -1 if not found
func IndexOf[E comparable](dataArray []E, data E, startIndex ...int) int {
	if len(startIndex) > 1 {
		panic("should be len(startIndex) <= 1")
	}

	from := 0
	if len(startIndex) == 1 {
		from = startIndex[0]
	}

	if from >= len(dataArray) || from < 0 {
		panic("should be 0 <= startIndex < len(dataArray)")
	}

	for i := from; i < len(dataArray); i++ {
		if dataArray[i] == data {
			return i
		}
	}

	return -1
}

func Map[V any](dataArray []V, handler func(data V) V) []V {
	transformed := make([]V, 0)
	if len(dataArray) == 0 {
		return transformed
	}

	for _, data := range dataArray {
		transformed = append(transformed, handler(data))
	}
	return transformed
}
