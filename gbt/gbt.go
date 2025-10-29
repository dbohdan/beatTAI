package main

import (
	"fmt"

	"github.com/brandondube/tai"
)

func main() {
	g := tai.Now().AsGregorian()

	// pull the hour, minute, second, and attosecond from the TAI time
	// and convert them to milliseconds
	h := g.Hour * 3_600_000
	m := g.Min * 60_000
	s := g.Sec * 1_000
	ms := g.Asec / 1_000_000_000_000_000

	// do the math for beatTAI and store it
	beatTAI := (float64(h+m+s) + float64(ms)) / 86400

	// print the value, rounded to two decimal places, padded with leading zeroes
	// if necessary and prefixed with : for proper beatTAI syntax
	fmt.Printf(":%06.2f\n", beatTAI)
}
