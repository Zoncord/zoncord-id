package services

import (
	"github.com/Zoncord/zoncord-id/errors"
	"math/rand"
	"os"
	"testing"
	"time"
)

type testOnePassword struct {
	password string
	expected error
}

type testTwoPasswords struct {
	password1, password2 string
	expected             error
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
		{createTestPassword(0), errors.PasswordTooShort},
		{createTestPassword(9), errors.PasswordTooShort},
		{createTestPassword(10), nil},
		{createTestPassword(64), nil},
		{createTestPassword(65), errors.PasswordTooLong},
		{createTestPassword(100), errors.PasswordTooLong},
	}
	for _, test := range tests {
		output := PasswordLengthValidation(test.password)
		if output != test.expected {
			t.Errorf("\nGot: %v\nExpected: %v", output, test.expected)
		}
	}
}

func TestPasswordsEquivalencyValidation(t *testing.T) {
	var tests = []testTwoPasswords{
		{"asdf", "asdf", nil},
		{"a", "asdf", errors.PasswordsDontMatch},
	}
	for _, test := range tests {
		output := PasswordsEquivalencyValidation(test.password1, test.password2)
		if output != test.expected {
			t.Errorf("\nGot: %v\nExpected: %v", output, test.expected)
		}
	}
}

// testing passwordValidation function
func TestPasswordValidation(t *testing.T) {
	longPassword := createTestPassword(65)
	var tests = []testTwoPasswords{
		{"a", "asdf", errors.PasswordsDontMatch},
		{"asdf", "asdf", errors.PasswordTooShort},
		{longPassword, longPassword, errors.PasswordTooLong},
		{"asdfasdfasdf", "asdfasdfasdf", errors.PasswordMustIncludeNumber},
		{"asdfasdfasdf1", "asdfasdfasdf1", nil},
	}

	for _, test := range tests {
		output := PasswordsValidation(test.password1, test.password2)
		if output != test.expected {
			t.Errorf("\nGot: %v\nExpected: %v", output, test.expected)
		}
	}
}

type testEmail struct {
	email    string
	expected error
}

func TestEmailValidation(t *testing.T) {
	var tests = []testEmail{
		{"", errors.InvalidEmailFormat},
		{"asdf", errors.InvalidEmailFormat},
		{"asdfasdf@", errors.InvalidEmailFormat},
		{"asdfasdf@asdfsafd", errors.InvalidEmailFormat},
		{"asdfasdf@asdfsafd.", errors.InvalidEmailFormat},
		{"asdfasdf@asdfsafd.a", errors.InvalidEmailFormat},
		{"asdfasdf@asdfasfd.asdf", nil},
		{"asdf-asdf@asdfasfd.asdf", nil},
		{"asdf_asdf@asdfasfd.asdf", nil},
		{"asdfasdf@asdf.asfd.asdf", nil},
		{"asdf.asdf@asdf.asfd.asdf", nil},
	}
	for _, test := range tests {
		output := EmailValidation(test.email)
		if output != test.expected {
			t.Errorf("\nEmail: %s\nGot: %v\nExpected: %v", test.email, output, test.expected)
		}
	}
}
