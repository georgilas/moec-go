package main

import "fmt"

func main() {
    var grapta, proforika int
    fmt.Print("Γραπτά: ")
    fmt.Scanln(&grapta)
    fmt.Print("Προφορικά: ")
    fmt.Scanln(&proforika)

    athrisma := grapta + proforika

    if grapta >= 10 && proforika >=10 && athrisma > 22 {
        fmt.Println("Συνολική βαθμολογία:", athrisma)
    } else {
        fmt.Println("ΕΠΑΝΕΞΕΤΑΣΗ")
    }
}
