package services

import (
	"math/rand"
	"testing"
)

// test for CreateTestString function
func TestCreateTestString(t *testing.T) {
	passwordLength := rand.Intn(100)
	password := CreateTestString(passwordLength)
	if len([]rune(password)) != passwordLength {
		t.Error("Length given to CreateTestString function and returned string length doesnt match")
	}
}
