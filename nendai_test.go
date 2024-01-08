package bdg

import (
	"fmt"
	"io"
	"log"
	"os"
	"testing"
	"time"
)

func TestGenerate(t *testing.T) {
	// Call function to test
	b := Nendai{}
	b.Generate()

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

	// Check if date generated is valid
	_, err := time.Parse("2006-01-02", b.Date)
	if err != nil {
		t.Errorf("Expected generated date %s isn't a valid date: %v", b.Date, err)
	}

	///
	// Test generation w/ year value present
	///
	y := Nendai{
		Years: 32,
	}
	y.Generate()

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
		t.Errorf("Expected generated date %s isn't a valid date: %v", y.Date, err)
	}

	///
	// Test generation w/ year and month values present
	///
	ym := Nendai{
		Years:  73,
		Months: 10,
	}
	ym.Generate()

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
		t.Errorf("Expected generated date %s isn't a valid date: %v", ym.Date, err)
	}

	///
	// Test generation w/ year, month and days values present
	///
	ymd := Nendai{
		Years:  28,
		Months: 42,
		Days:   13,
	}
	ymd.Generate()

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
		t.Errorf("Expected generated date %s isn't a valid date: %v", ymd.Date, err)
	}
}

// Test Stringer method
func TestString(t *testing.T) {

	// Create a Nendai
	n := Nendai{
		Years:  22,
		Months: 7,
		Days:   3,
	}

	// Redirect output for testing
	// Use a file as temporary standard output
	tmpOut := "/tmp/nendai.out"

	f, err := os.Create(tmpOut)
	if err != nil {
		log.Fatal(err)
	}

	// Close the file
	defer f.Close()

	// Preserve current standard output by assinging onto a variable
	oldStdout := os.Stdout

	// Redirect current standard output to the file
	os.Stdout = f

	// Ensure restoration of standard output location
	defer func() { os.Stdout = oldStdout }()

	fmt.Println(n)

	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	got := string(data)
	expected := "Years: 22, Months: 7, Days: 3, Date: 0000-00-00"
	if got != expected {
		t.Errorf("Expected: %s, Got %s", expected, got)
	}

	// Remove the temporary output file
	err = os.Remove(tmpOut)
	if err != nil {
		log.Fatal(err)
	}
}
