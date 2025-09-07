package main

import "fmt"


func binSearch(sl []int, numSearch int) int {
  return binSearchHelp(sl, numSearch, 0)
}

func binSearchHelp(sl []int, numSearch int, offset int) int {
    if len(sl) == 0 {
      return -1
    }
    mid := len(sl)/2
    val := sl[mid]
    if val == numSearch {
      return offset + mid
    } else if val < numSearch {
      return binSearchHelp(sl[mid + 1:], numSearch, mid + offset + 1)
    } else if val > numSearch {
      return binSearchHelp(sl[:mid], numSearch, offset)
    }
    return -1
  }



func main() {
  fmt.Println(binSearch([]int{1, 2, 4, 5, 10, 11}, 11))

}
