package main

import "fmt"

func main() {
	var(
		n int
		prostimo float32
	)

	fmt.Print("Αριθμός ξώβεργων; ")
	fmt.Scanln(&n)

	if n <= 72 {
		prostimo = 200
	} else {
		prostimo = 200 + float32(n - 72) * 10
	}

	fmt.Println("Το πρόστιμο για", n, "ξώβεργες είναι", prostimo, "ευρώ.")
}
