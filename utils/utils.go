package utils

import (
	"encoding/json"
	"math/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateRandomString(n int) string {
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func TransformData(data interface{}) string {
	b, _ := json.MarshalIndent(data, "", "  ")
	return string(b)
}
