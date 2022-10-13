package services

import (
	"math/rand"
	"testing"
	"time"
)

type testPasswordComplexityTest struct {
	password, expected string
}

// creating array with english letters
func getEngLetters() [26]string {
	var letters [26]string
	for i := 0; i < 26; i++ {
		letters[i] = string(rune('a' + i))
	}
	return letters
}

// creating random password with given length
func createPassword(length int) string {
	password := ""
	letters := getEngLetters()
	for i := 0; i < length; i++ {
		password += letters[rand.Intn(26)]
	}
	return password
}

func TestGetPassword(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	passwordLength := rand.Intn(100)
	password := createPassword(passwordLength)
	for i := 0; i < 10; i++ {
		if len([]rune(password)) != passwordLength {
			t.Error("Length given to createPassword function and returned string length doesnt match")
		}
	}
}

var testPasswordComplexityCheckData = []testPasswordComplexityTest{
	{createPassword(0), "password is too short"},
	{createPassword(1), "password is too short"},
	{createPassword(8), ""},
	{createPassword(40), ""},
	{createPassword(64), ""},
	{createPassword(66), "password is too long"},
	{createPassword(100), "password is too long"},
}

func TestPasswordComplexityCheck(t *testing.T) {
	for _, test := range testPasswordComplexityCheckData {
		output := PasswordComplexityCheck(test.password)
		var errorMessage string
		if output == nil {
			errorMessage = ""
		} else {
			errorMessage = output.Error()
		}
		if errorMessage != test.expected {
			t.Errorf("Password complexity check doesn't work right\nGiven: '%s', expected: '%s', got: '%s'", test.password, test.expected, errorMessage)
		}
	}
}
