package main

import (
	"fmt"
	"math"
)

func main() {
	var workers int = 5
	var threshold int = 500

	var name string
	var cases int

	var tot int = 0
	var diff int

	for i := 1; i <= workers; i++ {
		fmt.Print("Όνομα εργάτη: ")
		fmt.Scanln(&name)

		fmt.Print("Αριθμός κιβωτίων: ")
		fmt.Scanln(&cases)

		diff = cases - threshold

		tot += cases

		switch {
			case diff < 0:
				fmt.Println("Όνομα: ", name, " - Ρήτρα: ", math.Abs(float64(diff)), " ευρώ.")
			case diff == 0:
				fmt.Println("Όνομα: {name} - Μηδενικό επίδομα.")
			case diff >= 1 && diff <= 15:
				fmt.Println("Όνομα: ", name, " - Επίδομα: ", float64(diff) * 1.0, " ευρώ.")
			case diff >= 16 && diff <= 30:
				fmt.Println("Όνομα: :", name, " - Επίδομα: ", float64(diff) * 2.0, " ευρώ.")
			default:
				fmt.Println("Όνομα: ", name, " - Επίδομα: ", float64(diff) * 3.0, " ευρώ.")
		}
	}

	fmt.Println("Συνολικός αριθμός κιβωτίων: ", tot)
}
