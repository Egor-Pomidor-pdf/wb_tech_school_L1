package main

import (
	"fmt"
	"math/big"
)



 


func main() {
  a := big.NewInt(2000000000)
  b := big.NewInt(3000000000)

  sum := new(big.Int).Add(a, b)
  fmt.Printf("Sum: %v\n", sum)

  diff := new(big.Int).Sub(a, b)
  fmt.Printf("diff: %v\n", diff)

  mult := new(big.Int).Mul(a, b)
  fmt.Printf("mult: %v\n", mult)

  div := new(big.Int).Div(a, b)
  fmt.Printf("div: %v\n", div)
}
