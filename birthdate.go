package bdg

import (
	"math/rand"
	"time"
)

type BirthDate struct {
	Years  int
	Months int
	Days   int
	Date   string
}

// Generate outputs a BirthDate type
func (b *BirthDate) Generate() {
	now := time.Now()
	// Generate between 18 and 88 for years
	randYears := rand.Intn(71) + 18
	// 0-11 for month
	randMonths := rand.Intn(5) + rand.Intn(6)
	// 0-30 for days
	randDays := rand.Intn(15) + rand.Intn(15)

	// Generate date data
	if b.Years == 0 {
		b.Years = randYears
	}

	if b.Months == 0 {
		b.Months = randMonths
	}

	if b.Days == 0 {
		b.Days = randDays
	}

	// Compute the birth date
	b.Date = now.AddDate(-b.Years, b.Months, b.Days).Format("2006-01-02")
}
