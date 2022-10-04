// Package strings provides functions for generating random strings using cryptographically
// secure random number generator from `crypto/rand`.
package strings

import (
	"crypto/rand"
	"encoding/base64"
	"math/big"
)

type charset string

const (
	// Alphabet contains all lowercase characaters from english alphabet.
	AlphabetLower = charset("abcdefghijklmnopqrstuvwxyz")

	// Alphabet contains all uppercase characaters from english alphabet.
	AlphabetUpper = charset("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	// Alphabet contains lowercase and uppercase characters from english alphabet.
	Alphabet = charset(AlphabetLower + AlphabetUpper)

	// Digits contains all digits.
	Digits = charset("0123456789")

	// AlphabetDigits contains uppercase, lowercase characters from english alphabet and digits.
	AlphabetDigits = charset(Alphabet + Digits)

	// Specials contains all special characters.
	Specials = charset("~=+-_%^&*/()[]{}<>.,:;/!@#$?|\"'")

	// AllChars contains uppercase, lowercase characters from english alphabet, digits and special characters.
	AllChars = charset(AlphabetDigits + Digits + Specials)
)

// NewBase64 returns URL-safe, base64 encoded random string.
// For example It can be used for session id or csrf tokens.
func NewBase64(n int) (string, error) {
	b, err := generateRandomBytes(n)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), nil
}

// NewWithCharset returns securely generated random string from given charset.
func NewWithCharset(n int, cs charset) (string, error) {
	b := make([]byte, n)

	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(cs))))
		if err != nil {
			return "", err
		}

		b[i] = cs[num.Int64()]
	}

	return string(b), nil
}

func NewRandom(n int) (string, error)             { return NewWithCharset(n, AllChars) }
func NewWithAlphabetLower(n int) (string, error)  { return NewWithCharset(n, AlphabetLower) }
func NewWithAlphabetUpper(n int) (string, error)  { return NewWithCharset(n, AlphabetUpper) }
func NewWithAlphabet(n int) (string, error)       { return NewWithCharset(n, Alphabet) }
func NewWithSpecials(n int) (string, error)       { return NewWithCharset(n, Specials) }
func NewWithAlphabetDigits(n int) (string, error) { return NewWithCharset(n, AlphabetDigits) }
func NewWithDigits(n int) (string, error)         { return NewWithCharset(n, Digits) }

// NewWithCharsets returns random string generated from slice of charsets,
// if sets are not specified, string will be generated from alphabet
// with lower, uppercase characters and digits.
//
// Better to use `NewWithCharset` or `NewWith{{charset}}` functions for better perfomance
// or `NewRandom` if all characters are needed.
func NewWithCharsets(n int, sets ...charset) (string, error) {
	var chars charset
	if len(sets) > 0 {
		for _, set := range sets {
			chars += set
		}
	} else {
		chars = AlphabetDigits
	}

	return NewWithCharset(n, charset(chars))
}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
