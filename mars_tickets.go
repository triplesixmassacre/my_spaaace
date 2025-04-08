// задачка про билеты для путешествия на марс
// https://golangify.com/code-ticket

package main

import (
	"fmt"
	"math/rand"
)

func () {
	const secPerDay = 86400
	distance := 62100000
	company := ""
	trip := ""

	fmt.Println("Название                Days   Trip type   Price")
	fmt.Println("=================================================")

	for i := 0; i < 10; i++ {
		switch rand.Intn(3) {
		case 0:
			company = "Virgin Galactic"
		case 1:
			company = "SpaceX"
		case 2:
			company = "Space Adventures"
		}

		speed := rand.Intn(15) + 16
		duration := distance / speed / secPerDay
		cost := 20.0 + speed

		if rand.Intn(2) == 1 {
			trip = "Round-trip"
			cost = cost * 2
		} else {
			trip = "One-way"
		}

		fmt.Printf("%-16s %4v %-10s $%4v\n", company, duration, trip, cost)
	}
}
