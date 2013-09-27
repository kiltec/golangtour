package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) (float64, int) {
    z := x
    delta := 1e-10
    iters := 0
    for {
        n := z - (z * z - x) / (2 * z)
        iters++
        if math.Abs(n - z) < delta {
            break
        }
        z = n
    }
    return z, iters
}

func main() {
    const x = 1000
    result, iters := Sqrt(x)
    actual := math.Sqrt(x)
    diff := math.Abs(result - actual)
    fmt.Println(iters, result, actual, diff)

}
