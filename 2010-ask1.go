package main

import
    "fmt"

func main() {
    
    cost_per_gram := 0.017

    var weight_in_grams int
    fmt.Println("Weight in grams?")
    fmt.Scanln(&weight_in_grams)

    if weight_in_grams >=20 && weight_in_grams <= 500 {
        amount_payable := float64(weight_in_grams) * cost_per_gram
        fmt.Println("Amount payable:", amount_payable)
    } else {
        fmt.Println("WEIGHT OUT OF LIMITS")
    }
}
