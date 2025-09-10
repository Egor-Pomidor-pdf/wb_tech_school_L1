package main

import "fmt"

func delElementFromSlice[T any](sl []T, i int) []T {
  copy(sl[i:], sl[i+1:])
  var zero T
  sl[len(sl)-1] = zero
  return sl[:len(sl) - 1]
}

func main() {
  sl := []int{1, 2, 5, 7, 9}
  sl = delElementFromSlice(sl, 1)
  fmt.Println(sl)
}
