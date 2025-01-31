package tslice_test

import (
	"errors"
	"testing"

	"github.com/poteto-go/tslice"
)

func TestAt(t *testing.T) {
	tests := []struct {
		name     string
		targets  []int
		index    int
		expected int
	}{
		{"Test success At plus index", []int{1, 2, 3, 4}, 1, 2},
		{"Test success At minus index", []int{1, 2, 3, 4}, -1, 4},
	}

	for _, it := range tests {
		t.Run(it.name, func(t *testing.T) {
			result := tslice.At(it.targets, it.index)
			if result != it.expected {
				t.Errorf("unmatched | actual(%d) - expected(%d)", result, it.expected)
			}
		})
	}
}

func TestAtPanicCase(t *testing.T) {
	tests := []struct {
		name    string
		targets []int
		index   int
	}{
		{"Test panic index >= length", []int{1, 2, 3, 4}, 4},
		{"Test panic index < -length", []int{1, 2, 3, 4}, -5},
	}

	for _, it := range tests {
		t.Run(it.name, func(t *testing.T) {
			var err error
			defer func() {
				if rec := recover(); rec != nil {
					err = errors.New("error")
				}
			}()

			tslice.At(it.targets, it.index)

			if !errors.Is(err, errors.New("error")) {
				t.Error("unmatched not panic")
			}
		})
	}
}

func TestConcat(t *testing.T) {
	tests := []struct {
		name     string
		targets  []int
		targets2 []int
		expected []int
	}{
		{"Test success concat", []int{1, 2}, []int{3, 4}, []int{1, 2, 3, 4}},
	}

	for _, it := range tests {
		t.Run(it.name, func(t *testing.T) {
			concat := tslice.Concat(it.targets, it.targets2)

			for i := range it.expected {
				if concat[i] != it.expected[i] {
					t.Error("unmatched")
				}
			}
		})
	}
}

func TestFill(t *testing.T) {
	tests := []struct {
		name     string
		targets  []int
		mask     int
		from     int
		to       int
		expected []int
	}{
		{"Test success fill", []int{0, 1, -1, 2}, 5, 0, 4, []int{5, 5, 5, 5}},
		{"Test success fill from to", []int{0, 1, -1, 2}, 5, 1, 3, []int{0, 5, 5, 2}},
		{"Test just from case", []int{0, 1, -1, 2}, 5, 2, -1, []int{0, 1, 5, 5}},
		{"Test all fill case", []int{0, 1, -1, 2}, 5, -1, -1, []int{5, 5, 5, 5}},
	}

	for _, it := range tests {
		t.Run(it.name, func(t *testing.T) {
			filled := make([]int, 0)
			if it.to >= 0 {
				filled = tslice.Fill(it.targets, it.mask, it.from, it.to)
			} else if it.from >= 0 {
				filled = tslice.Fill(it.targets, it.mask, it.from)
			} else {
				filled = tslice.Fill(it.targets, it.mask)
			}

			for i := range filled {
				if filled[i] != it.expected[i] {
					t.Error("unmatched")
				}
			}
		})
	}
}

func TestFillPanicCase(t *testing.T) {
	tests := []struct {
		name    string
		targets []int
		mask    int
		from    int
		to      int
		add     int
	}{
		{"Test panic if args > 2", []int{1, 2}, 1, 1, 1, 1},
		{"Test panic if from >= to", []int{1, 2, 3}, 1, 2, 1, -1},
		{"Test panic if from <= -1", []int{1, 2}, 1, -1, 1, -1},
		{"Test panic if to <= 0", []int{1, 2}, 1, -1, 0, -1},
		{"Test panic if from >= length", []int{1}, 1, 2, 2, -1},
	}

	for _, it := range tests {
		t.Run(it.name, func(t *testing.T) {
			var err error
			defer func() {
				if rec := recover(); rec != nil {
					err = errors.New("error")
				}
			}()

			if it.add >= 0 {
				tslice.Fill(it.targets, it.mask, it.from, it.to, it.add)
			} else {
				tslice.Fill(it.targets, it.mask, it.from, it.to)
			}

			if !errors.Is(err, errors.New("error")) {
				t.Error("unmatched not panic")
			}
		})
	}
}

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
