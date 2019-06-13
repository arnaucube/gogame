package utils

import "math/rand"

func RandInRange(min, max int) int64 {
	return int64(rand.Intn(max-min) + min)
}
