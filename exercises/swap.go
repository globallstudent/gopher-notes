// Swapping two numbers without using temporary variable

package main

import "fmt"

func swap(a, b int) (int, int) {
    return b, a
}

func main() {
    x, y := swap(3, 4)
    fmt.Printf("Swapped %d and %d: %d, %d\n", 3, 4, x, y)
}
