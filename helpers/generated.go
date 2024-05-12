package helpers

import (
	"math/rand"
	"time"
)

const charset = "0123456789"

func GenerateAccountNumber(types string) string {
	prefix := "31"
	lenNumber := 8
	if types == "CREDIT" {
		prefix = "16"
	}
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, lenNumber)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return prefix + string(b)
}
