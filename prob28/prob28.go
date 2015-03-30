/*
    Project Euler Problem No. 28 - Number Spiral Diagonals

Starting with the number 1 and moving to the right in a clockwise direction 
a 5 by 5 spiral is formed as follows:

        21 22 23 24 25
        20  7  8  9 10
        19  6  1  2 11
        18  5  4  3 12
        17 16 15 14 13

It can be verified that the sum of the numbers on the diagonals is 101.

What is the sum of the numbers on the diagonals in a 1001 by 1001 spiral formed
in the same way?
 */

package main

import "fmt"

const Dim int = 1001

func main() {

    c:= `Project Euler Problem No. 28 - Number Spiral Diagonals
Starting with the number 1 and moving to the right in a clockwise direction 
a 5 by 5 spiral is formed as follows:

        21 22 23 24 25
        20  7  8  9 10
        19  6  1  2 11
        18  5  4  3 12
        17 16 15 14 13

It can be verified that the sum of the numbers on the diagonals is 101.

What is the sum of the numbers on the diagonals in a 1001 by 1001 spiral formed
in the same way?`
    fmt.Println(c)
    fmt.Println()

    tbl := new(SpiralTable)
    err := tbl.Create(Dim)
    if err != nil {
        fmt.Println("ERROR:", err)
    } else {
        fmt.Printf("The sum %d spiral table is %d.\n", Dim,
                tbl.CalculateDiagonalSum())
        }
}
