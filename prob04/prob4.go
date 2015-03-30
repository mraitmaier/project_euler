/*
    Project Euler Problem No. 4 - Largest Palindrome Product
A palindromic number reads the same both ways. 
The largest palindrome made from the product of two 2-digit numbers is 
9009 = 91 * 99.

Find the largest palindrome made from the product of two 3-digit numbers.
 */
package main

import (
    "fmt"
    "strconv"
)

func reverseStr(s string) string {
    runes := []rune(s)
    for i,j := 0, len(runes) - 1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func isPalindrome(s string) bool {
    reversed := reverseStr(s)
    for i := 0; i < len(s) - 1; i++ {
       if s[i] != reversed[i] { return false }
    }
    return true
}

func main() {

    c := `Project Euler Problem No. 4 - Largest Palindrome Product
A palindromic number reads the same both ways. 
The largest palindrome made from the product of two 2-digit numbers is 
9009 = 91 * 99.

Find the largest palindrome made from the product of two 3-digit numbers.`

    fmt.Println(c)
    fmt.Println()

    prod := 0
    largest := 0
    for i := 100; i < 1000  ; i++ {
        for j := 100; j < 1000; j++ {
            prod = i * j
            if isPalindrome(strconv.Itoa(prod)) {
                //fmt.Printf("i=%d, j=%d\n", i, j) // DEBUG
                if prod > largest { largest = prod }
            }
        }
    }
    fmt.Printf("The largest palindrome product is %d.\n", largest)
}
