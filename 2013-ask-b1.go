package main

import "fmt"

func main() {
    plithos := 0
    plithos_mioseon := 0
    var syn_sim_misthon float32 = 0.0
    var syn_mioseon float32 = 0.0
    var miosi float32

    var name string
    fmt.Print("Δώσε όνομα υπαλλήλου (ΤΕΛΟΣ για διακοπή): ")
    fmt.Scanln(&name)

    for name != "ΤΕΛΟΣ" {
        var simerinos_misthos float32
        fmt.Print("Δώσε σημερινό μισθό υπαλλήλου: ")
        fmt.Scanln(&simerinos_misthos)

        for simerinos_misthos <=0 {
            fmt.Print("ΛΑΘΟΣ ΜΙΣΘΟΣ - Δώσε σημερινό μισθό υπαλλήλου: ")
            fmt.Scanln(&simerinos_misthos)
        }

        plithos += 1
        syn_sim_misthon += simerinos_misthos

        if simerinos_misthos <= 1000 {
            miosi = 0.0
        } else if simerinos_misthos <= 2500 {
            miosi = (simerinos_misthos - 1000) * 0.05
            plithos_mioseon += 1
        } else if simerinos_misthos <= 4000 {
            miosi = 1500 * 0.05 + (simerinos_misthos - 2500) * 0.075
            plithos_mioseon += 1
        } else {
            miosi = 1500 * 0.05 + 1500 * 0.075 + (simerinos_misthos - 4000) * 0.10
//          plithos_mioseon += 1
        }

        syn_mioseon += miosi
        var neos_misthos float32 = simerinos_misthos - miosi
        fmt.Println("Σημερινός μισθός: ", simerinos_misthos, " - Μείωση: ", miosi, " = Νέος μισθός: ", neos_misthos)

        fmt.Print("Δώσε όνομα υπαλλήλου (ΤΕΛΟΣ για διακοπή): ")
        fmt.Scanln(&name)
    }

    fmt.Println("Πλήθος υπαλλήλων: ", plithos)
    fmt.Println("Σύνολο σημερινών μισθών (χωρίς μειώσεις): ", syn_sim_misthon)
    fmt.Println("Σύνολο μειώσεων: ", syn_mioseon)
    fmt.Println("Πλήθος υπαλλήληων με μείωση μισθού: ", plithos_mioseon)
}
