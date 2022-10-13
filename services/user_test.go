package services

import (
	"github.com/Zoncord/zoncord-id/errors"
	"math/rand"
	"os"
	"testing"
	"time"
)

type testOnePassword struct {
	password, expected string
}

type testTwoPasswords struct {
	password1, password2, expected string
}

func TestMain(m *testing.M) {
	rand.Seed(time.Now().UnixNano())
	os.Exit(m.Run())
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
func createTestPassword(length int) string {
	password := ""
	letters := getEngLetters()
	for i := 0; i < length; i++ {
		password += letters[rand.Intn(26)]
	}
	return password
}

// test for createTestPassword function
func TestCreateTestPassword(t *testing.T) {
	passwordLength := rand.Intn(100)
	password := createTestPassword(passwordLength)
	if len([]rune(password)) != passwordLength {
		t.Error("Length given to createTestPassword function and returned string length doesnt match")
	}
}

// checking tests for passwordComplexity function
func TestPasswordLengthValidation(t *testing.T) {
	var tests = []testOnePassword{
		{createTestPassword(0), errors.PasswordTooShort.Error()},
		{createTestPassword(9), errors.PasswordTooShort.Error()},
		{createTestPassword(10), ""},
		{createTestPassword(64), ""},
		{createTestPassword(65), errors.PasswordTooLong.Error()},
		{createTestPassword(100), errors.PasswordTooLong.Error()},
	}
	for _, test := range tests {
		output := PasswordLengthValidation(test.password)
		if output == nil {
			if test.expected != "" {
				t.Errorf("Got: nil error\nExpected: %s", test.expected)
			}
			continue
		}
		if output.Error() != test.expected {
			t.Errorf("Got: %s\nExpected: %s", output.Error(), test.expected)
		}
	}
}

func TestPasswordEquivalencyValidation(t *testing.T) {
	var tests = []testTwoPasswords{
		{"asdf", "asdf", ""},
		{"a", "asdf", errors.PasswordsDontMatch.Error()},
	}
	for _, test := range tests {
		output := PasswordEquivalencyValidation(test.password1, test.password2)
		if output == nil {
			if test.expected != "" {
				t.Errorf("Got: nil error\nExpected: %s", test.expected)
			}
			continue
		}
		if output.Error() != test.expected {
			t.Errorf("Got: %s\nExpected: %s", output.Error(), test.expected)
		}
	}
}

// testing passwordValidation function
func TestPasswordValidation(t *testing.T) {
	longPassword := createTestPassword(65)
	var tests = []testTwoPasswords{
		{"a", "asdf", errors.PasswordsDontMatch.Error()},
		{"asdf", "asdf", errors.PasswordTooShort.Error()},
		{longPassword, longPassword, errors.PasswordTooLong.Error()},
		{"asdfasdfasdf", "asdfasdfasdf", errors.PasswordMustIncludeNumber.Error()},
		{"asdfasdfasdf1", "asdfasdfasdf1", ""},
	}

	for _, test := range tests {
		output := PasswordsValidation(test.password1, test.password2)
		if output == nil {
			if test.expected != "" {
				t.Errorf("\nGot: nil error\nExpected: %s", test.expected)
			}
			continue
		}
		if output.Error() != test.expected {
			t.Errorf("\nGot: %s\nExpected: %s", output.Error(), test.expected)
		}
	}
}
