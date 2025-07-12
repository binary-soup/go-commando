package test

import (
	"math/rand"
)

// Generate a random ASCII string of printable characters.
func RandASCII(src rand.Source, len int) string {
	r := rand.New(src)

	bytes := make([]byte, len)
	for i := range bytes {
		// printable ASCII range
		bytes[i] = byte(r.Intn(95) + 32)
	}
	return string(bytes)
}
