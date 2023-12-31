package bdg

import (
	"math/rand"
	"time"
)

// genBirthAge generates the age
func genBirthAge(age int) (int, string) {
	now := time.Now()

	// Compute and generate the date for the random age
	birthDate := now.AddDate(-age, 0, 0).Format("2006-01-02")

	return age, birthDate
}

// RandBirthAge generates a random age from 18 to 88 and the birth date from time.Now()
func RandBirthAge() (int, string) {
	// Generate a random number between 18 and 88
	randAge := rand.Intn(71) + 18

	age, birthDate := genBirthAge(randAge)

	return age, birthDate
}

// SpecBirthAge generates the birthdate for the supplied age target
func SpecBirthAge(age int) string {
	_, birthDate := genBirthAge(age)

	return birthDate
}
