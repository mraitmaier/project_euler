/*
    Project Euler Problem No. 18 - Maximum path sum I

Starting at the top of the triangle below and moving to adjacent numbers on 
the row below, the maximum total from top to bottom is 23.

         3
        7 4
       2 4 6
      8 5 9 3

That is, 3 + 7 + 4 + 9 = 23.

Find the maximum total from top to bottom of the triangle below:

                75
               95 64
              17 47 82
             18 35 87 10
            20 04 82 47 65
           19 01 23 75 03 34
          88 02 77 73 07 63 67
         99 65 04 28 06 16 70 92
        41 41 26 56 83 40 80 70 33
       41 48 72 33 47 32 37 16 94 29
      53 71 44 65 25 43 91 52 97 51 14
     70 11 33 28 77 73 17 78 39 68 17 57
    91 71 52 38 17 14 91 43 58 50 27 29 48
   63 66 04 68 89 53 67 30 73 16 69 87 40 31
  04 62 98 27 23 09 70 98 73 93 38 53 60 04 23

NOTE: As there are only 16384 routes, it is possible to solve this problem by 
trying every route. However, Problem 67, is the same challenge with a triangle 
containing one-hundred rows; it cannot be solved by brute force and requires a 
clever method! ;o)
*/

package main

import (
    "fmt"
)

var l1 []int = []int{75}
var l2 []int = []int{95, 64}
var l3 []int = []int{17, 47, 82}
var l4 []int = []int{18, 35, 87, 10}
var l5 []int = []int{20,  4, 82, 47, 65}
var l6 []int = []int{19,  1, 23, 75,  3, 34}
var l7 []int = []int{88,  2, 77, 73,  7, 63, 67}
var l8 []int = []int{99, 65,  4, 28,  6, 16, 70, 92}
var l9 []int = []int{41, 41, 26, 56, 83, 40, 80, 70, 33}
var l10 []int =[]int{41, 48, 72, 33, 47, 32, 37, 16, 94, 29}
var l11 []int =[]int{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14}
var l12 []int =[]int{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57}
var l13 []int =[]int{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48}
var l14 []int =[]int{63, 66,  4, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31}
var l15 []int =[]int{ 4, 62, 98, 27, 23,  9, 70, 98, 73, 93, 38, 53, 60,  4, 23}
var Triangle  [][]int = [][]int{ l1, l2, l3, l4, l5, l6, l7, l8, l9,
                                 l10, l11, l12, l13, l14, l15 }
/* 
Takes two lines and reduces them to a single line by comparing the sums of 
adjacent values. Returns this new line (must have the same length as the 
shorter one).
*/
func reduce_triangle(shorter, longer []int) []int {
    //length := len(shorter)
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

func main() {
c := `    Project Euler Problem No. 18 - Maximum path sum I

Starting at the top of the triangle below and moving to adjacent numbers 
on the row below, the maximum total from top to bottom is 23.

         3
        7 4
       2 4 6
      8 5 9 3

That is, 3 + 7 + 4 + 9 = 23.

Find the maximum total from top to bottom of the triangle below:

                75
               95 64
              17 47 82
             18 35 87 10
            20 04 82 47 65
           19 01 23 75 03 34
          88 02 77 73 07 63 67
         99 65 04 28 06 16 70 92
        41 41 26 56 83 40 80 70 33
       41 48 72 33 47 32 37 16 94 29
      53 71 44 65 25 43 91 52 97 51 14
     70 11 33 28 77 73 17 78 39 68 17 57
    91 71 52 38 17 14 91 43 58 50 27 29 48
   63 66 04 68 89 53 67 30 73 16 69 87 40 31
  04 62 98 27 23 09 70 98 73 93 38 53 60 04 23 `

    fmt.Println(c)
    fmt.Println()

    sum := triangle_sum(Triangle)
    fmt.Printf("The largest sum in given triangle is %d.\n", sum)
}
