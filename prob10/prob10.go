/*
    Project Euler Project No. 10 - Summation of Primes
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
 */

package main

import "fmt"

const Max int = 2000000

/** 
 sieveOfEratoshenes - a basic implementation of the sieve of Eratosthenes
 
 This is of course basic stuff: find prime numbers up to given 'max' value.
 The function returns the slice of bool: the numbers are actually an index
 of the slice, while boolean values represent whether an index is a prime or
 not. If value for a particular index is 'false', the number is a prime.
 Note: why 'false'? Because the []bool is initialized to 'false' by compiler,
 so we use 'true' to mark the numbers that are not primes.
 
 Usage:
    primes := sieveOfEratosthenes(1000)
    for num, status := range primes {
        if !status {
           current_prime := num
           // do whatever you want with the current prime
        }
    }
 */
func sieveOfEratosthenes(max int) []bool {

    primes := make([]bool, max + 1)
    // NOTE: all slice elements are initialized to false

    primes[0] = true
    primes[1] = true
    for i := 2; i*i <= max; i++ {
        step := i
        if primes[i] { continue }
        for j := i + step; j <= max; j += step {
            primes[j] = true
        }
    }
    return primes

}

func main() {

    c := `Project Euler Project No. 10 - Summation of Primes
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.`
    fmt.Println(c)
    fmt.Println()

    primes := sieveOfEratosthenes(Max)
    // now we do the summation
    sum := 0
    for cnt, val := range primes {
        if !val { sum += cnt }
    }
    fmt.Printf("The sum of primes up to %d is %d.\n", Max, sum)
}
