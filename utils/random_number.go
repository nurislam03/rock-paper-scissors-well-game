package utils

import (
	"math/rand"
	"time"
)

// GenerateRandomNumberWithLimit returns random int following [0,limit)
func GenerateRandomNumberWithLimit(limit int) int {
	rand.Seed(time.Now().UnixNano())
	randNumber := rand.Intn(limit)
	return randNumber
}
