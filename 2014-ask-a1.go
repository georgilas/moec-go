package main

import "fmt"

func main() {
	var n int
	fmt.Print("Αριθμός μαθητή: ")
	fmt.Scanln(&n)

	var g string
	switch {
		case n >= 1 && n <= 8:
			g = "ΑΜΜΟΧΩΣΤΟΣ"
		case n >= 9 && n <= 16:
			g = "ΚΕΡΥΝΕΙΑ"
		case n >= 17 && n <= 25:
			g = "ΜΟΡΦΟΥ"
	}

	fmt.Println("Ομάδα: ", g)
}
