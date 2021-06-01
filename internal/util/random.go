package util

import (
	"math/rand"
	"strconv"
	"time"
)

const (
	charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	numset = "1234567890"
)

var seededRand = rand.New(
	rand.NewSource(time.Now().UnixNano()),
)

func RandomStringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandomIntLength(length int) int {
	str := RandomStringWithCharset(length, numset)
	i, _ := strconv.Atoi(str)
	return i
}

func RandomAlphaNumString(length int) string {
	return RandomStringWithCharset(length, charset)
}
