package services

import (
	"math/rand"
	"time"
)

// creating array with english letters
func getEngLetters() [26]string {
	var letters [26]string
	for i := 0; i < 26; i++ {
		letters[i] = string(rune('a' + i))
	}
	return letters
}

// CreateTestString creates random password with given length
func CreateTestString(length int) string {
	rand.Seed(time.Now().UnixNano())
	password := ""
	letters := getEngLetters()
	for i := 0; i < length; i++ {
		password += letters[rand.Intn(26)]
	}
	return password
}
