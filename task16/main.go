package main

import "fmt"


func quickSort(sl []int) {
  if len(sl) < 2 {
    return
  }

  pr := sl[0]
  var left, right []int

  for _, num := range sl[1:] {
    if num < pr {
      left = append(left, num)
    } else {
      right = append(right, num)
    }
  }
  quickSort(left)
  quickSort(right)

  result := append(append(left, pr), right...)
  copy(sl, result)

}

func main() {
  icx := []int{5, 2, 4, 7, 1, -2, 44, 22, 2}
  quickSort(icx)
  fmt.Println(icx)

}
