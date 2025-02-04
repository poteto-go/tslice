package tslice_test

import (
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

func TestCopyWithin(t *testing.T) {
	tests := []struct {
		name     string
		targets  []int
		start    int
		from     int
		to       int
		expected []int
	}{
		{"Test success from to CopyWithin single", []int{0, 1, -1, 2, -3, 4}, 1, 2, 3, []int{0, -1, -1, 2, -3, 4}},
		{"Test success from to CopyWithin multi", []int{0, 1, -1, 2, -3, 4}, 1, 3, 5, []int{0, 2, -3, 2, -3, 4}},
		{"Test success from to end", []int{0, 1, -1, 2, -3, 4}, 1, 3, -1, []int{0, 2, -3, 4, -3, 4}},
		{"Test success targetIndex > lenght", []int{0, 1, -1, 2, -3, 4}, 4, 3, -1, []int{0, 1, -1, 2, 2, -3}},
	}

	for _, it := range tests {
		t.Run(it.name, func(t *testing.T) {

			copied := make([]int, 0)
			if it.to >= 0 {
				copied = tslice.CopyWithin(it.targets, it.start, it.from, it.to)
			} else if it.from >= 0 {
				copied = tslice.CopyWithin(it.targets, it.start, it.from)
			} else {
				copied = tslice.CopyWithin(it.targets, it.start)
			}

			for i := range copied {
				if copied[i] != it.expected[i] {
					t.Errorf("unmatched [%d]: actual(%d) - expected(%d)", i, copied[i], it.expected[i])
				}
			}
		})
	}
}

func TestEvery(t *testing.T) {
	tests := []struct {
		name     string
		targets  []int
		expected bool
	}{
		{"Test true case", []int{1, 2, 3, 4, 5}, true},
		{"Test 0 length array input return true", []int{}, true},
		{"Test false case", []int{1, 2, 3, 4, -1}, false},
	}

	for _, it := range tests {
		t.Run(it.name, func(t *testing.T) {
			result := tslice.Every(it.targets, func(data int) bool {
				return data >= 0
			})

			if result != it.expected {
				t.Errorf("unmatched: actual(%v) - expect(%v)", result, it.expected)
			}
		})
	}
}

func TestEntries(t *testing.T) {
	dataArray := []string{"hello", "world", "!!"}
	expectedKeys := []int{0, 1, 2}

	iterators := tslice.Entries(dataArray)

	cnt := 0
	for i, v := range iterators {
		if i != expectedKeys[cnt] {
			t.Errorf("unmatched key: actual(%d) - expected(%d)", i, expectedKeys[cnt])
		}

		if v != dataArray[cnt] {
			t.Errorf("Unmatched value: actual(%s) - expected(%s)", v, dataArray[cnt])
		}
		cnt++
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

func TestForeach(t *testing.T) {
	cnt := 0

	targets := []int{1, 2, 3}
	tslice.Foreach(targets, func(data int) {
		cnt = data
	})

	if cnt != targets[len(targets)-1] {
		t.Errorf("unmatched called num: actual(%d) - expected(%d)", cnt, targets[len(targets)-1])
	}
}

func TestIncludes(t *testing.T) {
	tests := []struct {
		name     string
		targets  []int
		target   int
		expected bool
	}{
		{"Test true case", []int{1, 2, 3}, 1, true},
		{"Test false case", []int{1, 2, 3}, 0, false},
	}

	for _, it := range tests {
		t.Run(it.name, func(t *testing.T) {
			result := tslice.Includes(it.targets, it.target)
			if result != it.expected {
				t.Errorf("unmatched: actual(%v) - expected(%v)", result, it.expected)
			}
		})
	}
}

func TestIndexOf(t *testing.T) {
	tests := []struct {
		name     string
		targets  []int
		offset   int
		data     int
		expected int
	}{
		{"Test find case from 0", []int{1, 2, 3}, -1, 2, 1},
		{"Test find case form offset", []int{1, 2, 3, 2, 5}, 2, 2, 3},
		{"Test not found case", []int{1, 2, 3}, -1, 4, -1},
	}

	for _, it := range tests {
		t.Run(it.name, func(t *testing.T) {
			result := func() int {
				switch {
				case it.offset < 0:
					return tslice.IndexOf(it.targets, it.data)
				default:
					return tslice.IndexOf(it.targets, it.data, it.offset)
				}
			}()

			if result != it.expected {
				t.Errorf("unmatched actual(%d) - expected(%d)", result, it.expected)
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

func TestPop(t *testing.T) {
	tests := []struct {
		name     string
		targets  []int
		expected int
	}{
		{"Test should update array", []int{1, 2, 3}, 3},
	}

	for _, it := range tests {
		t.Run(it.name, func(t *testing.T) {
			n := len(it.targets)

			popped := tslice.Pop(&it.targets)
			if popped != it.expected {
				t.Errorf("unmatched value: actual(%d) - expected(%d)", popped, it.expected)
			}

			if len(it.targets) != (n - 1) {
				t.Error("unmatched size")
			}
		})
	}
}

func TestPush(t *testing.T) {
	tests := []struct {
		name     string
		targets  []int
		add      []int
		expected []int
	}{
		{"Test should update array", []int{1, 2, 3}, []int{4, 5}, []int{1, 2, 3, 4, 5}},
		{"Test 0 length added", []int{1, 2, 3}, []int{}, []int{1, 2, 3}},
	}

	for _, it := range tests {
		t.Run(it.name, func(t *testing.T) {
			n := tslice.Push(&it.targets, it.add...)
			if n != len(it.expected) {
				t.Errorf("unmatched size: actual(%d) - expected(%d)", n, len(it.expected))
			}

			for i := range it.expected {
				if it.targets[i] != it.expected[i] {
					t.Error("unmatched value")
				}
			}
		})
	}
}
