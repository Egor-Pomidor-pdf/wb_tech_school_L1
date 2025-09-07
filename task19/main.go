package main

import "fmt"

func reverseString(str string) string {
  result := make([]rune, len(str))
  for i, val := range []rune(str) {
    result[len(str) - i - 1] = val
  }
  return string(result)

}

func main() {
  fmt.Println(reverseString("abc"))
}