package bdg

import (
	"testing"
	"time"
)

func TestRandBirthAge(t *testing.T) {
	// Call function to test
	age, birthDate := randBirthAge()

	// Check if age is between 18 and 88
	if age < 18 || age > 88 {
		t.Errorf("Expected generated age to be between 18-88, got %d", age)
	}

	// Check if date generated is valid
	_, err := time.Parse("2006-01-02", birthDate)

	if err != nil {
		t.Errorf("Expected generated birth date %s isn't a valid date: %v", birthDate, err)
	}

}
