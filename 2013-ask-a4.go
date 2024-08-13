package main

import "fmt"

func main() {
    plithos := 0
    var synoliki_timi float32 = 0.0

    var timi float32
    fmt.Print("Δώσε τιμή προϊόντος: ")
    fmt.Scanln(&timi)

    for timi > 0.0 {
        fpa := float32(timi * 0.18)
        synoliki_timi = timi + fpa
        plithos += 1
        fmt.Print("Δώσε τιμή προϊόντος: ")
        fmt.Scanln(&timi)
    }

    var ekptosi float32
    if synoliki_timi > 100.0 {
        ekptosi = synoliki_timi * 0.10
    } else {
        ekptosi = synoliki_timi * 0.05
    }

    teliki_timi := synoliki_timi - ekptosi

    fmt.Println("Αριθμός προϊόντων: ", plithos)
    fmt.Println("Συνολική τελική τιμή: ", teliki_timi)
}
