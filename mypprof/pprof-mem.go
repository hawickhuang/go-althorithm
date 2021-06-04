package mypprof

import (
	"math/rand"
	"strings"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func Concat(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += randString(n)
	}
	return s
}

func Concat2(n int) string {
	var s strings.Builder
	for i := 0; i < n; i++ {
		s.WriteString(randString(n))
	}
	return s.String()
}