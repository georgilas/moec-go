package main

import (
	"fmt"
	"math"
)

func main() {
	var cost, subsidy float64

	fmt.Print("Enter cost: " )
	fmt.Scanln(&cost)

	subsidy = math.Min(1500, cost * 0.3)

	fmt.Println("Subsidy:", subsidy)
}
