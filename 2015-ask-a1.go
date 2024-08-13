package main

import "fmt"

func main() {
	var wallet float32 = 500.0
	var price float32

	fmt.Print("Τιμή υπολογιστή: ")
	fmt.Scanln(&price)

	if price > wallet {
		fmt.Println("ΑΓΟΡΑ ΑΔΥΝΑΤΗ")
	} else {
		fmt.Println("ΑΓΟΡΑ ΔΥΝΑΤΗ - Υπόλοιπο: ", wallet - price, " ευρώ.")
	}
}
