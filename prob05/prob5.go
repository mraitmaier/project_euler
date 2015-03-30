/*
 * Project Euler Problem No. 5 - Smallest Multiple
 * 2520 is the smallest number that can be divided by each of the numbers 
 * from 1 to 10 without any remainder.
 *
 * What is the smallest positive number that is evenly divisible by all 
 * of the numbers from 1 to 20?
 */
package main

import "fmt"

/* since we lift values by 20, we need to check only prime numbers and the 
squares, cubes etc. */

/* the first version of the test function uses a tree of if-else clauses */
func test_it1(val int) bool {
    if !(val % 19 == 0) {
            return false
    } else if !(val % 17 == 0) {
            return false
    } else if !(val % 16 == 0) {
            return false
    } else if !(val % 13 == 0) {
            return false
    } else if !(val % 11 == 0) {
            return false
    } else if !(val % 9 == 0) {
            return false
    } else if !(val % 7 == 0) {
            return false
    } else if !(val % 3 == 0) {
            return false
    }
    return true
}

/* the second version of the test function uses a switch clause */
func test_it2(val int) bool {
    switch {
        case !(val % 19 == 0): { return false }
        case !(val % 17 == 0): { return false }
        case !(val % 16 == 0): { return false }
        case !(val % 13 == 0): { return false }
        case !(val % 11 == 0): { return false }
        case !(val %  9 == 0): { return false }
        case !(val %  7 == 0): { return false }
        case !(val %  3 == 0): { return false }
    }
    return true
}


func main () {

    c := `Project Euler Problem No. 5 - Smallest Multiple
2520 is the smallest number that can be divided by each of the numbers 
from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all 
of the numbers from 1 to 20?`

    val := 20;
    for ; !(test_it2(val)); val += 20 { /* empty body */ }
    fmt.Println(c)
    fmt.Println()
    fmt.Printf("The smallest number is %d.\n", val)
}

