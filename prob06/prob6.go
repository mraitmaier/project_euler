/*
    Project Euler Problem No.6 - Sum Square Difference
The sum of the squares of the first ten natural numbers is,
        12 + 22 + ... + 102 = 385
The square of the sum of the first ten natural numbers is,
        (1 + 2 + ... + 10)2 = 552 = 3025
Hence the difference between the sum of the squares of the first ten 
natural numbers and the square of the sum is 3025 - 385 = 2640.

Find the difference between the sum of the squares of the first one 
hundred natural numbers and the square of the sum.
*/
package main

import (
    "fmt"
)

const Max int = 100

func sum_of_squares(max int) (sum int) {
    sum = 0
    for i := 1; i <= max; i++ {
        sum += i * i
    }
    return
}

func square_of_sum(max int) int {
    sum := 0
    for i := 1; i <= max; i++ {
        sum += i
    }
    return sum * sum
}

func main() {

    c:= `Project Euler Problem No.6 - Sum Square Difference
The sum of the squares of the first ten natural numbers is,
        12 + 22 + ... + 102 = 385
The square of the sum of the first ten natural numbers is,
        (1 + 2 + ... + 10)2 = 552 = 3025
Hence the difference between the sum of the squares of the first ten 
natural numbers and the square of the sum is 3025 - 385 = 2640.

Find the difference between the sum of the squares of the first one 
hundred natural numbers and the square of the sum.`

    diff := square_of_sum(Max) - sum_of_squares(Max)
    fmt.Println(c)
    fmt.Println()
    fmt.Printf("The difference is %d.\n", diff)
}

