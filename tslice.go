package tslice

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

// NOTE: return 0 value if not found
func Find[V any](dataArray []V, handler func(data V) bool) (found V) {
	if len(dataArray) == 0 {
		return found
	}

	for _, data := range dataArray {
		if handler(data) {
			return data
		}
	}

	return found
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

// NOTE: return 0 value if not found
func FindLast[V any](dataArray []V, handler func(data V) bool) (found V) {
	if len(dataArray) == 0 {
		return found
	}

	for i := len(dataArray) - 1; i >= 0; i-- {
		if handler(dataArray[i]) {
			return dataArray[i]
		}
	}

	return found
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
