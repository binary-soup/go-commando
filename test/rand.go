package test

import (
	"math/rand"
	"time"
)

// Test helper for generating random test values.
type Rand struct {
	*rand.Rand
}

// Create a new rand object from a custom source.
func NewRand(src rand.Source) Rand {
	return Rand{
		Rand: rand.New(src),
	}
}

// Create a new rand object using the time as the seed.
func NewRandFromTime() Rand {
	return NewRand(rand.NewSource(time.Now().UnixNano()))
}

// Generate a random ASCII string of printable characters.
func (r Rand) ASCII(len int) string {
	bytes := make([]byte, len)
	for i := range bytes {
		// printable ASCII range
		bytes[i] = byte(r.IntRange(32, 126))
	}
	return string(bytes)
}

// Generate a random int between [min, max].
func (r Rand) IntRange(min, max int) int {
	return r.Intn(max-min+1) + min
}
