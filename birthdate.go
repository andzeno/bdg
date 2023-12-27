package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	now := time.Now()
	rand.Seed(now.UnixNano())

	// Generate a random number between 18 and 88
	randAge := rand.Intn(71) + 18

	//
	nYearsAgo := now.AddDate(-randAge, 0, 0)
	fmt.Printf("Date %d Years Ago: %s\n", randAge, nYearsAgo)
}

