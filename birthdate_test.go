package bdg

import (
	"fmt"
	"testing"
	"time"
)

func TestGenerate(t *testing.T) {
	// Call function to test
	b := Nendai{}
	b.GenerateDate()

	///
	// Test generation w/o input values
	///
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

	///
	// Test generation w/ year value present
	///
	y := Nendai{
		Years: 32,
	}
	y.GenerateDate()

	// Check if years equal input
	if y.Years != 32 {
		t.Errorf("Expected generated years to be 32, got %d", y.Years)
	}

	// Check if months is beyond 11
	if y.Months > 11 {
		t.Errorf("Expected generated months doesn't exceed 11, got %d", y.Months)
	}

	// Check if days is beyond 30
	if y.Days > 30 {
		t.Errorf("Expected generated days doesn't exceed 30, got %d", b.Days)
	}

	// Check if date generated is valid
	_, err = time.Parse("2006-01-02", y.Date)

	if err != nil {
		t.Errorf("Expected generated birth date %s isn't a valid date: %v", y.Date, err)
	}

	///
	// Test generation w/ year and month values present
	///
	ym := Nendai{
		Years:  73,
		Months: 10,
	}
	ym.GenerateDate()

	// Check if years equal input
	if ym.Years != 73 {
		t.Errorf("Expected generated years to be 73, got %d", ym.Years)
	}

	// Check if months are equal
	if ym.Months != 10 {
		t.Errorf("Expected generated months to be 10 , got %d", ym.Months)
	}
	// Check if days is beyond 30
	if ym.Days > 30 {
		t.Errorf("Expected generated days doesn't exceed 30, got %d", ym.Days)
	}

	// Check if date generated is valid
	_, err = time.Parse("2006-01-02", ym.Date)

	if err != nil {
		t.Errorf("Expected generated birth date %s isn't a valid date: %v", ym.Date, err)
	}

	///
	// Test generation w/ year, month and days values present
	///
	ymd := Nendai{
		Years:  28,
		Months: 42,
		Days:   13,
	}
	ymd.GenerateDate()

	// Check if years equal input
	if ymd.Years != 28 {
		t.Errorf("Expected generated years to be 28, got %d", ymd.Years)
	}

	// Check if months are equal
	if ymd.Months != 42 {
		t.Errorf("Expected generated months to be 42 , got %d", ymd.Months)
	}

	// Check if days are equal
	if ymd.Days != 13 {
		t.Errorf("Expected generated days to be 13, got %d", ymd.Days)
	}

	// Check if date generated is valid
	_, err = time.Parse("2006-01-02", ymd.Date)

	if err != nil {
		t.Errorf("Expected generated birth date %s isn't a valid date: %v", ymd.Date, err)
	}
}
