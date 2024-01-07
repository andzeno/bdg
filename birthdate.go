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

// Generate outputs a BirthDate type
func (n *Nendai) Generate() {
	now := time.Now()
	// Generate between 18 and 88 for years
	randYears := rand.Intn(71) + 18
	// 0-11 for month
	randMonths := rand.Intn(5) + rand.Intn(6)
	// 0-30 for days
	randDays := rand.Intn(15) + rand.Intn(15)

	// Generate date data
	if n.Years == 0 {
		n.Years = randYears
	}

	if n.Months == 0 {
		n.Months = randMonths
	}

	if n.Days == 0 {
		n.Days = randDays
	}

	// Compute the birth date
	n.Date = now.AddDate(-n.Years, n.Months, n.Days).Format("2006-01-02")
}
