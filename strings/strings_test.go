package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
			assert.ErrorIs(t, err, tt.err)
			assert.NoError(t, err)
			assert.Equal(t, tt.length, uint(len(got)))
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
			assert.Equal(t, tt.length, uint(len(got)))
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
			assert.Equal(t, tt.length, uint(len(got)))
		})
	}
}
