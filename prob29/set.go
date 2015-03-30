/*
Implements the Set type(s).
*/
package main

import (
    "fmt"
    "math/big"
    "sort"
)

/* we define large integers as keys as string, type big.Int is a pointer and
 cannot be used as key 
 */
type BigIntSet map[string]bool

func (self BigIntSet) In(num big.Int) bool { return self[num.String()] }

func (self BigIntSet) String() string {
    return fmt.Sprintf("Set has currently %d elements.\n", len(self))
}

func (self BigIntSet) Add(num big.Int) {
    if !self.In(num) { self[num.String()] = true }
}

func (self BigIntSet) Len() int { return len(self) }

func (self BigIntSet) Remove(num big.Int) { delete(self, num.String()) }

func (self BigIntSet) Elements() string {
    s := ""
    for key, _ := range self {
        s = fmt.Sprintf("%s%s ", s, key)
    }
    return fmt.Sprintf("%s\n", s)
}

func (self BigIntSet) Sorted() []string {
    keys := make([]string, 0)
    for key, _ := range self {
        keys = append(keys, key)
    }
    sort.Strings(keys)
    return keys
}
