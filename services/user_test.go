package services

import (
	"math/rand"
	"testing"
	"time"
)

type testPasswordComplexityCheck struct {
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
func createTestPassword(length int) string {
	rand.Seed(time.Now().UnixNano())
	password := ""
	letters := getEngLetters()
	for i := 0; i < length; i++ {
		password += letters[rand.Intn(26)]
	}
	return password
}

// test for createTestPassword function
func TestCreateTestPassword(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	passwordLength := rand.Intn(100)
	password := createTestPassword(passwordLength)
	for i := 0; i < 10; i++ {
		if len([]rune(password)) != passwordLength {
			t.Error("Length given to createTestPassword function and returned string length doesnt match")
		}
	}
}

// tests for passwordComplexity function
var testsPasswordComplexityCheckData = []testPasswordComplexityCheck{
	{createTestPassword(0), "password is too short"},
	{createTestPassword(1), "password is too short"},
	{createTestPassword(8), ""},
	{createTestPassword(40), ""},
	{createTestPassword(64), ""},
	{createTestPassword(66), "password is too long"},
	{createTestPassword(100), "password is too long"},
}

// checking tests for passwordComplexity function
func TestPasswordComplexityCheck(t *testing.T) {
	for _, test := range testsPasswordComplexityCheckData {
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

// test structure
type testPasswordValidation struct {
	password1, password2 string
	expectedBool         bool
	expectedErr          string
}

// array with tests for passwordValidation function
var testsPasswordValidation []testPasswordValidation

// tests with non-equal length passwords
func prepareNonEqualLengthPasswordsTestsPasswordValidation() {
	for i := 0; i < 10; i++ {
		password1 := createTestPassword(10)
		password2 := createTestPassword(8)
		var test testPasswordValidation = testPasswordValidation{
			password1,
			password2,
			false,
			"passwords do not match",
		}
		testsPasswordValidation = append(testsPasswordValidation, test)
	}
}

// tests with doesn't match passwords
func prepareDoesntMatchPasswordsTestsPasswordValidation() {
	for i := 0; i < 10; i++ {
		password1 := createTestPassword(9) + "a"
		password2 := createTestPassword(9) + "b"
		var test testPasswordValidation = testPasswordValidation{
			password1,
			password2,
			false,
			"passwords do not match",
		}
		testsPasswordValidation = append(testsPasswordValidation, test)
	}
}

// tests with equal passwords
func prepareMatchPasswordsTestsPasswordValidation() {
	for i := 0; i < 10; i++ {
		password := createTestPassword(10)
		var test testPasswordValidation = testPasswordValidation{
			password,
			password,
			true,
			"",
		}
		testsPasswordValidation = append(testsPasswordValidation, test)
	}
}

// testing passwordValidation function
func TestPasswordValidation(t *testing.T) {
	prepareNonEqualLengthPasswordsTestsPasswordValidation()
	prepareDoesntMatchPasswordsTestsPasswordValidation()
	prepareMatchPasswordsTestsPasswordValidation()

	for _, test := range testsPasswordValidation {
		isEqual, err := PasswordValidation(test.password1, test.password2)
		if !isEqual && err == nil || isEqual && err != nil {
			t.Error("Returned values doesnt match")
		}
		if (err == nil && test.expectedErr != "") || (err != nil && err.Error() != test.expectedErr) {
			t.Error("Returned error doesnt match to expected error")
		}
	}
}
