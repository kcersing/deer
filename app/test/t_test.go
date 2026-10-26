package test

import (
	"testing"
)

func TestToUint64(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected uint64
	}{
		{"int positive", int(100), 100},
		{"int negative", int(-1), 18446744073709551615}, // uint64 overflow behavior
		{"int8", int8(8), 8},
		{"int16", int16(16), 16},
		{"int32", int32(32), 32},
		{"int64", int64(64), 64},
		{"uint", uint(10), 10},
		{"uint8", uint8(8), 8},
		{"uint16", uint16(16), 16},
		{"uint32", uint32(32), 32},
		{"uint64", uint64(64), 64},
		{"uintptr", uintptr(123), 123},
		{"float32", float32(3.14), 3},
		{"float64", float64(9.99), 9},
		{"string invalid", "hello", 0},
		{"nil input", nil, 0},
		{"bool invalid", true, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			result := ToUint64(tt.input)

			if result != tt.expected {
				t.Errorf("ToUint64(%v) = %d; want %d", tt.input, result, tt.expected)
			} else {
				t.Logf("ToUint64(%v) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func ToUint64[T any](a T) uint64 {
	switch v := any(a).(type) {
	case int:
		return uint64(v)
	case int8:
		return uint64(v)
	case int16:
		return uint64(v)
	case int32:
		return uint64(v)
	case int64:
		return uint64(v)
	case uint:
		return uint64(v)
	case uint8:
		return uint64(v)
	case uint16:
		return uint64(v)
	case uint32:
		return uint64(v)
	case uint64:
		return v
	case uintptr:
		return uint64(v)
	case float32:
		return uint64(v)
	case float64:
		return uint64(v)
	default:
		return 0
	}
}
