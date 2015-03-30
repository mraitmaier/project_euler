/*
    Project Euler Problem No. 3 - Largest Prime Factor

The prime factors of 13195 are 5, 7, 13 and 29.
What is the largest prime factor of the number 600851475143?
 */
package main

import "fmt"

const Num int = 600851475143

func isPrime(n int) bool {
    for i := 2; i < n / 2; i++ {
        if n % i == 0 {
            return false
        }
    }
    return true
}

func main() {
    c := `Project Euler Problem No. 3 - Largest Prime Factor

The prime factors of 13195 are 5, 7, 13 and 29.
What is the largest prime factor of the number 600851475143?`
    fmt.Println(c)
    fmt.Println()

   for i := 2; i < Num / 2; i++ {
        if Num % i == 0 {
            pair := Num / i
            if isPrime(pair) {
                fmt.Printf("The largest prime factor of %d is %d.\n", Num, pair)
                break
            }
        }
    }
}
