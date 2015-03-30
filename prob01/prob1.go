/*
 * Project Euler problem No. 1
 *
 * If we list all the natural numbers below 10 that are multiples of 3 or 5, 
 * we get 3, 5, 6 and 9. The sum of these multiples is 23.
 * Find the sum of all the multiples of 3 or 5 below 1000.
 */
package main

import (
    "fmt"
)

const (
    Max int = 1000
)

func main() {

    c := `If we list all the natural numbers below 10 that are multiples of 3
or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
Find the sum of all the multiples of 3 or 5 below 1000.`

    sum := 0
    for i := 1; i < Max; i ++ {
        if i % 3 == 0  || i % 5 == 0 {
            sum += i
        }
    }
    fmt.Println(c)
    fmt.Println()
    fmt.Printf("The sum is %d.\n", sum)
}
