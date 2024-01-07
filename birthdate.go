package bdg

import (
	"math/rand"
	"time"
)

type Nendai struct {
	Years  int
	Months int
	Days   int
	Date   string
}

// GenerateDate gives a Nendai Date member based on duration
func (n *Nendai) GenerateDate() {
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

	// Compute the date based on duration data
	n.Date = now.AddDate(-n.Years, n.Months, n.Days).Format("2006-01-02")
}
