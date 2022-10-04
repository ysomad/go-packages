// Package strings provides functions for generating random strings using cryptographically
// secure random number generator from `crypto/rand`.
package strings

import (
	"reflect"
	"testing"
)

func TestNewBase64(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBase64(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBase64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewBase64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewWithCharset(t *testing.T) {
	type args struct {
		n  int
		cs charset
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewWithCharset(tt.args.n, tt.args.cs)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewWithCharset() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewWithCharset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRandom(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRandom(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRandom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewRandom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewWithAlphabetLower(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewWithAlphabetLower(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewWithAlphabetLower() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewWithAlphabetLower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewWithAlphabetUpper(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewWithAlphabetUpper(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewWithAlphabetUpper() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewWithAlphabetUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewWithAlphabet(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewWithAlphabet(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewWithAlphabet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewWithAlphabet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewWithSpecials(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewWithSpecials(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewWithSpecials() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewWithSpecials() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewWithAlphabetDigits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewWithAlphabetDigits(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewWithAlphabetDigits() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewWithAlphabetDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewWithDigits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewWithDigits(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewWithDigits() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewWithDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewWithCharsets(t *testing.T) {
	type args struct {
		n    int
		sets []charset
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewWithCharsets(tt.args.n, tt.args.sets...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewWithCharsets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewWithCharsets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateRandomBytes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateRandomBytes(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateRandomBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateRandomBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
