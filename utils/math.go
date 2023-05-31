package utils

import "math/rand"

func RandMax(max uint) int {
	return rand.Intn(int(max + 1))
}

func RandMinMax(min uint, max uint) int {
	return int(min) + RandMax(max-min)
}
