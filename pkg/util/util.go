package util

import (
	"math/rand"
)

func RandStr(n int) string {
	var letters = []byte("asdfghjklqwertyuiopzxcvbnmASDFGHJKLQWERTYUIOPZXCVBNM")
	result := make([]byte, n)

	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
