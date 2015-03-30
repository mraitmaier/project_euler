/*
    Project Euler Problem No. 50 - Consecutive Prime Sum

The prime 41, can be written as the sum of six consecutive primes:
        41 = 2 + 3 + 5 + 7 + 11 + 13
This is the longest sum of consecutive primes that adds to a prime below
one-hundred.
The longest sum of consecutive primes below one-thousand that adds to a
prime, contains 21 terms, and is equal to 953.

Which prime, below one-million, can be written as the sum of the most
consecutive primes?
*/
package main

import "fmt"

/* let's define an auxiliary type for consecutive sums */
type consecutiveSum struct {

    // a sum of consecutive primes
    sum int

    // number of primes needed to get to the sum
    primes int
}

const Max int = 1e6

/* we use basic sieve of Eratosthenes to find the primes */
func sieveOfEratosthenes(max int) []bool {

	primes := make([]bool, max+1)
	// NOTE: all slice elements are initialized to false

	primes[0] = true
	primes[1] = true
	for i := 2; i*i <= max; i++ {
		step := i
		if primes[i] {
			continue
		}
		for j := i + step; j <= max; j += step {
			primes[j] = true
		}
	}
	return primes
}

/* let's seed the right prime number and number of conscutive primes */
func seedThePrime(max int) (int, int) {

    // we find all primes up to limit
	primes := sieveOfEratosthenes(max)

    // and we calculate all prime-consecutive sums that are smaller than limit
    sums := findAllSums(max, primes)
    l := len(sums)

    // we must now go through a list of sums and check if it is a prime and how
    // many primes are needed to get to that sum
	prime := 0
    num := 0
    for i := 0; i < l; i++ {
        if !primes[sums[i].sum] && sums[i].primes > num {
            prime = sums[i].sum
            num = sums[i].primes
        }
	}
	return prime, num
}

/* find all possible prime-consecutive sums up to limit ('max') and return the
 * list of these sums */
func findAllSums(max int, primes []bool) []consecutiveSum {

    sums := make([]consecutiveSum, 0)

    for low := 2; low <= max; low++ {
        cnt, s := 0, 0
        if !primes[low] {
            for high := low; high <= max; high++ {
                if !primes[high] {
                    s = s + high
                    cnt++
                    // we append the current sum...
                    if s < max  && cnt > 1 {
                        sums = append(sums, consecutiveSum{s, cnt})
                    }
                    // if current sum exceeds limit...
                    if s > max { break }
                }
            }
        }
   }
   return sums
}

func main() {
    c := `Project Euler Problem No. 50 - Consecutive Prime Sum

The prime 41, can be written as the sum of six consecutive primes:
        41 = 2 + 3 + 5 + 7 + 11 + 13
This is the longest sum of consecutive primes that adds to a prime below
one-hundred.
The longest sum of consecutive primes below one-thousand that adds to a
prime, contains 21 terms, and is equal to 953.

Which prime, below one-million, can be written as the sum of the most
consecutive primes?`
    fmt.Println(c)
    fmt.Println()

	prime, num := seedThePrime(Max)

	fmt.Printf("The prime is %d, it is a sum of %d consecutive primes.\n",
            prime, num)
}
