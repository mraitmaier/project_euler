/*
    Project Euler Problem No. 14 - Longest Collatz Sequence

The following iterative sequence is defined for the set of positive integers:

    n  --> n/2 (n is even)
    n  --> 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:
        13 -> 40 -> 20 -> 10 -> 5 -> 16 -> 8 -> 4 -> 2 -> 1
It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.
Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.
 */

package main

import "fmt"

const Limit int = 1000000
//const Limit int = 100

func CollatzSequence(start int) []int {

    // we initialize the default Collatz Sequence to 1 element and set the 
    // starting element of the sequence
    seq := make([]int, 1)
    seq[0] = start

    // now we create the sequence
    current := start
    for i :=1; current > 1; i++ {
        if current % 2 == 0 {
            current = current / 2
        } else {
            current = 3 * current + 1
        }
        seq = append(seq, current)
    }
    return seq
}

func main() {

    c := ` Project Euler Problem No. 14 - Longest Collatz Sequence
The following iterative sequence is defined for the set of positive integers:
    n  --> n/2 (n is even)
    n  --> 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:
        13 -> 40 -> 20 -> 10 -> 5 -> 16 -> 8 -> 4 -> 2 -> 1
It can be seen that this sequence (starting at 13 and finishing at 1) contains
10 terms. Although it has not been proved yet (Collatz Problem), it is thought
that all starting numbers finish at 1.
Which starting number, under one million, produces the longest chain? `
    fmt.Println(c)
    fmt.Println()

    longest, l, number := 0, 0, 0
    for i := 2; i <= Limit; i++ {
        s := CollatzSequence(i)
        l = len(s)
        if l > longest {
            longest = l
            number = i
        }
    }
    fmt.Printf("The longest sequence %d is produced by %d.\n", longest, number)
}
