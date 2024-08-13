package main

import "fmt"

func main() {
	var (
		timi_isitiriou float32
		ilikia_pelati int
	)

	fmt.Print("Τιμή εισιτηρίου; ")
	fmt.Scanln(&timi_isitiriou)

	fmt.Print("Ηλικία πελάτη; ")
	fmt.Scanln(&ilikia_pelati)

	if ilikia_pelati > 18 && ilikia_pelati <= 65 {
		fmt.Println("ΚΑΜΜΙΑ ΕΚΠΤΩΣΗ - Τιμή εισιτηρίου: ", timi_isitiriou, " ευρώ.")
	} else {
		fmt.Println("Τιμή εισιτηρίου: ", timi_isitiriou * 0.6, " ευρώ.")
	}
}
