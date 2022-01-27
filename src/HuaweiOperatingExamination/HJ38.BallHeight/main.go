package main

import "fmt"

func main() {
    var n int
    fmt.Scanln(&n)
    fmt.Printf("%0.6f\n",float64(n)*2.875)
    fmt.Printf("%0.6f\n",0.03125*float64(n))
}