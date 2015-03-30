/*
    Project Euler Problem No. 7 - 10001st Prime

By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see 
that the 6th prime is 13.

What is the 10001st prime number?
 */
package main

import "fmt"

const Max int = 105000 // Max set by trial-and-error method

func main() {
    c := `Project Euler Problem No. 7 - 10001st Prime
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see 
that the 6th prime is 13.

What is the 10001st prime number?`
    fmt.Println(c)
    fmt.Println()

    primes := make([]bool, Max)
    // NOTE: all slice elements are initialized to false

    // implementing the basic sieve of erathostenes
    // we assume that 0 and 1 are not primes; just to shorten loop a bit
    primes[0] = true
    primes[1] = true
    for i := 2; i*i <= Max; i = i+1 {
        step := i
        if primes[i] { continue }
        for j := i + step; j < Max; j += step {
            primes[j] = true
        }
    }
    // now we seed the right prime
    numOfPrimes := 0
    theRightPrime := 0
    for cnt, val := range primes {
        if !val {
            numOfPrimes++
            if numOfPrimes == 10001 {
                theRightPrime = cnt
                break
            }
        }
    }
    fmt.Printf("The 10001th prime is %d.\n", theRightPrime)
}
