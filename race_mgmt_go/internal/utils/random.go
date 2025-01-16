package utils

import (
	"math/rand"
	"time"
)

// Create a package-level random source and generator for reusable random operations.
var random = rand.New(rand.NewSource(time.Now().UnixNano()))

// RandomInt generates a random integer between min and max (inclusive).
func RandomInt(min, max int) int {
	if min > max {
		min, max = max, min
	}
	return random.Intn(max-min+1) + min
}

// RandomFloat generates a random float64 between min and max.
func RandomFloat(min, max float64) float64 {
	if min > max {
		min, max = max, min
	}
	return min + random.Float64()*(max-min)
}

// RandomBool generates a random boolean value.
func RandomBool() bool {
	return random.Intn(2) == 1
}

// RandomChance returns true with the given probability (0-100%).
func RandomChance(probability int) bool {
	if probability <= 0 {
		return false
	}
	if probability >= 100 {
		return true
	}
	return random.Intn(100) < probability
}

// min returns the smaller of two integers.
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
