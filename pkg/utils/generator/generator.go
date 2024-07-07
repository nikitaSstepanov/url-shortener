package generator

import "math/rand"

const (
	lenOfSymbols = 62
)

var (
	symbols = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
)

func GetRandomString(length int) string {
	res := make([]rune, length)

	for i := range res {
		res[i] = symbols[rand.Intn(lenOfSymbols)]
	}

	return string(res)
}