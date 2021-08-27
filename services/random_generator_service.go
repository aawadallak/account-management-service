package services

import (
	"math/rand"
	"time"
)

type RandomGenerator struct {
}

func (r *RandomGenerator) GenerateRandomString(length int) string {

	var seedRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seedRand.Intn(len(charset))]
	}

	randomCode := string(b)

	return randomCode
}
