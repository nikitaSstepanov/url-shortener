package generator

import "math/rand"

var (
	symbols = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	lenOfSymbols = 62
)

func GetRandomString(length int) string {
	res := make([]rune, length)

	for i := range res {
		res[i] = symbols[rand.Intn(lenOfSymbols)]
	}

	return string(res)
}