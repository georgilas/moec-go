package main

import "fmt"

func main() {
	var days int
	var tariff, discount, amount float32

	tariff = 6

	fmt.Print("Days? ")
	fmt.Scanln(&days)

	if days > 7 {
		discount = 0.2
	} else {
		discount = 0.0
	}

	amount = (float32(days) * tariff) * (1.0 - discount)

	fmt.Println("Amount payable for", days, "days:", amount, "euro.")
}
