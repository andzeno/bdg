package bdg

import (
	"math/rand"
	"time"
)

// randBirthAge generates a random age from 18 to 88 and the birth date from time.Now()
func randBirthAge() (int, string) {
	now := time.Now()

	// Generate a random number between 18 and 88
	age := rand.Intn(71) + 18

	// Compute and generate the date for the random age
	birthDate := now.AddDate(-age, 0, 0).Format("2006-01-02")

	return age, birthDate
}
