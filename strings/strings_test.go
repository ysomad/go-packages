package strings

import (
	"errors"
	"testing"
)

func TestNewUnique(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		length uint
		err    error
	}{
		{
			name:   "success",
			length: 10,
			err:    nil,
		},
		{
			name:   "0 length",
			length: 0,
			err:    errZeroLength,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUnique(tt.length)

			if !errors.Is(err, tt.err) {
				t.Errorf("want err %v, got err %v", tt.err, err)
			}
			if uint(len(got)) != tt.length {
				t.Errorf("want length %b, got length %b", tt.length, len(got))
			}
		})
	}
}

func TestNewSafe(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		length uint
	}{
		{
			name:   "success",
			length: 35,
		},
		{
			name:   "0 length",
			length: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSafe(tt.length)

			if uint(len(got)) != tt.length {
				t.Errorf("want length %b, got length %b", tt.length, len(got))
			}
		})
	}
}

func TestNewRandom(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		length uint
	}{
		{
			name:   "success",
			length: 33,
		},
		{
			name:   "0 length",
			length: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRandom(tt.length)

			if uint(len(got)) != tt.length {
				t.Errorf("want length %b, got length %b", tt.length, len(got))
			}
		})
	}
}
