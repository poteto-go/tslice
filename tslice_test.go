package tslice_test

import (
	"testing"

	"github.com/poteto-go/tslice"
)

func TestFilter(t *testing.T) {
	dataArr := []int{0, 1, -1, 2}

	filtered := tslice.Filter(dataArr, func(data int) bool {
		return data >= 0
	})

	if len(filtered) != 3 {
		t.Errorf("Unmatched filter: %v", filtered)
	}
}
