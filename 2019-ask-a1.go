package main

import "fmt"

func main() {
	var epitr_orio_tax, tax_odigou int

	fmt.Print("Επιτρεπόμενο όριο ταχύτητας; ")
	fmt.Scanln(&epitr_orio_tax)

	fmt.Print("Ταχύτητα οδηγού; ")
	fmt.Scanln(&tax_odigou)

	if tax_odigou > epitr_orio_tax {
		ypervasi := tax_odigou - epitr_orio_tax
		prostimo := float32(ypervasi) * 5
		fmt.Println("Το πρόστιμο για υπέρβαση τού επιτρεπόμενου ορίου ταχύτητας κατά", ypervasi, "χλμ ανά ώρα είναι", prostimo, "ευρώ.")
	} else {
		fmt.Println("Δεν υπάρχει πρόστιμο.")
	}
}
