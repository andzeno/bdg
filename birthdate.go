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
func (b *BirthDate) Generate(values ...int) {
	now := time.Now()
	// Generate between 18 and 88 for years
	randYears := rand.Intn(71) + 18
	// 0-11 for month
	randMonths := rand.Intn(5) + rand.Intn(6)
	// 0-30 for days
	randDays := rand.Intn(15) + rand.Intn(15)

	// Generate birthdate data based on parameter count
	paramsCnt := len(values)
	if paramsCnt == 1 {
		b.Years = values[0]
		b.Months = randMonths
		b.Days = randDays
	} else if paramsCnt == 2 {
		b.Years = values[0]
		b.Months = values[1]
		b.Days = randDays
	} else if paramsCnt >= 3 {
		b.Years = values[0]
		b.Months = values[1]
		b.Days = values[2]
	} else if paramsCnt == 0 {
		b.Years = randYears
		b.Months = randMonths
		b.Days = randDays
	}

	// Compute and generate the date for the random age
	b.Date = now.AddDate(-b.Years, b.Months, b.Days).Format("2006-01-02")

}
