/*
    Project Euler Problem No. 9 - Speacial Pythagorean Triplet

A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
                        sqr(a) + sqr(b) = sqr(c)
For example, sqr(3) + sqr(4) = 9 + 16 = 25 = sqr(5).

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product a*b*c.
 */
package main

import (
    "fmt"
)

func check_triplet(a, b, c int) bool {
    pitagorean := (a*a + b*b) == c*c
    if pitagorean && (a + b + c == 1000) { return true }
    return false
}

func main() {
    c := `Project Euler Problem No. 9 - Speacial Pythagorean Triplet

A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
                        sqr(a) + sqr(b) = sqr(c)
For example, sqr(3) + sqr(4) = 9 + 16 = 25 = sqr(5).

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product a*b*c.`

    var i,j,k int
    for i = 100; i < 500; i++ {
        for j = i + 1; j < 500; j++ {
            for k = j + 1; k < 500; k++ {
                if i == j && i == k { continue }
                if check_triplet(i, j, k) { goto out }
            }
        }
    }

out:

    fmt.Println(c)
    fmt.Println()
    fmt.Printf("The sum of %d, %d, %d is %d.\n", i, j, k, i + j + k)
    fmt.Printf("The product of %d, %d, %d is %d.\n", i, j, k, i * j * k)
}

