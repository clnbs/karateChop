package mathtool

import (
	"math/rand"
	"time"
)

//RandomIntBetween returns a random int between two values (int)
func RandomIntBetween(a, b int) int {
	if a == b {
		return a
	}
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	if a <= 0 && b <= 0 {
		a *= -1
		b *= -1
		if b < a {
			return (r1.Intn(a-b) + b) * -1
		}
		return (r1.Intn(b-a) + a) * -1
	}

	if b < a {
		return r1.Intn(a-b) + b
	}
	return r1.Intn(b-a) + a
}
