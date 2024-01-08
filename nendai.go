package bdg

import (
	"fmt"
	"math/rand"
	"time"
)

// Use a constant date layout throughout
const layout = "2006-01-02"

/*
A Nendai consists of:

	Years:  elapsed years
	Months: elapsed months
	Days:   elapsed days
	Date:   total elapsed combined duration in yyyy-mm-dd format
*/
type Nendai struct {
	Years  int
	Months int
	Days   int
	Date   string
}

// Generate populates all Nendai field members
func (n *Nendai) Generate() {
	now := time.Now()
	// Generate between 18 and 88 for years
	randYears := rand.Intn(71) + 18
	// 0-11 for month
	randMonths := rand.Intn(5) + rand.Intn(6)
	// 0-30 for days
	randDays := rand.Intn(15) + rand.Intn(15)

	// Generate duration data
	if n.Years == 0 {
		n.Years = randYears
	}

	if n.Months == 0 {
		n.Months = randMonths
	}

	if n.Days == 0 {
		n.Days = randDays
	}

	// Compute the Nendai Date based on duration data
	n.Date = now.AddDate(-n.Years, n.Months, n.Days).Format(layout)
}

// Define a Stringer method to aid in display
func (n *Nendai) String() string {
	return fmt.Sprintf("Years: %d, Months: %d, Days: %d, Date: %s", n.Years, n.Months, n.Days, n.Date)
}

// GenerateDate Nendai Date member based on existing duration
// func (n *Nendai) GenerateData() {
// 	nTime, err := time.Parse(layout, n.Date)

// }
