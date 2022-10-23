package validation

import (
	"fmt"
	"github.com/Zoncord/zoncord-id/services"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	rand.Seed(time.Now().UnixNano())
	os.Exit(m.Run())
}

type isEqTest struct {
	err1, err2 error
	expected   bool
}

func TestIsEq(t *testing.T) {
	var tests = []isEqTest{
		{fmt.Errorf("a"), fmt.Errorf("a"), true},
		{nil, nil, true},
		{fmt.Errorf("b"), fmt.Errorf("a"), false},
		{nil, fmt.Errorf("a"), false},
	}
	for _, test := range tests {
		if IsEq(test.err1, test.err2) != test.expected {
			t.Errorf("\nExpected: %v\nGot: %v", test.err1, test.err2)
		}
	}
}

type testSimpleValidation struct {
	validationValue validationValue
	expected        error
}

func TestSimpleValidation(t *testing.T) {
	errorsProvider := newTestValidationValue("")
	var validationValues = []testSimpleValidation{
		{newTestValidationValue(""), errorsProvider.ErrRequiredValue()},
		{newTestValidationValue("my name"), errorsProvider.ErrSpacesForbidden()},
		{newTestValidationValue("my1name"), errorsProvider.ErrNumbersForbidden()},
		{newTestValidationValue("my$name"), errorsProvider.ErrSpecialCharactersForbidden()},
		{newTestValidationValue(services.CreateTestString(65)), errorsProvider.ErrValueTooLong()},
	}

	for _, test := range validationValues {
		err := CheckRules(SimpleValidationRules(test.validationValue))
		if !IsEq(err, test.expected) {
			t.Errorf("\nExcpected: %v\nGot: %v", test.expected, err)
		}
	}
}

type testOnePassword struct {
	password string
	expected error
}

type testTwoPasswords struct {
	password1, password2 string
	expected             error
}

// checking tests for passwordComplexity function
func TestPasswordLengthRule(t *testing.T) {
	errorsProvider := newPasswordValidationValue("")
	var tests = []testOnePassword{
		{services.CreateTestString(0), errorsProvider.ErrRequiredValue()},
		{services.CreateTestString(9), errorsProvider.ErrValueTooShort()},
		{services.CreateTestString(10), errorsProvider.numbersRequiredRule()},
		{services.CreateTestString(63) + "1", nil},
		{services.CreateTestString(65), errorsProvider.ErrValueTooLong()},
		{services.CreateTestString(100), errorsProvider.ErrValueTooLong()},
	}
	for _, test := range tests {
		output := PasswordValidation(test.password)
		if !IsEq(output, test.expected) {
			t.Errorf("\nGot: %v\nExpected: %v", output, test.expected)
		}
	}
}

func TestPasswordsEquivalencyRule(t *testing.T) {
	var tests = []testTwoPasswords{
		{"asdf", "asdf", nil},
		{"a", "asdf", ErrPasswordsDontMatch},
	}
	for _, test := range tests {
		var validationVal = twoValidationValues{
			newTestValidationValue(test.password1),
			newTestValidationValue(test.password2),
		}
		output := validationVal.equivalencyRule()
		if !IsEq(output, test.expected) {
			t.Errorf("\nGot: %v\nExpected: %v", output, test.expected)
		}
	}
}

// testing passwordValidation function
func TestPasswordValidation(t *testing.T) {
	longPassword := services.CreateTestString(65)
	errorsProvider := newPasswordValidationValue("")
	var tests = []testTwoPasswords{
		{"a", "asdf", ErrPasswordsDontMatch},
		{"", "", errorsProvider.ErrRequiredValue()},
		{"asdf", "asdf", errorsProvider.ErrValueTooShort()},
		{longPassword, longPassword, errorsProvider.ErrValueTooLong()},
		{"asdfasdfasdf", "asdfasdfasdf", errorsProvider.ErrValueMustIncludeNumber()},
		{"asdfasdfasdf1", "asdfasdfasdf1", nil},
	}

	for _, test := range tests {
		output := PasswordsValidation(test.password1, test.password2)
		if !IsEq(output, test.expected) {
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
		{"", ErrInvalidEmailFormat},
		{"asdf", ErrInvalidEmailFormat},
		{"asdfasdf@", ErrInvalidEmailFormat},
		{"asdfasdf@asdfsafd", ErrInvalidEmailFormat},
		{"asdfasdf@asdfsafd.", ErrInvalidEmailFormat},
		{"asdfasdf@asdfsafd.a", ErrInvalidEmailFormat},
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
