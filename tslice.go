package tslice

import (
	"iter"
	"slices"
	"sort"

	"golang.org/x/exp/constraints"
)

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

// NOTE: internal call slices.All
func Entries[V any](dataArray []V) iter.Seq2[int, V] {
	return slices.All(dataArray)
}

func Every[V any](dataArray []V, yield func(data V) bool) bool {
	if len(dataArray) == 0 {
		return true
	}

	for _, data := range dataArray {
		if !yield(data) {
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

func Filter[V any](dataArray []V, yield func(data V) bool) []V {
	filtered := make([]V, 0)
	if len(dataArray) == 0 {
		return filtered
	}

	for _, data := range dataArray {
		if yield(data) {
			filtered = append(filtered, data)
		}
	}

	return filtered
}

// NOTE: return (0, false) if not found
func Find[V any](dataArray []V, yield func(data V) bool) (found V, ok bool) {
	if len(dataArray) == 0 {
		return found, ok
	}

	for _, data := range dataArray {
		if yield(data) {
			return data, true
		}
	}

	return found, ok
}

// NOTE: returned -1 if not found
func FindIndex[V any](dataArray []V, yield func(data V) bool) int {
	if len(dataArray) == 0 {
		return -1
	}

	for i, data := range dataArray {
		if yield(data) {
			return i
		}
	}

	return -1
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

// NOTE: return (0, false) if not found
func FindLast[V any](dataArray []V, yield func(data V) bool) (found V, ok bool) {
	if len(dataArray) == 0 {
		return found, ok
	}

	for i := len(dataArray) - 1; i >= 0; i-- {
		if yield(dataArray[i]) {
			return dataArray[i], true
		}
	}

	return found, ok
}

// NOTE: returned -1 if not found
func FindLastIndex[V any](dataArray []V, yield func(data V) bool) int {
	if len(dataArray) == 0 {
		return -1
	}

	for i := len(dataArray) - 1; i >= 0; i-- {
		if yield(dataArray[i]) {
			return i
		}
	}

	return -1
}

// return -1 if not found
func LastIndexOf[E comparable](dataArray []E, data E, startIndex ...int) int {
	if len(startIndex) > 1 {
		panic("should be len(startIndex) <= 1")
	}

	from := len(dataArray) - 1
	if len(startIndex) == 1 {
		from = startIndex[0]
	}

	if from >= len(dataArray) || from < 0 {
		panic("should be 0 <= startIndex < len(dataArray)")
	}

	for i := from; i >= 0; i-- {
		if dataArray[i] == data {
			return i
		}
	}

	return -1
}

func ForEach[V any](dataArray []V, yield func(data V)) {
	for _, data := range dataArray {
		yield(data)
	}
}

// NOTE: Internal call slices.Contains
func Includes[E comparable](dataArray []E, data E) bool {
	return slices.Contains(dataArray, data)
}

func Map[V any](dataArray []V, yield func(data V) V) []V {
	transformed := make([]V, 0)
	if len(dataArray) == 0 {
		return transformed
	}

	for _, data := range dataArray {
		transformed = append(transformed, yield(data))
	}
	return transformed
}

func Pop[V any](dataArray *[]V) V {
	if len(*dataArray) == 0 {
		panic("should be len(dataArray) >= 1")
	}

	lastValue := (*dataArray)[len(*dataArray)-1]

	newArr := make([]V, len(*dataArray)-1)
	copy(newArr, *dataArray)
	*dataArray = newArr

	return lastValue
}

func Push[V any](dataArray *[]V, add ...V) int {
	if len(add) == 0 {
		return len(*dataArray)
	}

	n := len(*dataArray)

	// 配列を拡張
	newArr := make([]V, n+len(add))
	copy(newArr, *dataArray)
	*dataArray = newArr

	for i := 0; i < len(add); i++ {
		(*dataArray)[n] = add[i]
		n++
	}
	return n
}

func Reduce[V any, T any](dataArray []V, yield func(acc T, cur V) T, acc0 ...T) T {
	if len(acc0) > 1 {
		panic("should be len(acc0) <= 1")
	}

	var start T
	if len(acc0) == 1 {
		start = acc0[0]
	}

	now := start
	for _, data := range dataArray {
		now = yield(now, data)
	}

	return now
}

func ReduceRight[V any, T any](dataArray []V, yield func(acc T, cur V) T, acc0 ...T) T {
	if len(acc0) > 1 {
		panic("should be len(acc0) <= 1")
	}

	var start T
	if len(acc0) == 1 {
		start = acc0[0]
	}

	now := start
	for i := len(dataArray) - 1; i >= 0; i-- {
		now = yield(now, dataArray[i])
	}

	return now
}

// * call w/ yield
// * if you want call w/o yield
// * use SortO instead
// this internal calls sort.SliceStable()
func Sort[V any](dataArray []V, yield func(left, right V) int) {
	if len(dataArray) == 0 {
		return
	}

	wrapper := func(i, j int) bool {
		left := dataArray[i]
		right := dataArray[j]

		result := yield(left, right)
		return result < 0
	}

	sort.SliceStable(dataArray, wrapper)
}

func ToSorted[V any](dataArray []V, yield func(left, right V) int) []V {
	newArray := make([]V, len(dataArray))
	copy(newArray, dataArray)

	Sort(newArray, yield)

	return newArray
}

func SortO[O constraints.Ordered](dataArray []O) {
	if len(dataArray) == 0 {
		return
	}

	sort.SliceStable(dataArray, func(i, j int) bool {
		return (dataArray)[i] < (dataArray)[j]
	})
}

func ToSortedO[O constraints.Ordered](dataArray []O) []O {
	newArray := make([]O, len(dataArray))
	copy(newArray, dataArray)

	SortO(newArray)

	return newArray
}

func ToString[V any](dataArray []V) string {
	if len(dataArray) == 0 {
		return ""
	}

	var text string
	for i, data := range dataArray {
		text += toString(data)
		if i != len(dataArray)-1 {
			text += ", "
		}
	}
	return text
}
