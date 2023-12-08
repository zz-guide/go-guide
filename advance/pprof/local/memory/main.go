package main

import (
	"github.com/pkg/profile"
	"math/rand"
	"strings"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

/**
	go tool pprof -http=:8888 /var/folders/pr/b1mkwkcd5dbcr5j4s13s5p9w0000gn/T/profile101715994/mem.pprof

 */
func main() {
	defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()
	concat(100)
}

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func concat(n int) string {

	var s strings.Builder
	for i := 0; i < n; i++ {
		s.WriteString(randomString(n))
	}
	return s.String()
	/*
	s := ""
	for i := 0; i < n; i++ {
		s += randomString(n)
	}
	return s*/
}