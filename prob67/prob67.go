/*
    Project Euler Problem No. 67 - Maximum path sum II

Starting at the top of the triangle below and moving to adjacent numbers on 
the row below, the maximum total from top to bottom is 23.

         3
        7 4
       2 4 6
      8 5 9 3

That is, 3 + 7 + 4 + 9 = 23.

Find the maximum total from top to bottom in 'triangle.txt', a 15K text 
file containing a triangle with one-hundred rows.

NOTE: This is a much more difficult version of Problem 18. It is not possible 
to try every route to solve this problem, as there are 2**99 altogether! If 
you could check one trillion (10**12) routes every second it would take over 
twenty billion years to check them all. There is an efficient algorithm to 
solve it. ;o)
*/

package main

import (
    "fmt"
    "strings"
    "strconv"
)

var DataFilePath = "triangle.txt"

/* 
Takes two lines and reduces them to a single line by comparing the sums of 
adjacent values. Returns this new line (must have the same length as the 
shorter one).
*/
func reduce_triangle(shorter, longer []int) []int {
    new_line := make([]int, 0)
    // iterate over shorter line and compare the adjacent values in longer line
    for ix, val := range shorter {
        sum1 := val + longer[ix]
        sum2 := val + longer[ix + 1]
        if sum1 > sum2 {
            new_line = append(new_line, sum1)
        } else {
            new_line = append(new_line, sum2)
        }
    }
    return new_line
}

/* 
Calculates the largest sum in triangle.
We take the "bottom to top" approach. We reduce the problem to two lines at the
time and we substitute those lines with a single line containing the best
sum possibilities. We reduce the triangle in such a way until we reach the
shortest line (the top of the triangle) and we're left with a slice that holds
a single value: this is our sum.
*/
func triangle_sum(triangle [][]int) int {
    sum := 0
    // starting index: we go bottom to top
    ix := len(triangle) - 1 // go slices are indexed starting from 1!
    var new_line []int
    for ix > 0 {
        new_line = reduce_triangle(triangle[ix-1], triangle[ix])
        // delete the longer line
        triangle[ix] = nil
        // and substitute the shorter line with the new calculated one
        triangle[ix-1] = new_line
        ix--
    }
    sum = new_line[0]
    return sum
}

/* Read a data file and return the triangle of integer values that are to be
processed. */
func read_file(path string) [][]int {

    triangle := make([][]int, 0)
    lines, err := ReadLines(path)
    if err != nil { return triangle }
    for _, line := range lines {
        curr := make([]int, 0)
        s := strings.Split(line, " ")
        // convert the []string into []int
        for _, val := range s {
            dsf, _ := strconv.Atoi(val)
            curr = append(curr, dsf)
        }
        triangle = append(triangle, curr)
    }
    return triangle
}

func main() {

c := `    Project Euler Problem No. 67 - Maximum path sum II

Starting at the top of the triangle below and moving to adjacent numbers on 
the row below, the maximum total from top to bottom is 23.

         3
        7 4
       2 4 6
      8 5 9 3

That is, 3 + 7 + 4 + 9 = 23.

Find the maximum total from top to bottom in 'triangle.txt', a 15K text 
file containing a triangle with one-hundred rows.`

    fmt.Println(c)
    fmt.Println()

    sum := triangle_sum(read_file(DataFilePath))
    fmt.Printf("The largest sum in given triangle is %d.\n", sum)
}
