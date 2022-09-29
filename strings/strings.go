package strings

import (
	cryptoRand "crypto/rand"
	"errors"
	mathRand "math/rand"
)

const (
	chars   = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	special = " !\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
)

// NewUnique generates random cryptographically secure string
func NewUnique(length uint) (string, error) {
	if length == 0 {
		return "", errors.New("length must be greater than 0")
	}

	bytes := make([]byte, length)
	if _, err := cryptoRand.Read(bytes); err != nil {
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = chars[b%byte(len(chars))]
	}

	return string(bytes), nil
}

// NewSafe generates random URL safe string
func NewSafe(length uint) string {
	bytes := make([]byte, length)

	for i := range bytes {
		bytes[i] = chars[mathRand.Intn(len(chars))]
	}

	return string(bytes)
}

// NewRandom generates random string with special characters
func NewRandom(length uint) string {
	bytes := make([]byte, length)

	c := chars + special
	for i := range bytes {
		bytes[i] = c[mathRand.Intn(len(chars))]
	}

	return string(bytes)
}
