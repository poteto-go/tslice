package tslice_test

import (
	"errors"
	"testing"

	"github.com/poteto-go/tslice"
)

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

func TestCopyWithinPanicCase(t *testing.T) {
	tests := []struct {
		name    string
		targets []int
		start   int
		from    int
		to      int
		add     int
	}{
		{"Test panic if args > 2", []int{1, 2}, 1, 1, 1, 1},
		{"Test panic if from >= to", []int{1, 2, 3}, 1, 2, 1, -1},
		{"Test panic if from <= -1", []int{1, 2}, 1, -1, 1, -1},
		{"Test panic if to <= 0", []int{1, 2}, 1, -1, 0, -1},
		{"Test panic if from >= length", []int{1, 2}, 1, 2, 2, -1},
		{"Test panic if start >= length", []int{1, 2}, 3, 0, 1, -1},
		{"Test panic if start < 0", []int{1, 2}, -1, 0, 1, -1},
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
				tslice.CopyWithin(it.targets, it.start, it.from, it.to, it.add)
			} else {
				tslice.CopyWithin(it.targets, it.start, it.from, it.to)
			}

			if !errors.Is(err, errors.New("error")) {
				t.Error("unmatched not panic")
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

func TestIndexOfPanicCase(t *testing.T) {
	tests := []struct {
		name    string
		targets []int
		offset  int
		data    int
		add     int
	}{
		{"Test over 2 args case", []int{1, 2, 3}, 0, 2, 1},
		{"Test startIndex < 0 case", []int{1, 2, 3}, -1, 2, -1},
		{"Test startIndex >= len(array) case", []int{1, 2, 3}, 5, 2, -1},
	}

	for _, it := range tests {
		t.Run(it.name, func(t *testing.T) {
			var err error
			defer func() {
				if rec := recover(); rec != nil {
					err = errors.New("error")
				}
			}()

			if it.offset < -1 {
				tslice.IndexOf(it.targets, it.data)
			} else if it.add < 0 {
				tslice.IndexOf(it.targets, it.data, it.offset)
			} else {
				tslice.IndexOf(it.targets, it.data, it.offset, it.add)
			}
			if !errors.Is(err, errors.New("error")) {
				t.Error("unmatched not panic")
			}
		})
	}
}

func TestPopPanicCase(t *testing.T) {
	var err error
	defer func() {
		if rec := recover(); rec != nil {
			err = errors.New("error")
		}
	}()

	dataArray := []int{}
	tslice.Pop(&dataArray)

	if !errors.Is(err, errors.New("error")) {
		t.Error("unmatched not panic")
	}
}

func TestReducePanicCase(t *testing.T) {
	var err error
	defer func() {
		if rec := recover(); rec != nil {
			err = errors.New("error")
		}
	}()

	dataArray := []int{}
	tslice.Reduce(dataArray, func(acc int, cur int) int {
		return acc + cur
	}, 1, 2)

	if !errors.Is(err, errors.New("error")) {
		t.Error("unmatched not panic")
	}
}

func TestReduceRightPanicCase(t *testing.T) {
	var err error
	defer func() {
		if rec := recover(); rec != nil {
			err = errors.New("error")
		}
	}()

	dataArray := []int{}
	tslice.ReduceRight(dataArray, func(acc int, cur int) int {
		return acc + cur
	}, 1, 2)

	if !errors.Is(err, errors.New("error")) {
		t.Error("unmatched not panic")
	}
}
