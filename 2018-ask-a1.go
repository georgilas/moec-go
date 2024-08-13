package main

import "fmt"

func main() {
	var leof int

	fmt.Print("Πόσα λεωφορεία θα χρειασστούν; ")
	fmt.Scanln(&leof)

	var timi_ana_leof float32 = 50

	var pososto_ekp float32

	if  leof > 15 {
		pososto_ekp = 0.2
	} else {
		pososto_ekp = 0.1
	}

	teliki_timi_pro_ekp := float32(leof) * timi_ana_leof

	ekptosi := teliki_timi_pro_ekp * pososto_ekp

	teliki_timi := teliki_timi_pro_ekp - ekptosi

	fmt.Println("Έκπτωση: ", ekptosi, " ευρώ.")
	fmt.Println("Τελικό ποσό: ", teliki_timi, " ευρώ.")
}
