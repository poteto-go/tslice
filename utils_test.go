package tslice

import "testing"

func TestPrivateToString(t *testing.T) {
	type User struct {
		name string
	}

	tests := []struct {
		name     string
		data     any
		expected string
	}{
		{"test string case", string("5"), "5"},
		{"test []byte case", []byte("5"), "5"},
		{"test []rune case", []rune("5"), "5"},
		{"test int case", int(5), "5"},
		{"test int8 case", int8(5), "5"},
		{"test int16 case", int16(5), "5"},
		{"test int32 case", int32(5), "5"},
		{"test int64 case", int64(5), "5"},
		{"test unit case", uint(5), "5"},
		{"test uint8 case", uint8(5), "5"},
		{"test unit16 case", uint16(5), "5"},
		{"test unit32 case", uint32(5), "5"},
		{"test uint64 case", uint64(5), "5"},
		{"test float32 case", float32(5.010), "5.01"},
		{"test float64 case", float64(5.010), "5.01"},
		{"test bool case", bool(false), "false"},
		{"test struct case", User{"name"}, "{name}"},
	}

	for _, it := range tests {
		t.Run(it.name, func(t *testing.T) {
			result := toString(it.data)
			if result != it.expected {
				t.Errorf(
					"unmatched: actual(%s) - expected(%s)",
					result, it.expected,
				)
			}
		})
	}
}
