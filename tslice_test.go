package tslice_test

import (
	"testing"

	"github.com/poteto-go/tslice"
)

func TestFilter(t *testing.T) {
	tests := []struct {
		name         string
		targets      []int
		expectLength int
	}{
		{"Test success filter", []int{0, 1, -1, 2}, 3},
		{"Test 0 length case", []int{}, 0},
	}

	for _, it := range tests {
		t.Run(it.name, func(t *testing.T) {
			filtered := tslice.Filter(it.targets, func(data int) bool {
				return data >= 0
			})

			if len(filtered) != it.expectLength {
				t.Errorf("unmatched filter: %v", filtered)
			}
		})
	}
}

func TestFind(t *testing.T) {
	tests := []struct {
		name     string
		targets  []int
		expected int
	}{
		{"Test success find", []int{0, 1, -1, 2}, 1},
		{"Test 0 length case", []int{}, 0},
		{"Test nothing clear rule", []int{-1}, 0},
	}

	for _, it := range tests {
		t.Run(it.name, func(t *testing.T) {
			found, _ := tslice.Find(it.targets, func(data int) bool {
				return data >= 1
			})

			if found != it.expected {
				t.Errorf("unmatched | actual(%d) - expected(%d)", found, it.expected)
			}
		})
	}
}

func TestFindIndex(t *testing.T) {
	tests := []struct {
		name     string
		targets  []int
		expected int
	}{
		{"Test success findIndex", []int{0, 1, -1, 2}, 1},
		{"Test 0 length case", []int{}, -1},
		{"Test nothing clear rule", []int{-1}, -1},
	}

	for _, it := range tests {
		t.Run(it.name, func(t *testing.T) {
			found := tslice.FindIndex(it.targets, func(data int) bool {
				return data >= 1
			})

			if found != it.expected {
				t.Errorf("unmatched | actual(%d) - expected(%d)", found, it.expected)
			}
		})
	}
}

func TestFindLast(t *testing.T) {
	tests := []struct {
		name     string
		targets  []int
		expected int
	}{
		{"Test success find last one", []int{0, 1, -1, 2}, 2},
		{"Test 0 length case", []int{}, 0},
		{"Test nothing clear rule", []int{-1}, 0},
	}

	for _, it := range tests {
		t.Run(it.name, func(t *testing.T) {
			found, _ := tslice.FindLast(it.targets, func(data int) bool {
				return data >= 1
			})

			if found != it.expected {
				t.Errorf("unmatched | actual(%d) - expected(%d)", found, it.expected)
			}
		})
	}
}

func TestFindLastIndex(t *testing.T) {
	tests := []struct {
		name     string
		targets  []int
		expected int
	}{
		{"Test success findLastIndex", []int{0, 1, -1, 2}, 3},
		{"Test 0 length case", []int{}, -1},
		{"Test nothing clear rule", []int{-1}, -1},
	}

	for _, it := range tests {
		t.Run(it.name, func(t *testing.T) {
			found := tslice.FindLastIndex(it.targets, func(data int) bool {
				return data >= 1
			})

			if found != it.expected {
				t.Errorf("unmatched | actual(%d) - expected(%d)", found, it.expected)
			}
		})
	}
}

func TestMap(t *testing.T) {
	tests := []struct {
		name     string
		targets  []int
		expected []int
	}{
		{"Test success *2 case", []int{0, 1, 2, 3}, []int{0, 2, 4, 6}},
		{"Test 0 length case", []int{}, []int{}},
	}

	for _, it := range tests {
		t.Run(it.name, func(t *testing.T) {
			transformed := tslice.Map(it.targets, func(data int) int {
				return data * 2
			})

			for i := range transformed {
				if transformed[i] != it.expected[i] {
					t.Error("unmatched")
				}
			}
		})
	}
}
