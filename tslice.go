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

// NOTE: returned 0 value if not found
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
