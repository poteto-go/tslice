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
