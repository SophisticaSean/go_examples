// https://projecteuler.net/problem=1

//If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
//Find the sum of all the multiples of 3 or 5 below 1000.

package main

import (
"fmt"
"os"
"strconv"
)

func main() {
  help_string := "to get the sum of all multiples of 3 or 5 below n type the following: go run euler_1.go n"
  if len(os.Args) != 2 {
    fmt.Println(help_string)
  }
  var limit, _ = strconv.Atoi(os.Args[1])

  if limit == 0 {
    fmt.Println(help_string)
    fmt.Println("n needs to be an integer greater than 3")
  }

  var s []int
  for i := 1; i < limit; i++ {
    if i%3 == 0 {
      s = append(s, i)
    }

    if i%5 == 0 {
      s = append(s, i)
    }
  }
  total := 0
  for i := len(s) - 1; i >= 0; i-- {
    total = total + s[i]
  }
  fmt.Println(total)
}
