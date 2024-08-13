package main

import "fmt"

func main() {
	var arith_dianykt float32
	fmt.Print("Αριθμός διανυκτερεύσεων: ")
	fmt.Scanln(&arith_dianykt)

	var timi_dianykt float32
	fmt.Print("Τιμή ανά διανυκτέρευση: ")
	fmt.Scanln(&timi_dianykt)

	syn_timi := float32(arith_dianykt) * timi_dianykt

	var pos_ekp float32

	if arith_dianykt > 3 && timi_dianykt >= 100 {
		pos_ekp = 0.3
	} else {
		pos_ekp = 0.1
	}

	ekptosi := syn_timi * pos_ekp
	tel_timi := syn_timi - ekptosi

	fmt.Println("Τελική τιμή: ", tel_timi, " ευρώ.")
}
