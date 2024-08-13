package main

import "fmt"

func main() {
	var kat, poso float32

	fmt.Print("Κατανάλωση; ")
	fmt.Scanln(&kat)

	if kat <= 80 {
		poso = kat * 0.7
	} else {
		poso = kat * 0.95 + 10
	}

	fmt.Println("Ποσό πληρωμής για κατανάλωση", kat, "κ.μ. είναι", poso, "ευρώ.")
}
