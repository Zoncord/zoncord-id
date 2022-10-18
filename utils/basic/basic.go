package basic

import (
	"math/rand"
	"time"
)

func getEngLetters() [26]string {
	// creating array with english letters
	var letters [26]string
	for i := 0; i < 26; i++ {
		letters[i] = string(rune('a' + i))
	}
	return letters
}

func CreateTestString(length int) string {
	// CreateTestString creates random password with given length
	rand.Seed(time.Now().UnixNano())
	password := ""
	letters := getEngLetters()
	for i := 0; i < length; i++ {
		password += letters[rand.Intn(26)]
	}
	return password
}
