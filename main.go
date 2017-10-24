package main

import (
	"fmt"

	"github.com/alexsasharegan/codewars-golang/kata"
)

func main() {
	fmt.Printf("Actual  : %v\nExpected: %v\n", kata.NbMonths(2000, 8000, 1000, 1.5), [2]int{6, 766})
	fmt.Printf("Actual  : %v\nExpected: %v\n", kata.NbMonths(12000, 8000, 1000, 1.5), [2]int{0, 4000})
	fmt.Printf("Actual  : %v\nExpected: %v\n", kata.NbMonths(8000, 12000, 500, 1.0), [2]int{8, 597})
	fmt.Printf("Actual  : %v\nExpected: %v\n", kata.NbMonths(18000, 32000, 1500, 1.25), [2]int{8, 332})
	fmt.Printf("Actual  : %v\nExpected: %v\n", kata.NbMonths(7500, 32000, 300, 1.55), [2]int{25, 122})
}
