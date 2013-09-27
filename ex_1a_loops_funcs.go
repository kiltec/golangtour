package main

import (
    "fmt"
)

func Sqrt(x float64) float64 {
    z := 100.0
    for i := 0; i <= 10; i++ {
        n := z - (z * z - x) / (2 * z)
        z = n
    }
    return z
}

func main() {
    const x = 2
    fmt.Println(Sqrt(x))
}
