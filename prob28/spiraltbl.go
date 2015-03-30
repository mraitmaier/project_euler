/*
    spiraltbl.go - an implementation of the SpiralTable struct

    Project Euler Problem No. 28 - Number Spiral Diagonals
    
 */

package main

import (
    "fmt"
    "errors"
)


/*
    We define a spiral table as structure that holds the table itself and a
    side dimension (this value comes handy...).
 */
type SpiralTable struct {

    // dimension of the table
    dim int

    // table itself
    table [][]int
}

/* create a table with given dimension (must be odd value!) */
func (s *SpiralTable) Create(dim int) error {
    // check dimension value
    if dim % 2 == 0 {
        msg := fmt.Sprintf("Table dimensions cannot be even: %d", dim)
        return errors.New(msg)
    }
    // init struct
    s.dim = dim
    s.table = make([][]int, s.dim)
    for i := range s.table {
        s.table[i] = make([]int, s.dim)
    }
    // we locate the middle of the table as a starting point; 
    start := int(dim / 2)

    //  let's begin
    i, j, cnt := start, start, 1
    s.table[i][j] = cnt
    limit := s.dim*s.dim
    // As we spiral along the table, the step is incremented and the following
    // rule applies:
    // - if step value is even, we go first left and then up;
    // - if step value is odd, we go first right and then down.
    // We always end up creating table in the upper right corner while moving
    // right - so we need to check the current value so that we don't break the
    // table borders.
    for step := 1; ; step++ {
        if step % 2 == 0 {
            // we move left
                i, j, cnt = s.left(i, j, step, cnt)
            // and then up
                i, j, cnt  = s.up(i, j, step, cnt)
        } else {
            // we move right
                i, j, cnt = s.right(i, j, step, cnt)
                if cnt >= limit { break } //<-- loop break point
            // and down
                i, j, cnt = s.down(i, j, step, cnt)
        }
    }
    return nil
}

/* creating table - move left */
func (s *SpiralTable) left(x, y, step, val int) (int, int, int) {
    new_y := 0
    new_val := val
    for i := 1; i <= step; i++ {
        new_y = y - i
        new_val = val + i
        s.table[x][new_y] = new_val
    }
    return x, new_y, new_val
}

/*
 creating table - move right 
 we always end up creating a table moving right, so we must check we didn't
 exceed the value limit (that is a square of a table dimension)
 */
func (s *SpiralTable) right(x, y, step, val int) (int, int, int) {
    new_y := 0
    new_val := val
    for i := 1; i <= step; i++ {
        new_y = y + i
        new_val = val + i
        if new_val > (s.dim*s.dim)  { break } // <--- check value
        s.table[x][new_y] = new_val
    }
    return x, new_y, new_val
}

/* creating table - move up */
func (s *SpiralTable) up(x, y, step, val int) (int, int, int) {
    new_x := 0
    new_val := val
    for i := 1; i <= step; i++ {
        new_x = x - i
        new_val = val + i
        s.table[new_x][y] = new_val
    }
    return new_x, y, new_val
}

/* creating table - move down */
func (s *SpiralTable) down(x, y, step, val int) (int, int, int) {
    new_x := 0
    new_val := val
    for i := 1; i <= step; i++ {
        new_x = x + i
        new_val = val + i
        s.table[new_x][y] = new_val
    }
    return new_x, y, new_val
}

func (s *SpiralTable) Error() string {
    return "Table dimensions cannot be even"
}

func (s *SpiralTable) String() string {
    return fmt.Sprintf("The Spiral table has a dimension of %d\n", s.dim)
}

/* this is the Project Euler Problem No. 28: we need to calculate the sum of
 * elements that are on the main diagonals; this method does just that */
func (s *SpiralTable) CalculateDiagonalSum() int {
    sum := 0
    // the first diagonal
    for i := 0; i <= (s.dim - 1); i++ {
        sum += s.table[i][i]
    }
    // the second diagonal
    for i, j := 0, s.dim - 1; i <= (s.dim - 1) && s.dim >= 0; i,j = i+1, j-1 {
        sum += s.table[i][j]
    }
    // XXX: the actual sum must be decremented: since the middle of the table
    // is summed twice and this is our starting point for creating the
    // table - it holds value of 1, of course.
    return sum - 1
}

