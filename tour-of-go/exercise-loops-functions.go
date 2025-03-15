package main

import (
    "fmt"
    "math"
)

// Sqrt computes the square root of x using Newton's method.
func Sqrt(x float64) float64 {
    z := 1.0
    for i := 0; i < 10; i++ {
        z -= (z*z - x) / (2 * z)
        fmt.Println("Iteration", i+1, ":", z)
    }
    return z
}

func main() {
    x := 2.0
    fmt.Println("Computed sqrt:", Sqrt(x))
    fmt.Println("math.Sqrt:", math.Sqrt(x))
}

