package bdg

import (
	"fmt"
	"testing"
	"time"
)

func TestGenerate(t *testing.T) {
	// Call function to test
	b := BirthDate{}
	b.Generate()

	// Check if years is between 18 and 88
	if b.Years < 18 || b.Years > 88 {
		t.Errorf("Expected generated years to be between 18-88, got %d", b.Years)
	}

	// Check if months is beyond 11
	if b.Months > 11 {
		t.Errorf("Expected generated months doesn't exceed 11, got %d", b.Months)
	}

	// Check if days is beyond 30
	if b.Days > 30 {
		t.Errorf("Expected generated days doesn't exceed 30, got %d", b.Days)
	}

	fmt.Println(b.Date)

	// Check if date generated is valid
	_, err := time.Parse("2006-01-02", b.Date)

	b.Generate(32)
	// Check if years is between 18 and 88
	if b.Years != 32 {
		t.Errorf("Expected generated years to be 32, got %d", b.Years)
	}

	// Check if months is beyond 11
	if b.Months > 11 {
		t.Errorf("Expected generated months doesn't exceed 11, got %d", b.Months)
	}

	// Check if days is beyond 30
	if b.Days > 30 {
		t.Errorf("Expected generated days doesn't exceed 30, got %d", b.Days)
	}

	// Check if date generated is valid
	_, err = time.Parse("2006-01-02", b.Date)

	if err != nil {
		t.Errorf("Expected generated birth date %s isn't a valid date: %v", b.Date, err)
	}

	fmt.Println(b.Date)
}
